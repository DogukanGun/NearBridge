// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

import "./BridgeReceiver.sol";

contract Bridge {

    mapping(string=>address) users;
    address _owner;

    event BridgeCall(address contractAddress, string message);

    constructor(){
        _owner = msg.sender;
    }

    modifier onlyAdmin(){
        require(msg.sender == _owner);
        _;
    }

    function addUserContract(string calldata userIdentifier,address userContractAddress) external onlyAdmin{
        users[userIdentifier] = userContractAddress;
    }

    function sendMessage(string calldata userIdentifier,string calldata message) external {
        require(users[userIdentifier] != address(0),"User is not registered");
        bytes memory payload = abi.encodeWithSignature("sendMessage(string)",message);
        (bool success,) = users[userIdentifier].call(payload);
        require(success);
    }

    function callBridge(string calldata message) external{
        emit BridgeCall(msg.sender,message);
    }

    function getUserContractAddress(string calldata userIdentifier) view external onlyAdmin returns(address){
        return users[userIdentifier];
    }

}