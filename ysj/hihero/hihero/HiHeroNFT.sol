// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "./interfaces/IHiHeros.sol";

contract HiHeroNFT is ERC721URIStorage, Ownable, IHeroNFT {

    using SafeMath for uint256;

    using Counters for Counters.Counter;

    string constant private _extensionName = ".json";

    Counters.Counter private _tokenIds;

    //白名单
    EnumerableSet.AddressSet private _whitelist;

    //是否开启铸造
    bool public _isStartMint;

    string private _baseUri;

    //卖出此数量的NFT后将开启游戏
    uint256 private _transferTimes;

    //项目方预挖地址
    address public firstMint;


    constructor() ERC721("Fighting APE NFT", "FIGHT") {
        _isStartMint = true;
        _transferTimes = 0;
        addWhitelist(msg.sender);
        firstMint = msg.sender;
    }

    function setBaseUri(string memory baseUri_) external onlyOwner {
         _baseUri = baseUri_;
    }

    function setIsStartMint(bool isStartMint_) external onlyOwner {
        if(_isStartMint != isStartMint_) {
            _isStartMint = !_isStartMint;
        }
    }

     function _baseURI() internal view virtual override returns (string memory) {
        return _baseUri;
    }

    function addWhitelist(address account_) public onlyOwner returns (bool) {
        return EnumerableSet.add(_whitelist, account_);
    }

    function removeWhitelist(address account_) public onlyOwner returns (bool) {
        return EnumerableSet.remove(_whitelist, account_);
    }

    function at(uint256 index_) public view returns (address) {
        return EnumerableSet.at(_whitelist, index_);
    }

    function isInWhitelist(address account_) external view override returns (bool) {
        return _isInWhitelist(account_);
    }

    function mintNFT(address player_, uint256 amount_) public  onlyOwner returns (uint256) {
        require(_isStartMint, "mint NFT not start.");
        require(_isInWhitelist(player_), " require whitelist accounts.");

        uint256 newItemId_ = _tokenIds.current();
        string memory baseURI = _baseURI();

        for(uint256 i = 0; i < amount_; ++i) {
            _mint(player_, newItemId_);
            _tokenIds.increment();
            
            _setTokenURI(newItemId_, string(abi.encodePacked(baseURI, Strings.toString(newItemId_), _extensionName)));
            newItemId_ = _tokenIds.current();
        }

        return newItemId_;
    }

    function tokenURI(uint256 tokenId_) public view virtual override returns (string memory) {

         require(_exists(tokenId_), "ERC721: invalid token ID");
        string memory baseURI = _baseURI();

        return bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI, Strings.toString(tokenId_),_extensionName)) : "";
    }

    function getUserCounts() external view returns (uint256) {
        return _transferTimes;
    }

    function withdraw() public onlyOwner {
        uint balance = address(this).balance;
        payable(msg.sender).transfer(balance);
    }

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual override {
        if(firstMint == from) {
            _transferTimes = _transferTimes.add(1);
        }
    }

    function _isInWhitelist(address account_) internal view returns (bool) {
        return EnumerableSet.contains(_whitelist, account_);
    }

    function getHeroAmount() external view virtual  override returns (uint256) {
        return _tokenIds.current();
    }
}