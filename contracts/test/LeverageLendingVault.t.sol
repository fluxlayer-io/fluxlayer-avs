// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../src/LeverageLendingVault.sol";
import "../src/ERC20Mock.sol";

// Unit tests for Layer III's leverage lending vault against the formal model in the thesis
// (Section "Layer III: Under-Collateralised Leverage Lending Vault", Eq. 8-14 and Theorem 5.2's
// proof). Parameters match the thesis's own committed values: lambda_max = 10x,
// Delta = 15 minutes.
contract LeverageLendingVaultTest is Test {
    LeverageLendingVault public vault;
    ERC20Mock public token;

    address public keeper = address(this); // vault owner acts as the fluxlayer-mpc keeper
    address public lp = address(0x1000);
    address public maker = address(0x2000);

    uint256 public constant LAMBDA_MAX_PCT = 1000; // 10x, scaled by LEVERAGE_DENOMINATOR=100
    uint256 public constant DELTA = 15 minutes;

    bytes32 public orderId = keccak256("order-1");
    address public execAddress = address(0x3000); // stand-in for the maker's Cobo MPC wallet

    function setUp() public {
        token = new ERC20Mock();
        vault = new LeverageLendingVault(IERC20(address(token)), LAMBDA_MAX_PCT, DELTA);

        token.mint(lp, 1_000_000 ether);
        vm.prank(lp);
        vault.deposit(1_000_000 ether);

        vault.registerInternalExecutionAddress(orderId, execAddress);
        token.mint(maker, 10_000 ether);
    }

    function test_OpenPosition_WithinLeverageCap_Succeeds() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether; // exactly 10x, the leverage cap boundary

        uint256 liquidityBefore = vault.availableLiquidity();
        vm.prank(maker);
        token.approve(address(vault), collateral);
        uint256 deadline = vault.openPosition(orderId, maker, collateral, loan);

        assertEq(deadline, block.timestamp + DELTA, "deadline must be t0 + Delta");
        assertEq(token.balanceOf(execAddress), loan, "loan must be disbursed to the registered exec address");
        assertEq(
            vault.availableLiquidity(), liquidityBefore - loan, "LP liquidity must be debited by exactly the loan"
        );

        LeverageLendingVault.Position memory p = vault.getPosition(orderId);
        assertEq(p.collateral, collateral);
        assertEq(p.loanAmount, loan);
        assertEq(uint8(p.status), uint8(LeverageLendingVault.PositionStatus.Open));
    }

    function test_OpenPosition_ExceedingLeverageCap_Reverts() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether + 1; // 1 wei over the 10x cap

        vm.expectRevert(abi.encodeWithSelector(LeverageLendingVault.LeverageExceeded.selector, loan, 1_000 ether));
        vault.openPosition(orderId, maker, collateral, loan);
    }

    function test_OpenPosition_UnregisteredExecutionAddress_Reverts() public {
        bytes32 unregisteredOrder = keccak256("order-unregistered");
        vm.expectRevert(LeverageLendingVault.UnregisteredExecutionAddress.selector);
        vault.openPosition(unregisteredOrder, maker, 100 ether, 1_000 ether);
    }

    function test_OpenPosition_InsufficientLiquidity_Reverts() public {
        uint256 collateral = 1_000_000 ether; // large enough to satisfy leverage cap alone
        uint256 loan = 10_000_000 ether; // exceeds the LP pool's 1,000,000 ether
        vm.expectRevert(
            abi.encodeWithSelector(LeverageLendingVault.InsufficientLiquidity.selector, loan, 1_000_000 ether)
        );
        vault.openPosition(orderId, maker, collateral, loan);
    }

    function test_CollateralIsNotWithdrawableByLPs() public {
        // Regression test: an earlier draft of this contract credited posted collateral into
        // availableLiquidity, which would have let LPs withdraw a maker's escrowed collateral.
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        uint256 remainingLiquidity = vault.availableLiquidity();
        vm.prank(lp);
        vm.expectRevert("insufficient available liquidity");
        vault.withdraw(remainingLiquidity + 1);

        // the LP can still withdraw exactly what is genuinely undeployed
        vm.prank(lp);
        vault.withdraw(remainingLiquidity);
        assertEq(vault.availableLiquidity(), 0);
        // but the maker's collateral, sitting in the vault's token balance, was never touched
        assertEq(token.balanceOf(address(vault)), collateral);
    }

    function test_Repay_Full_ClosesPositionAndReturnsCollateral() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        token.mint(keeper, loan);
        token.approve(address(vault), loan);
        vault.repay(orderId, loan);

        LeverageLendingVault.Position memory p = vault.getPosition(orderId);
        assertEq(uint8(p.status), uint8(LeverageLendingVault.PositionStatus.Repaid));
        assertEq(p.collateral, 0, "collateral must be released on full repayment");
        assertEq(token.balanceOf(maker), 10_000 ether, "maker's collateral must be returned in full");
    }

    function test_Repay_Partial_KeepsPositionOpen() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        token.mint(keeper, 400 ether);
        token.approve(address(vault), 400 ether);
        vault.repay(orderId, 400 ether);

        LeverageLendingVault.Position memory p = vault.getPosition(orderId);
        assertEq(uint8(p.status), uint8(LeverageLendingVault.PositionStatus.Open));
        assertEq(p.repaidAmount, 400 ether);
        assertEq(p.collateral, collateral, "collateral stays escrowed until fully repaid");
    }

    function test_Repay_AfterDeadline_Reverts() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        uint256 deadline = vault.openPosition(orderId, maker, collateral, loan);

        vm.warp(deadline + 1);
        token.mint(keeper, loan);
        token.approve(address(vault), loan);
        vm.expectRevert(abi.encodeWithSelector(LeverageLendingVault.DeadlinePassed.selector, orderId, deadline));
        vault.repay(orderId, loan);
    }

    function test_Repay_ExceedingOutstanding_Reverts() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        token.mint(keeper, loan + 1 ether);
        token.approve(address(vault), loan + 1 ether);
        vm.expectRevert(abi.encodeWithSelector(LeverageLendingVault.RepayExceedsOutstanding.selector, loan + 1 ether, loan));
        vault.repay(orderId, loan + 1 ether);
    }

    function test_Liquidate_BeforeDeadline_Reverts() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        vm.expectRevert(
            abi.encodeWithSelector(LeverageLendingVault.DeadlineNotPassed.selector, orderId, block.timestamp + DELTA)
        );
        vault.liquidate(orderId);
    }

    // At lambda_max = 10x, collateral (C) covers only 1/10th of an unpaid loan (L): if a
    // position is never repaid at all, liquidation cannot make the LP pool whole purely from
    // collateral. This is the exact risk the thesis describes as "a modest increase in risk
    // relative to over-collateralised lending" (Table "Comparison of Lending Types") and which
    // Delta = 15 minutes / assumption A4 (high-liquidity collateral, negligible price movement)
    // is relied on to keep rare in practice by ensuring repayment normally completes first.
    function test_Liquidate_NoRepayment_SeizesCollateralOnly_LPsBearShortfall() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);
        uint256 liquidityAfterOpen = vault.availableLiquidity();

        vm.warp(block.timestamp + DELTA + 1);
        vault.liquidate(orderId);

        LeverageLendingVault.Position memory p = vault.getPosition(orderId);
        assertEq(uint8(p.status), uint8(LeverageLendingVault.PositionStatus.Liquidated));
        assertEq(p.collateral, 0);
        // the full 100 ether collateral is seized (it is less than the 1,000 ether outstanding)
        assertEq(vault.availableLiquidity(), liquidityAfterOpen + collateral);
        // the maker receives no surplus back, having failed to repay
        assertEq(token.balanceOf(maker), 10_000 ether - collateral);
    }

    function test_Liquidate_AfterPartialRepay_SeizesOnlyRemainingOutstanding() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);

        // repay enough that the remaining outstanding (50 ether) is now less than collateral
        token.mint(keeper, 950 ether);
        token.approve(address(vault), 950 ether);
        vault.repay(orderId, 950 ether);

        uint256 liquidityAfterRepay = vault.availableLiquidity();
        vm.warp(block.timestamp + DELTA + 1);
        vault.liquidate(orderId);

        // outstanding = 50 ether < collateral = 100 ether, so only 50 ether is seized and the
        // remaining 50 ether surplus collateral is returned to the maker
        assertEq(vault.availableLiquidity(), liquidityAfterRepay + 50 ether);
        assertEq(token.balanceOf(maker), 10_000 ether - collateral + 50 ether);
    }

    function test_Liquidate_ThenRepay_Reverts() public {
        uint256 collateral = 100 ether;
        uint256 loan = 1_000 ether;
        vault.openPosition(orderId, maker, collateral, loan);
        vm.warp(block.timestamp + DELTA + 1);
        vault.liquidate(orderId);

        token.mint(keeper, loan);
        token.approve(address(vault), loan);
        vm.expectRevert(abi.encodeWithSelector(LeverageLendingVault.PositionNotOpen.selector, orderId));
        vault.repay(orderId, loan);
    }
}
