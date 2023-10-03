// SPDX-License-Identifier: MIT
pragma solidity >=0.8.2 <0.9.0;

import "./BridgeReceiver.sol";

contract BridgeReceiverExample is BridgeReciever{

    string public message ;

    function sendMessage(string calldata _message) external{
        message = _message;
    }

}