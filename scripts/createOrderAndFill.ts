import {Contract, ethers} from "ethers";
import axios from "axios";
import * as path from "node:path";
import * as fs from "fs";
import * as dotenv from 'dotenv';
// @ts-ignore
import OrderBookABI from './orderbookAbi.json';
// @ts-ignore
import SettlementABI from './settlementAbi.json';

dotenv.config();

export type Order = {
    orderId: number;
    from: string;
    receiver: string;
    sellToken: string;
    sellAmount: string;
    buyToken: string;
    buyAmount: string;
    validTo: number;
    targetNetworkNumber: number;
    signature: string;
}

async function main() {
    // read json from contracts/script/output/flux_layer_avs_deployment_output.json
    const jsonFilePath = path.join(__dirname, '../contracts/script/output/flux_layer_avs_deployment_output.json');
    const jsonData = fs.readFileSync(jsonFilePath, 'utf-8');
    const {addresses} = JSON.parse(jsonData);
    // taker on holesky
    const taker1 = new ethers.Wallet(process.env.PRIVATE_KEY!, new ethers.JsonRpcProvider(process.env.HOLESKY_RPC_URL!));
    // taker on sepolia
    const taker2 = new ethers.Wallet(process.env.PRIVATE_KEY!, new ethers.JsonRpcProvider(process.env.SEPOLIA_RPC_URL!));
    // read order from avs
    const orders = await axios.get<Order[]>('http://127.0.0.1:8080/api/v1/orders').then(res => res.data);
    // // get order with max order id
    const order = orders.reduce((a, b) => a.orderId > b.orderId ? a : b);
    // create order by maker on holesky, by calling createOrder in contract OrderBook
    const orderBook = new Contract(addresses.orderBook, OrderBookABI, taker1);
    const tx = await orderBook.createOrder([order.orderId, order.from, order.receiver, order.sellToken, order.sellAmount, order.buyToken, order.buyAmount, order.validTo, order.targetNetworkNumber], order.signature);
    await tx.wait();
    // log etherscan url
    console.log(`https://holesky.etherscan.io/tx/${tx.hash}`);
    // fulfill order by taker on sepolia, by calling fillOrder in contract Settlement
    const settlement = new Contract(addresses.settlement, SettlementABI, taker2);
    const tx2 = await settlement.fulfill([order.orderId, order.from, order.receiver, order.sellToken, order.sellAmount, order.buyToken, order.buyAmount, order.validTo, order.targetNetworkNumber], 100, ethers.getBytes('0x00'), order.signature);
    await tx2.wait();
    // log etherscan url
    console.log(`https://sepolia.etherscan.io/tx/${tx2.hash}`);
}

main();
