
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;



import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";

library Commonlib {

    //total supply of token 
    uint256 public constant supply = 398e25;

    // isssue rate of unlock
    uint256 public constant issueRate = 60; 

    //endtime of issue
    uint256 public constant issueEndtime = issueRate * 20;

    //distributed address
    address public constant holder1 = 0x4BD5cA2f03C788F1068168D7741481BBBb19219E;
    address public constant holder2 = 0x965D98061E99A02E821B45E20B32963B3cd1ef12;
    address public constant holder3 = 0x7a52B23903f25Fc56118b9D6E1386c282594c440;
    address public constant holder4 = 0xD91576b60a14bd88880cFB78BB7BC4D1a486570a;
    address public constant holder5 = 0x33C67Cfd627b82494b42Ee52Ae7cF29D63541B55;
    address public constant holder6 = 0x8Ac115d1A4938736D593368C8334F55C88D9cAD5;
    address public constant holder7 = 0x852Ec9A172C675d94E1ffA4A3AAE131e7CdE4E5E;
    address public constant holder8 = 0xD74b84b943aCDC61bC8cDDB47Ac5Cb096E43D27b;
    
}

//token contract 
contract TokenMEETest is ERC20, Ownable {

    using SafeMath for uint256;

    uint256 public depolyTimestamp;

    mapping(address => uint256) public lockTime;

    mapping(address => uint256) public lockAmount;

    //information with all users
    EnumerableSet.AddressSet private _users;

    constructor( ) ERC20("Mee games", "MEE") {

        _init();
    }

    function _init() internal {
        uint256 current_ = block.timestamp;
        depolyTimestamp = current_;

        //lock information of horder1
        uint256 holder1Amount = Commonlib.supply.mul(25).div(1000);
        _mint(Commonlib.holder1, holder1Amount);
        lockAmount[Commonlib.holder1] = holder1Amount;
        lockTime[Commonlib.holder1] = current_.add(60 * 10);

        //lock information of horder2
        uint256 holder2Amount = Commonlib.supply.mul(32).div(100);
        _mint(Commonlib.holder2, holder2Amount);
        lockAmount[Commonlib.holder2] = holder2Amount;
        lockTime[Commonlib.holder2] = current_.add(60 * 5);

        //holder3 ~ holder5
        _mint(Commonlib.holder3, Commonlib.supply.mul(15).div(100));
        _mint(Commonlib.holder4, Commonlib.supply.mul(30).div(100));
        _mint(Commonlib.holder5, Commonlib.supply.mul(25).div(1000));

        //lock information of horder6
        uint256 holder6Amount = Commonlib.supply.mul(8).div(100);
        _mint(Commonlib.holder6, holder6Amount);
        lockAmount[Commonlib.holder6] = holder6Amount;
        lockTime[Commonlib.holder6] = current_.add(60 * 2);

        //lock information of horder7
        uint256 holder7Amount = Commonlib.supply.mul(5).div(100);
        _mint(Commonlib.holder7, holder7Amount);
        lockAmount[Commonlib.holder7] = holder7Amount;
        lockTime[Commonlib.holder7] = current_.add(60 * 2);

        //lock information of horder8
        _mint(Commonlib.holder8, holder7Amount);
         lockAmount[Commonlib.holder8] = holder7Amount;
        lockTime[Commonlib.holder8] = current_.add(60 * 2);
       
    }

    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override virtual {

        if(from == address(0)) {
            return ;
        }

        require( block.timestamp >= lockTime[from], "account on locking."); 
        require(getAvailableAmount(from) >= amount, "exceed available amount.");
    }

    function getAvailableAmount(address account_)  public view returns (uint256)  {
        //transfer of non-lock account
        if(lockTime[account_] == 0) {
            return  balanceOf(account_);
        }

        if(block.timestamp <= lockTime[account_]) {
            return 0;
        }

        //exceed  time of lock
        if(block.timestamp >= lockTime[account_].add(Commonlib.issueEndtime)) {
            return balanceOf(account_);
        }

        uint256 index_ = (block.timestamp - lockTime[account_]) / Commonlib.issueRate;

        //issue finished with 20 months.
        uint256 reserve_ = lockAmount[account_].mul(20 - index_).mul(5).div(100);

        uint256 available = balanceOf(account_).sub(reserve_);
        
        return available;
    }

    /*
    *  holder information
    */
    function _afterTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override virtual {
        EnumerableSet.add(_users, from);
        EnumerableSet.add(_users, to);
    }

    /*
    * get address of all users
    */
    function getAccountByIndex(uint256 index_) public view  returns (address) {
        return EnumerableSet.at(_users, index_);
    }

    
    /*
    * get amount of holders.
    */
    function getAccountLength() public view  returns  (uint256) {
        return EnumerableSet.length(_users);
    }

}