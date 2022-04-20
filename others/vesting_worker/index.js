const ethers = require('ethers');
const {MongoClient} = require('mongodb');
const sleep = require('sleep-promise');

require('dotenv').config();

const RPC = process.env.RPC || "https://api.ezchain.com/ext/bc/C/rpc"
const provider = new ethers.providers.JsonRpcProvider({ url: RPC, timeout: 60000 })
const MongoUri = process.env.MONGO_URI || "mongodb://root:password@127.0.0.1:27017?retryWrites=true&w=majority";


let vestingIDPromises = [];
let vestingSchedulePromises = [];

async function main() {
    const client = await MongoClient.connect(MongoUri, { useNewUrlParser: true });
    await client.connect();
    let db = client.db(process.env.MONGO_DATABASE || 'coingecko');

    let abi = [{ "type": "constructor", "stateMutability": "nonpayable", "inputs": [{ "type": "address", "name": "token_", "internalType": "address" }] }, { "type": "event", "name": "OwnershipTransferred", "inputs": [{ "type": "address", "name": "previousOwner", "internalType": "address", "indexed": true }, { "type": "address", "name": "newOwner", "internalType": "address", "indexed": true }], "anonymous": false }, { "type": "event", "name": "Released", "inputs": [{ "type": "uint256", "name": "amount", "internalType": "uint256", "indexed": false }], "anonymous": false }, { "type": "event", "name": "Revoked", "inputs": [], "anonymous": false }, { "type": "fallback", "stateMutability": "payable" }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "bytes32", "name": "", "internalType": "bytes32" }], "name": "computeNextVestingScheduleIdForHolder", "inputs": [{ "type": "address", "name": "holder", "internalType": "address" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "uint256", "name": "", "internalType": "uint256" }], "name": "computeReleasableAmount", "inputs": [{ "type": "bytes32", "name": "vestingScheduleId", "internalType": "bytes32" }] }, { "type": "function", "stateMutability": "pure", "outputs": [{ "type": "bytes32", "name": "", "internalType": "bytes32" }], "name": "computeVestingScheduleIdForAddressAndIndex", "inputs": [{ "type": "address", "name": "holder", "internalType": "address" }, { "type": "uint256", "name": "index", "internalType": "uint256" }] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "createVestingSchedule", "inputs": [{ "type": "address", "name": "_beneficiary", "internalType": "address" }, { "type": "uint256", "name": "_start", "internalType": "uint256" }, { "type": "uint256", "name": "_cliff", "internalType": "uint256" }, { "type": "uint256", "name": "_duration", "internalType": "uint256" }, { "type": "uint256", "name": "_slicePeriodSeconds", "internalType": "uint256" }, { "type": "bool", "name": "_revocable", "internalType": "bool" }, { "type": "uint256", "name": "_amount", "internalType": "uint256" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "tuple", "name": "", "internalType": "struct TokenVesting.VestingSchedule", "components": [{ "type": "bool", "name": "initialized", "internalType": "bool" }, { "type": "address", "name": "beneficiary", "internalType": "address" }, { "type": "uint256", "name": "cliff", "internalType": "uint256" }, { "type": "uint256", "name": "start", "internalType": "uint256" }, { "type": "uint256", "name": "duration", "internalType": "uint256" }, { "type": "uint256", "name": "slicePeriodSeconds", "internalType": "uint256" }, { "type": "bool", "name": "revocable", "internalType": "bool" }, { "type": "uint256", "name": "amountTotal", "internalType": "uint256" }, { "type": "uint256", "name": "released", "internalType": "uint256" }, { "type": "bool", "name": "revoked", "internalType": "bool" }, { "type": "bool", "name": "locked", "internalType": "bool" }] }], "name": "getLastVestingScheduleForHolder", "inputs": [{ "type": "address", "name": "holder", "internalType": "address" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "address", "name": "", "internalType": "address" }], "name": "getToken", "inputs": [] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "bytes32", "name": "", "internalType": "bytes32" }], "name": "getVestingIdAtIndex", "inputs": [{ "type": "uint256", "name": "index", "internalType": "uint256" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "tuple", "name": "", "internalType": "struct TokenVesting.VestingSchedule", "components": [{ "type": "bool", "name": "initialized", "internalType": "bool" }, { "type": "address", "name": "beneficiary", "internalType": "address" }, { "type": "uint256", "name": "cliff", "internalType": "uint256" }, { "type": "uint256", "name": "start", "internalType": "uint256" }, { "type": "uint256", "name": "duration", "internalType": "uint256" }, { "type": "uint256", "name": "slicePeriodSeconds", "internalType": "uint256" }, { "type": "bool", "name": "revocable", "internalType": "bool" }, { "type": "uint256", "name": "amountTotal", "internalType": "uint256" }, { "type": "uint256", "name": "released", "internalType": "uint256" }, { "type": "bool", "name": "revoked", "internalType": "bool" }, { "type": "bool", "name": "locked", "internalType": "bool" }] }], "name": "getVestingSchedule", "inputs": [{ "type": "bytes32", "name": "vestingScheduleId", "internalType": "bytes32" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "tuple", "name": "", "internalType": "struct TokenVesting.VestingSchedule", "components": [{ "type": "bool", "name": "initialized", "internalType": "bool" }, { "type": "address", "name": "beneficiary", "internalType": "address" }, { "type": "uint256", "name": "cliff", "internalType": "uint256" }, { "type": "uint256", "name": "start", "internalType": "uint256" }, { "type": "uint256", "name": "duration", "internalType": "uint256" }, { "type": "uint256", "name": "slicePeriodSeconds", "internalType": "uint256" }, { "type": "bool", "name": "revocable", "internalType": "bool" }, { "type": "uint256", "name": "amountTotal", "internalType": "uint256" }, { "type": "uint256", "name": "released", "internalType": "uint256" }, { "type": "bool", "name": "revoked", "internalType": "bool" }, { "type": "bool", "name": "locked", "internalType": "bool" }] }], "name": "getVestingScheduleByAddressAndIndex", "inputs": [{ "type": "address", "name": "holder", "internalType": "address" }, { "type": "uint256", "name": "index", "internalType": "uint256" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "uint256", "name": "", "internalType": "uint256" }], "name": "getVestingSchedulesCount", "inputs": [] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "uint256", "name": "", "internalType": "uint256" }], "name": "getVestingSchedulesCountByBeneficiary", "inputs": [{ "type": "address", "name": "_beneficiary", "internalType": "address" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "uint256", "name": "", "internalType": "uint256" }], "name": "getVestingSchedulesTotalAmount", "inputs": [] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "uint256", "name": "", "internalType": "uint256" }], "name": "getWithdrawableAmount", "inputs": [] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [{ "type": "uint256", "name": "total", "internalType": "uint256" }], "name": "investorWithdraw", "inputs": [{ "type": "bool", "name": "keepWrapped", "internalType": "bool" }] }, { "type": "function", "stateMutability": "view", "outputs": [{ "type": "address", "name": "", "internalType": "address" }], "name": "owner", "inputs": [] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "release", "inputs": [{ "type": "bytes32", "name": "vestingScheduleId", "internalType": "bytes32" }, { "type": "uint256", "name": "amount", "internalType": "uint256" }] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "renounceOwnership", "inputs": [] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "revoke", "inputs": [{ "type": "bytes32", "name": "vestingScheduleId", "internalType": "bytes32" }] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "setLock", "inputs": [{ "type": "bytes32", "name": "vestingScheduleId", "internalType": "bytes32" }, { "type": "bool", "name": "locked", "internalType": "bool" }] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "transferOwnership", "inputs": [{ "type": "address", "name": "newOwner", "internalType": "address" }] }, { "type": "function", "stateMutability": "nonpayable", "outputs": [], "name": "withdraw", "inputs": [{ "type": "uint256", "name": "amount", "internalType": "uint256" }] }, { "type": "receive", "stateMutability": "payable" }]
    let contractAddress = "0x05E4dfbB6f26E568D846C95C0C716C4338fd1C0A";

    try {
        while(true) {
            let contract = new ethers.Contract(contractAddress, abi, provider);
            let vestingCount = await contract.getVestingSchedulesCount();
            let data = {};
            let vestingIDs = [];
            for (let i = 0; i < vestingCount; i++) {
                let id = await contract.getVestingIdAtIndex(i);
                console.log("list ids: ", id)
                vestingIDs.push({index: i, vestingID: id});
            }

            let schedules = []
            for (const id of vestingIDs) {
                try {
                    let schedule = await contract.callStatic.getVestingSchedule(id.vestingID);
                    schedules.push({vestingID: id.vestingID, schedule: schedule});
                } catch (e) {
                    continue;
                }
            }

            for (const e of schedules) {
                const address = e.schedule.beneficiary;
                if (data.hasOwnProperty(address)) {
                    var k = data[address];
                    k.push(e.vestingID)
                    data[address] = k;
                } else {
                    data[address] = [e.vestingID];
                }
            }
            var k = []
            for (var key in data) {
                k.push.apply(k, data[key]);
            }
            let total = 0;
            for (var v of k) {
                try {
                    let result = await contract.computeReleasableAmount(v)
                    if (total == 0) {
                        total = result;
                    } else {
                        total = total.add(result);
                    }
                } catch (e) {
                    continue;
                }
            }
            console.log(total.div(1000000000).div(1000000000));
            await db.collection(process.env.MONGO_COLLECTION || "tokens").updateOne(
                {symbol: 'ezc', id: 'ezchain'},
                {$set: {vesting: total.div(1000000000).div(1000000000).toString()}},
                {upsert: false}
            );

            await sleep(process.env.TIME || 1000*60*15)
        }
    } catch (e) {
        console.log(e);
    }


}

main()
