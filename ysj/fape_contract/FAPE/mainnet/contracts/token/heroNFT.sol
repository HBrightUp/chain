// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import  "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";   
import "../interfaces/IHeroNFT.sol";


contract HeroNFT is ERC721URIStorage, IHeroNFT, Ownable {
    using Counters for Counters.Counter;
    
    bool public isTicketTransfer;

    bool public isStartMinit;
    
    Counters.Counter private _tokenIds;

    address public proxy;

    mapping (HeroLevel => uint256) public heroSupply;

    

    constructor() ERC721("TicketNFT", "TCK") {
        isTicketTransfer = false;
        isStartMinit = false;
    }

    modifier onlyProxy() {
        require(msg.sender == proxy, "only proxy access.");
        _;
    }

    function setProxy(address proxy_)  external onlyOwner {
        proxy = proxy_;
    }

     function _beforeTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal override virtual {
        if(!isTicketTransfer) {
            require(false, "cann't transfer ticket.");
        }
    }

    function setTransferSwitch(bool isTicketTransfer_) external  onlyOwner returns (bool) {
        if(isTicketTransfer != isTicketTransfer_) {
            isTicketTransfer = !isTicketTransfer;
        }

        return true;
    }

    function setHeroLevelSupply(HeroLevel heroLevel_, uint256 supply_) external onlyOwner returns (bool) {
        heroSupply[heroLevel_] = supply_;

        return true;
    }

    function getHeroLevelSupply(HeroLevel heroLevel_) external view override returns (uint256) {
        return heroSupply[heroLevel_];
    }

    function setIsStartMinit(bool isStartMinit_) external onlyOwner returns (bool) {
        if(isStartMinit != isStartMinit_) {
            isStartMinit = !isStartMinit;
        }

        return true;
    }

    function mintToken(address player, string memory tokenURI) external override onlyOwner returns (uint256)
    {
        uint256 newItemId = _tokenIds.current();
        _mint(player, newItemId);
        _setTokenURI(newItemId, tokenURI);

        _tokenIds.increment();
        return newItemId;
    }

    function getUserCounts() external view override  returns (uint256) {
        return _tokenIds.current();
    }
}