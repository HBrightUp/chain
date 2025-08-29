// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;





contract Receiver {
    event Received(address caller, uint amount, string message);

    receive() external payable {
        emit Received(msg.sender,  msg.value, "Received is called." );
    }

    fallback() external payable {
         emit Received(msg.sender,  msg.value, "Fallback is called." );
    }

    function f1oo(string memory message_, uint result) public payable returns(uint) {
        emit Received(msg.sender,  msg.value, message_ );
        return result;
    }

    function getAddress() public view returns (address) {
        return address(this);
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }
}


contract  Caller  {
    Receiver public receiver;

    constructor() {
        receiver = new Receiver();
    }
    
    event Response(bool success, bytes data);

    function sendMethod(address payable to_) payable external {
        bool isSend = to_.send(msg.value);
        require(isSend, "send failed.");
    }

    function transferMethod(address payable to_) payable external {
        to_.transfer(msg.value);
    }

    function testCall(address payable addr_, uint result) public payable {

        //method 1:
        //(bool Success, ) = addr_.call{value: msg.value}("");

        //method 2:
        //(bool Success, ) = addr_.call{value: msg.value}("abc");

        //method 3:
        (bool Success, bytes memory data) = addr_.call{value: msg.value}(abi.encodeWithSignature("foo(string,uint256)", "call foo", result));

        emit Response(Success, data);  
    } 

    function getAddress() public view returns (address) {
        return address(this);
    }

    function getBalance() public view returns (uint) {
        return address(this).balance;
    }





}