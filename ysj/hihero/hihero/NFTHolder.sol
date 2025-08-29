// SPDX-License-Identifier: GPL-3.0
pragma solidity ^0.8.9;

import "@openzeppelin/contracts-upgradeable/token/ERC721/IERC721Upgradeable.sol";


contract  NFTHolder {


    address public nft = 0x5771009e0987a038DAe684a481a415862b929476;


    function getHolder(uint256 index_) public view  returns (address) {
        return IERC721Upgradeable(nft).ownerOf(index_);
    }
}
