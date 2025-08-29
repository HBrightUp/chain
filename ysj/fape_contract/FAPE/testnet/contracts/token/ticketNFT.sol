// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "../interfaces/ITicketNFT.sol";

contract TicketNFT is ERC721URIStorage, ITicketNFT, Ownable {
    using Counters for Counters.Counter;
    
    bool public isTicketTransfer;

    bool public isStartMinit;
    
    Counters.Counter private _tokenIds;

    address public proxy;

    

    constructor() ERC721("TicketNFT", "TCK") {
        isTicketTransfer = false;
        isStartMinit = true;
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
        
        if(!isTicketTransfer && from != (address(0))) {
            require(false, "cann't transfer ticket.");
        }
    }

    function setTransferSwitch(bool isTicketTransfer_) external  onlyOwner returns (bool) {
        if(isTicketTransfer != isTicketTransfer_) {
            isTicketTransfer = !isTicketTransfer;
        }

        return true;
    }

    function setIsStartMinit(bool isStartMinit_) external  onlyOwner returns (bool) {
        if(isStartMinit != isStartMinit_) {
            isStartMinit = !isStartMinit;
        }

        return true;
    }

    function mintToken(address player, string memory tokenURI) public override onlyProxy returns (uint256)
    {
        require(isStartMinit, "ticket mint not start.");
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