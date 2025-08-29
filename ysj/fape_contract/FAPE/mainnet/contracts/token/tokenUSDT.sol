
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract TokenUSDT  is ERC20, Ownable {

    using SafeMath for uint256;

    address private _govermentAddr;

    constructor( ) ERC20("usdt coin", "USDT") {
        _mint(msg.sender, 1e26);
    }  

    function mintToken() external onlyOwner {
         _mint(msg.sender, 1e26);
    }
}