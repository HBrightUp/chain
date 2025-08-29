// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


contract Sam {



    uint256 private a;
    uint256 private b;
    uint256 private c;


    constructor() {
        a = 1;
        b = 2; 
        c = 3;
    }

    function getValue() external view returns (uint256, uint256, uint256) {
        return (a,b,c);
    }
}