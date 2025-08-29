// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.9;


import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/access/Ownable.sol";

contract TokenUSDT is ERC20, Ownable {

    

    constructor() ERC20("usdt test", "TUSDT") {
        _mint(msg.sender,  10**20);
    }

    function decimals() public view virtual override returns (uint8) {
        return 6;
    }

    function mint() external onlyOwner {
        _mint(msg.sender,  10**20);
    }
}