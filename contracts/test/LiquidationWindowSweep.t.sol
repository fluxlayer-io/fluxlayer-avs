// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.12;

import "forge-std/Test.sol";
import "../src/LeverageLendingVault.sol";
import "../src/ERC20Mock.sol";

// Parameter-sensitivity sweep of the liquidation window (Delta, Eq. 13) for Layer III's
// leverage lending vault, crossed against how much of that window has actually elapsed when
// the keeper attempts to close out a position. Each (Delta, elapsed-fraction) cell is run as
// an independent Foundry test against a freshly deployed vault, mirroring the methodology
// used for the DVN quorum-threshold (tau) sweep: reproducible via
// `forge test --match-contract LiquidationWindowSweepTest -vv`.
//
// A cell where elapsed <= Delta exercises the honest path (repay() before the deadline); a
// cell where elapsed > Delta exercises the forced-closeout path (liquidate() after it). Both
// close the position, so both are reported as "Closed" together with which path closed it and
// its gas cost, rather than as an accept/reject outcome as in the tau sweep.
contract LiquidationWindowSweepTest is Test {
    ERC20Mock public token;

    address public lp = address(0x1000);
    address public maker = address(0x2000);
    address public execAddress = address(0x3000);
    bytes32 public orderId = keccak256("order-1");

    uint256 public constant LAMBDA_MAX_PCT = 1000; // 10x
    uint256 public constant COLLATERAL = 100 ether;
    uint256 public constant LOAN = 1_000 ether;

    function _deployVault(uint256 delta) internal returns (LeverageLendingVault) {
        token = new ERC20Mock();
        LeverageLendingVault vault = new LeverageLendingVault(IERC20(address(token)), LAMBDA_MAX_PCT, delta);

        token.mint(lp, 1_000_000 ether);
        vm.prank(lp);
        vault.deposit(1_000_000 ether);

        vault.registerInternalExecutionAddress(orderId, execAddress);
        token.mint(maker, COLLATERAL);
        return vault;
    }

    // elapsedPct is elapsed time as a percentage of delta (e.g. 90 = 90% of the window).
    function _runCell(uint256 delta, uint256 elapsedPct) internal {
        LeverageLendingVault vault = _deployVault(delta);
        vault.openPosition(orderId, maker, COLLATERAL, LOAN);

        uint256 elapsed = (delta * elapsedPct) / 100;
        vm.warp(block.timestamp + elapsed);

        bool repaidBeforeDeadline = elapsed <= delta;
        uint256 gasBefore = gasleft();
        string memory path;
        if (repaidBeforeDeadline) {
            token.mint(address(this), LOAN);
            token.approve(address(vault), LOAN);
            vault.repay(orderId, LOAN);
            path = "REPAY";
        } else {
            vault.liquidate(orderId);
            path = "LIQUIDATE";
        }
        uint256 gasUsed = gasBefore - gasleft();

        LeverageLendingVault.Position memory p = vault.getPosition(orderId);
        LeverageLendingVault.PositionStatus expectedStatus =
            repaidBeforeDeadline ? LeverageLendingVault.PositionStatus.Repaid : LeverageLendingVault.PositionStatus.Liquidated;
        assertEq(uint8(p.status), uint8(expectedStatus), "position must close via the expected path");

        console2.log(
            string.concat(
                "SWEEP delta=",
                vm.toString(delta),
                " elapsedPct=",
                vm.toString(elapsedPct),
                " path=",
                path,
                " gas=",
                vm.toString(gasUsed)
            )
        );
    }

    // Delta = 5 minutes
    function test_Sweep_Delta5min_Elapsed50pct() public { _runCell(5 minutes, 50); }
    function test_Sweep_Delta5min_Elapsed90pct() public { _runCell(5 minutes, 90); }
    function test_Sweep_Delta5min_Elapsed100pct() public { _runCell(5 minutes, 100); }
    function test_Sweep_Delta5min_Elapsed110pct() public { _runCell(5 minutes, 110); }

    // Delta = 15 minutes (the thesis's committed default)
    function test_Sweep_Delta15min_Elapsed50pct() public { _runCell(15 minutes, 50); }
    function test_Sweep_Delta15min_Elapsed90pct() public { _runCell(15 minutes, 90); }
    function test_Sweep_Delta15min_Elapsed100pct() public { _runCell(15 minutes, 100); }
    function test_Sweep_Delta15min_Elapsed110pct() public { _runCell(15 minutes, 110); }

    // Delta = 30 minutes
    function test_Sweep_Delta30min_Elapsed50pct() public { _runCell(30 minutes, 50); }
    function test_Sweep_Delta30min_Elapsed90pct() public { _runCell(30 minutes, 90); }
    function test_Sweep_Delta30min_Elapsed100pct() public { _runCell(30 minutes, 100); }
    function test_Sweep_Delta30min_Elapsed110pct() public { _runCell(30 minutes, 110); }

    // Delta = 60 minutes
    function test_Sweep_Delta60min_Elapsed50pct() public { _runCell(60 minutes, 50); }
    function test_Sweep_Delta60min_Elapsed90pct() public { _runCell(60 minutes, 90); }
    function test_Sweep_Delta60min_Elapsed100pct() public { _runCell(60 minutes, 100); }
    function test_Sweep_Delta60min_Elapsed110pct() public { _runCell(60 minutes, 110); }
}
