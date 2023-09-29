// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

import "./BridgeReceiver.sol";

contract Bridge {

    mapping(string=>address) users;
    address _owner;

    constructor(){
        _owner = msg.sender;
    }

    modifier onylAdmin(){
        require(msg.sender == _owner);
        _;
    }

    function addUserContract(string calldata userIdentifier,address userContractAddress) external onylAdmin{
        users[userIdentifier] = userContractAddress;
    }

    function sendMessage(string calldata userIdentifier,string calldata message) external {
        require(users[userIdentifier] != address(0),"User is not registered");
        bytes memory payload = abi.encodeWithSignature("sendMessage(string)",message);
        (bool success,) = users[userIdentifier].call(payload);
        require(success);
    }

}