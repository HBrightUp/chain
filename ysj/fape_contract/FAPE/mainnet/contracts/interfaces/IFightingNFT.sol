// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


interface IFightingNFT {
     function isInWhitelist(address account_) external view returns (bool);
}