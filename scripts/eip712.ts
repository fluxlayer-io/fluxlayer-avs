import {ethers, TypedDataDomain} from "ethers";
import * as dotenv from 'dotenv';

dotenv.config();


// The named list of all type definitions
const types = {
    Order: [
        {name: 'orderId', type: 'uint32'},
        {name: 'maker', type: 'address'},
        {name: 'taker', type: 'address'},
        {name: 'inputToken', type: 'address'},
        {name: 'inputAmount', type: 'uint256'},
        {name: 'outputToken', type: 'address'},
        {name: 'outputAmount', type: 'uint256'},
        {name: 'expiry', type: 'uint256'},
        {name: 'targetNetworkNumber', type: 'uint32'},
    ]
};

// The data to sign
export const sampleOrder = {
    orderId: 11,
    maker: '0x6f43BB66B9B22057bb5B752a70481efb12166d0D',
    taker: '0x0000000000000000000000000000000000000000',
    inputToken: '0xC36E19B16c276D10BE40c7b9D3c026E8f482f2aD',
    inputAmount: ethers.getBigInt('10000'),
    outputToken: '0x6CE1551CE42d92e2E70De97d963C027E5a70F621',
    outputAmount: ethers.getBigInt('10000'),
    expiry: 9999999999,
    targetNetworkNumber: 11155111,
};

async function main() {
    const signer = new ethers.Wallet(process.env.MAKER_PRIVATE_KEY!);
    // read maker address
    console.log(signer.address);
    const domain: TypedDataDomain = {
        name: 'OrderBook',
        version: '1.0',
        chainId: 17000,
        verifyingContract: '0x6D26A188C826Af13fF97C8cB70245676af187a2F'
    };
    const signature = await signer.signTypedData(domain, types, sampleOrder);
    // log hashed domain
    console.log(ethers.TypedDataEncoder.hashDomain(domain));
    console.log(signature);
}

main();
