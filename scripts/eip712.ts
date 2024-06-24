import {ethers, TypedDataDomain} from "ethers";
import * as dotenv from 'dotenv';

dotenv.config();


// The named list of all type definitions
const types = {
    Order: [
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
const value = {
    maker: '0x2e4B14254CE56D195922Cf1F4C7E97745EE90005',
    taker: '0x0000000000000000000000000000000000000000',
    inputToken: '0xc7183455a4C133Ae270771860664b6B7ec320bB1',
    inputAmount: 100,
    outputToken: '0xa0Cb889707d426A7A386870A03bc70d1b0697598',
    outputAmount: 100,
    expiry: 1001,
    targetNetworkNumber: 17000,
};

async function main() {
    const signer = new ethers.Wallet(process.env.PRIVATE_KEY!);
    const domain: TypedDataDomain = {
        name: 'Settlement',
        version: '1.0',
        chainId: 17000,
        verifyingContract: '0xF62849F9A0B5Bf2913b396098F7c7019b51A820a'
    };
    const signature = await signer.signTypedData(domain, types, value);
    // log hashed domain
    console.log(ethers.TypedDataEncoder.hashDomain(domain));
    console.log(signature);
    // 0xd8829e185303d3eb7347d38e7cbfc4019721fec60b9341886d783354766ed5f4368f39e39202cb861152c0332d66e50960e46f0f5e8f15a122505fa1b3a801dc1c
}

main();