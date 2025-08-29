// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract Sam5 {

    receive() payable external{

    }

    function withdraw(address to) external  {
        payable(to).transfer(address(this).balance);
    }
}