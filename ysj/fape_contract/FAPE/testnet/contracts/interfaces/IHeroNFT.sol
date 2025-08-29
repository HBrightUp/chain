// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


interface IHeroNFT {
     enum HeroLevel {plain, delicate, epic, legend}
     function mintToken(address player, string memory tokenURI) external  returns (uint256);
     function getUserCounts() external view   returns (uint256);
     function getHeroLevelSupply(HeroLevel heroLevel_) external view returns (uint256);
}