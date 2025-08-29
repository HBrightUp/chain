// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


interface ITicketNFT {
     function mintToken(address player, string memory tokenURI_) external  returns (uint256);
     function getUserCounts() external view returns (uint256);
}