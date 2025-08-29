// contracts/GameItem.sol
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/security/ReentrancyGuard.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/utils/Counters.sol";
import "@openzeppelin/contracts/utils/Strings.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "../interfaces/IStoneNFT.sol";
import "../interfaces/IGoverment.sol";
import "../libraries/commonlib.sol";

contract StoneNFT is ERC721URIStorage, Ownable, ReentrancyGuard {

    using SafeMath for uint256;

    using Counters for Counters.Counter;

    address constant  public tokenFAPE = 0xaA3575C28C471539e82049bD2db72063B4Fc852f;

    address constant  public blackHole = 0x000000000000000000000000000000000000dEaD;

    address constant public goverment = 0x0901aa37E5Dd4ae22191cCf89f44cA79EcF5B49A;

    string constant private _extensionName = ".json";

    Counters.Counter private _tokenIds;

    // white list
    EnumerableSet.AddressSet private _blacklist;

    bool public isStartMint;

    string private _baseUri;

    uint256 public generateFee = 50e22;

    // award rate of mint NFT
    uint256 public awardRateOfMint;

    address public dapp;

    address public gameAwardAddr;

    struct stCardInfo{
        uint256[] code;
    }

    //card number of tokenId
    mapping(uint256 => uint256) public stone;

    //all tokenId hold by account
    mapping(address => stCardInfo) private _holders;

    constructor() ERC721("Fighting Stone NFT", "FSN") {
        isStartMint = true;
        awardRateOfMint = 5;
        dapp = 0xA025d787A6D8cF3a8996D7FD415aD96c889F0C52;
        gameAwardAddr = 0x1501EfA522462D29e71a9702380EfbE4b622B2C0;
    }

    modifier onlyDapp() {
        require(msg.sender == dapp, "only dapp access.");
        _;
    }

    function setDapp(address dapp_) external onlyOwner {
         dapp = dapp_;
    }

    function setGameAwardAddr(address gameAwardAddr_) external onlyOwner {
         gameAwardAddr = gameAwardAddr_;
    }

    function setBaseUri(string memory baseUri_) external onlyOwner {
         _baseUri = baseUri_;
    }

    function setIsStartMint(bool isStartMint_) external onlyOwner {
        if(isStartMint != isStartMint_) {
            isStartMint = !isStartMint;
        }
    }


     function _baseURI() internal view virtual override returns (string memory) {
        return _baseUri;
    }

    function addBlacklist(address account_) public onlyOwner returns (bool) {
        return EnumerableSet.add(_blacklist, account_);
    }

    function removeBlacklist(address account_) public onlyOwner returns (bool) {
        return EnumerableSet.remove(_blacklist, account_);
    }

    function at(uint256 index_) public view returns (address) {
        return EnumerableSet.at(_blacklist, index_);
    }

    function isInBlacklist(address account_) external view returns (bool) {
        return _isInBlacklist(account_);
    }

    function setGeneratefee(uint256 genfee_) external {
        generateFee = genfee_;
    }

     function setRateAwardOfMintNFT(uint256 rate_) external {
        awardRateOfMint = rate_;
    }

    function mintNFT(address account_, uint256 cardNo_) public onlyDapp nonReentrant returns (uint256)  {
        require(isStartMint, "mint NFT not start.");
        _checkCardNo(cardNo_);

        SafeERC20.safeTransferFrom(IERC20(tokenFAPE), account_, blackHole, generateFee);

        address superior_ = IGoverment(goverment).getSuperior(account_);
        if(superior_ != address(0)) {
            uint256 amount_ = generateFee.mul(awardRateOfMint).div(100);
            SafeERC20.safeTransferFrom(IERC20(tokenFAPE), gameAwardAddr, superior_, amount_);
        }

        uint256 newItemId_ = _tokenIds.current();
        string memory baseURI = _baseURI();

        _mint(account_, newItemId_);
        stone[newItemId_] = cardNo_;
        _tokenIds.increment();
        
        _setTokenURI(newItemId_, string(abi.encodePacked(baseURI, Strings.toString(newItemId_), _extensionName)));
        newItemId_ = _tokenIds.current();
        
        return newItemId_;
    }

    function _checkCardNo(uint256 cardNo_) internal pure {
        uint256 base = 10;
        uint256 card_ = cardNo_;
        uint256 times_ = 0;
        uint256 tail = 0;

        while(card_ > 0) {
            tail = card_ % base;
            require(tail > 0 && tail < 4, "invalid cardNo.");

            card_ = card_ / base;
            
            if(times_++ > 6) {
                require(false, "cardNo too large.");
            }
        }

        if(times_ < 7) {
            require(false, "cardNo too small.");
        }
    }

    function tokenURI(uint256 tokenId_) public view virtual override returns (string memory) {

         require(_exists(tokenId_), "ERC721: invalid token ID");
        string memory baseURI = _baseURI();

        return bytes(baseURI).length > 0 ? string(abi.encodePacked(baseURI, Strings.toString(tokenId_),_extensionName)) : "";
    }

    function withdraw() public onlyOwner {
        uint balance = address(this).balance;
        payable(msg.sender).transfer(balance);
    }

    function _isInBlacklist(address account_) internal view returns (bool) {
        return EnumerableSet.contains(_blacklist, account_);
    }

    function _afterTokenTransfer(
        address from,
        address to,
        uint256 tokenId
    ) internal virtual override {
        if(from != address(0)) {
            _deleteTokenId(from, tokenId);
        }

        if(to != address(0)) {
           _addTokenId(to, tokenId);
        }

    }

    function _deleteTokenId(address account_, uint256 tokenId_) internal {
        uint256 len_ = _holders[account_].code.length;
        bool isFind_;

        for(uint256 i = 0; i < len_; i++) {
            if(_holders[account_].code[i] == tokenId_) {
                isFind_ = true;
                delete _holders[account_].code[i];
                break;
            }
        }

        require(isFind_, "tokenId not found.");
    }

    function _addTokenId(address account_, uint256 tokenId_) internal {
        uint256 len_ = _holders[account_].code.length;
        bool isFind_;
        bool isAdd_;

        for(uint256 i = 0; i < len_; i++) {
            if(_holders[account_].code[i] == tokenId_) {
                isFind_ = true;
                break;
            }

            if(_holders[account_].code[i] == 0) {
                isAdd_ = true;
                _holders[account_].code[i] = tokenId_;
                break;
            }
        }

        require(!isFind_, "tokenId existed.");
        if(!isAdd_) {
            _holders[account_].code.push(tokenId_);
        }
    }

    function getTokenIdByHHolders(address account_) public view returns (uint256[] memory) {
        
        uint256 len_ = _holders[account_].code.length;
        uint256 counts = 0;

        if(ownerOf(0) == account_) {
            counts++;
        }

        for(uint256 i = 0; i < len_; i++) {
            if(_holders[account_].code[i] != 0) {
                counts++;
            }
        }

        uint256[] memory arrTokenId_ = new uint256[](counts);
        counts = 0;
        if(ownerOf(0) == account_) {
            arrTokenId_[counts++] = 0;
        }

        for(uint256 i = 0; i < len_; i++) {
            if(_holders[account_].code[i] != 0) {
                arrTokenId_[counts++] = _holders[account_].code[i];
            }
        }

        return arrTokenId_;
    }

    function getCardNo(address account_) public view returns (uint256[] memory) {
       uint256[] memory tokenId_ =  getTokenIdByHHolders(account_);
       uint256[] memory cardNo_ = new uint256[](tokenId_.length);

       for(uint256 i = 0; i < tokenId_.length; i++) {
           cardNo_[i] = stone[tokenId_[i]];
       }

        return cardNo_;
    }
}