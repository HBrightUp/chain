
// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";
import "@openzeppelin/contracts/utils/math/SafeMath.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "../libraries/commonlib.sol";

contract TokenFAPE is ERC20, Ownable {

    using SafeMath for uint256;

    //timestamp of lock funds address
    uint256 public fundsLockTime;

    //information with all users
    EnumerableSet.AddressSet private _users;


    constructor( ) ERC20("Fight APE", "FAPE") {

        //FAPE distribute mechanism
        _mint(Commonlib.apeAirdrop, Commonlib.totalSupply.mul(10).div(100));
        _mint(Commonlib.gameRewardsFunds, Commonlib.totalSupply.mul(50).div(100));
        _mint(Commonlib.funds, Commonlib.totalSupply.mul(23).div(100));
        _mint(Commonlib.stakeMint, Commonlib.totalSupply.mul(17).div(100));
        fundsLockTime = block.timestamp.add( 3600 * 24 * 365);
    }

    /*
    * lock FAPE coin
    */
    function _beforeTokenTransfer(
        address from,
        address to,
        uint256 amount
    ) internal override virtual {

        //lock one year with funds address
        if(from == Commonlib.funds) {
            require(block.timestamp >= fundsLockTime, "funds on locking.");
        }
    }
    
    /*
    * airdrop to some users on furture
    */
    function airTransfer(address[] memory recipients_, uint[] memory values_)  public returns (bool) {
        require(recipients_.length > 0, "no address input.");
        require(recipients_.length == values_.length, "input not match.");

        for(uint i = 0; i < recipients_.length; i++){
            transfer(recipients_[i], values_[i]);
        }

        return true;
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

    
    /*
    * get list of holders.
    */
    function getAllAccount() public view  returns (address[] memory) {
        return EnumerableSet.values(_users);
    }
}