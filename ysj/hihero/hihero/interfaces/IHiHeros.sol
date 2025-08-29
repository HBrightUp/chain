// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;


abstract contract  IHeroNFT {
     function isInWhitelist(address account_) external view virtual returns (bool);
     function getHeroAmount() external view virtual  returns (uint256);
}