// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;


import "@openzeppelin/contracts/utils/structs/EnumerableSet.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract Samwell is Ownable {
    
    EnumerableSet.AddressSet  private _users;

    constructor() {

    }

    function add(address account_) external  {
        EnumerableSet.add(_users, account_);
    }

    function remove(address account_) external  {
        EnumerableSet.remove(_users, account_);
    }

    function contains(address account_) external  view returns (bool) {
        return EnumerableSet.contains(_users, account_);
    }

    function length() external view returns (uint256) {
        return EnumerableSet.length(_users);
    }

    function at(uint256 index_) external view returns (address) {
       return EnumerableSet.at(_users, index_);
    }

    function values() external view returns (address[] memory) {
        return EnumerableSet.values(_users);
    }


}