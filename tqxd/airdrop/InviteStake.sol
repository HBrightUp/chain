// SPDX-License-Identifier: MIT
pragma solidity ^0.7.6;


library SafeMath {
    /**
     * @dev Returns the addition of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryAdd(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        uint256 c = a + b;
        if (c < a) return (false, 0);
        return (true, c);
    }

    /**
     * @dev Returns the substraction of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function trySub(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        if (b > a) return (false, 0);
        return (true, a - b);
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, with an overflow flag.
     *
     * _Available since v3.4._
     */
    function tryMul(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        // Gas optimization: this is cheaper than requiring 'a' not being zero, but the
        // benefit is lost if 'b' is also tested.
        // See: https://github.com/OpenZeppelin/openzeppelin-contracts/pull/522
        if (a == 0) return (true, 0);
        uint256 c = a * b;
        if (c / a != b) return (false, 0);
        return (true, c);
    }

    /**
     * @dev Returns the division of two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryDiv(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        if (b == 0) return (false, 0);
        return (true, a / b);
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers, with a division by zero flag.
     *
     * _Available since v3.4._
     */
    function tryMod(uint256 a, uint256 b) internal pure returns (bool, uint256) {
        if (b == 0) return (false, 0);
        return (true, a % b);
    }

    /**
     * @dev Returns the addition of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `+` operator.
     *
     * Requirements:
     *
     * - Addition cannot overflow.
     */
    function add(uint256 a, uint256 b) internal pure returns (uint256) {
        uint256 c = a + b;
        require(c >= a, "SafeMath: addition overflow");
        return c;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting on
     * overflow (when the result is negative).
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b <= a, "SafeMath: subtraction overflow");
        return a - b;
    }

    /**
     * @dev Returns the multiplication of two unsigned integers, reverting on
     * overflow.
     *
     * Counterpart to Solidity's `*` operator.
     *
     * Requirements:
     *
     * - Multiplication cannot overflow.
     */
    function mul(uint256 a, uint256 b) internal pure returns (uint256) {
        if (a == 0) return 0;
        uint256 c = a * b;
        require(c / a == b, "SafeMath: multiplication overflow");
        return c;
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting on
     * division by zero. The result is rounded towards zero.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath: division by zero");
        return a / b;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting when dividing by zero.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b) internal pure returns (uint256) {
        require(b > 0, "SafeMath: modulo by zero");
        return a % b;
    }

    /**
     * @dev Returns the subtraction of two unsigned integers, reverting with custom message on
     * overflow (when the result is negative).
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {trySub}.
     *
     * Counterpart to Solidity's `-` operator.
     *
     * Requirements:
     *
     * - Subtraction cannot overflow.
     */
    function sub(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b <= a, errorMessage);
        return a - b;
    }

    /**
     * @dev Returns the integer division of two unsigned integers, reverting with custom message on
     * division by zero. The result is rounded towards zero.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryDiv}.
     *
     * Counterpart to Solidity's `/` operator. Note: this function uses a
     * `revert` opcode (which leaves remaining gas untouched) while Solidity
     * uses an invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function div(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        return a / b;
    }

    /**
     * @dev Returns the remainder of dividing two unsigned integers. (unsigned integer modulo),
     * reverting with custom message when dividing by zero.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {tryMod}.
     *
     * Counterpart to Solidity's `%` operator. This function uses a `revert`
     * opcode (which leaves remaining gas untouched) while Solidity uses an
     * invalid opcode to revert (consuming all remaining gas).
     *
     * Requirements:
     *
     * - The divisor cannot be zero.
     */
    function mod(uint256 a, uint256 b, string memory errorMessage) internal pure returns (uint256) {
        require(b > 0, errorMessage);
        return a % b;
    }
}


library EnumerableSet {
    // To implement this library for multiple types with as little code
    // repetition as possible, we write it in terms of a generic Set type with
    // bytes32 values.
    // The Set implementation uses private functions, and user-facing
    // implementations (such as AddressSet) are just wrappers around the
    // underlying Set.
    // This means that we can only create new EnumerableSets for types that fit
    // in bytes32.

    struct Set {
        // Storage of set values
        bytes32[] _values;

        // Position of the value in the `values` array, plus 1 because index 0
        // means a value is not in the set.
        mapping (bytes32 => uint256) _indexes;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function _add(Set storage set, bytes32 value) private returns (bool) {
        if (!_contains(set, value)) {
            set._values.push(value);
            // The value is stored at length-1, but we add 1 to all indexes
            // and use 0 as a sentinel value
            set._indexes[value] = set._values.length;
            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function _remove(Set storage set, bytes32 value) private returns (bool) {
        // We read and store the value's index to prevent multiple reads from the same storage slot
        uint256 valueIndex = set._indexes[value];

        if (valueIndex != 0) { // Equivalent to contains(set, value)
            // To delete an element from the _values array in O(1), we swap the element to delete with the last one in
            // the array, and then remove the last element (sometimes called as 'swap and pop').
            // This modifies the order of the array, as noted in {at}.

            uint256 toDeleteIndex = valueIndex - 1;
            uint256 lastIndex = set._values.length - 1;

            // When the value to delete is the last one, the swap operation is unnecessary. However, since this occurs
            // so rarely, we still do the swap anyway to avoid the gas cost of adding an 'if' statement.

            bytes32 lastvalue = set._values[lastIndex];

            // Move the last value to the index where the value to delete is
            set._values[toDeleteIndex] = lastvalue;
            // Update the index for the moved value
            set._indexes[lastvalue] = toDeleteIndex + 1; // All indexes are 1-based

            // Delete the slot where the moved value was stored
            set._values.pop();

            // Delete the index for the deleted slot
            delete set._indexes[value];

            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function _contains(Set storage set, bytes32 value) private view returns (bool) {
        return set._indexes[value] != 0;
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function _length(Set storage set) private view returns (uint256) {
        return set._values.length;
    }

   /**
    * @dev Returns the value stored at position `index` in the set. O(1).
    *
    * Note that there are no guarantees on the ordering of values inside the
    * array, and it may change when more values are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function _at(Set storage set, uint256 index) private view returns (bytes32) {
        require(set._values.length > index, "EnumerableSet: index out of bounds");
        return set._values[index];
    }

    // Bytes32Set

    struct Bytes32Set {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _add(set._inner, value);
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(Bytes32Set storage set, bytes32 value) internal returns (bool) {
        return _remove(set._inner, value);
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(Bytes32Set storage set, bytes32 value) internal view returns (bool) {
        return _contains(set._inner, value);
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(Bytes32Set storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

   /**
    * @dev Returns the value stored at position `index` in the set. O(1).
    *
    * Note that there are no guarantees on the ordering of values inside the
    * array, and it may change when more values are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function at(Bytes32Set storage set, uint256 index) internal view returns (bytes32) {
        return _at(set._inner, index);
    }

    // AddressSet

    struct AddressSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(AddressSet storage set, address value) internal returns (bool) {
        return _add(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(AddressSet storage set, address value) internal returns (bool) {
        return _remove(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(AddressSet storage set, address value) internal view returns (bool) {
        return _contains(set._inner, bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Returns the number of values in the set. O(1).
     */
    function length(AddressSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

   /**
    * @dev Returns the value stored at position `index` in the set. O(1).
    *
    * Note that there are no guarantees on the ordering of values inside the
    * array, and it may change when more values are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function at(AddressSet storage set, uint256 index) internal view returns (address) {
        return address(uint160(uint256(_at(set._inner, index))));
    }


    // UintSet

    struct UintSet {
        Set _inner;
    }

    /**
     * @dev Add a value to a set. O(1).
     *
     * Returns true if the value was added to the set, that is if it was not
     * already present.
     */
    function add(UintSet storage set, uint256 value) internal returns (bool) {
        return _add(set._inner, bytes32(value));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the value was removed from the set, that is if it was
     * present.
     */
    function remove(UintSet storage set, uint256 value) internal returns (bool) {
        return _remove(set._inner, bytes32(value));
    }

    /**
     * @dev Returns true if the value is in the set. O(1).
     */
    function contains(UintSet storage set, uint256 value) internal view returns (bool) {
        return _contains(set._inner, bytes32(value));
    }

    /**
     * @dev Returns the number of values on the set. O(1).
     */
    function length(UintSet storage set) internal view returns (uint256) {
        return _length(set._inner);
    }

   /**
    * @dev Returns the value stored at position `index` in the set. O(1).
    *
    * Note that there are no guarantees on the ordering of values inside the
    * array, and it may change when more values are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function at(UintSet storage set, uint256 index) internal view returns (uint256) {
        return uint256(_at(set._inner, index));
    }
}


library EnumerableMap {
    // To implement this library for multiple types with as little code
    // repetition as possible, we write it in terms of a generic Map type with
    // bytes32 keys and values.
    // The Map implementation uses private functions, and user-facing
    // implementations (such as Uint256ToAddressMap) are just wrappers around
    // the underlying Map.
    // This means that we can only create new EnumerableMaps for types that fit
    // in bytes32.

    struct MapEntry {
        bytes32 _key;
        bytes32 _value;
    }

    struct Map {
        // Storage of map keys and values
        MapEntry[] _entries;

        // Position of the entry defined by a key in the `entries` array, plus 1
        // because index 0 means a key is not in the map.
        mapping (bytes32 => uint256) _indexes;
    }

    /**
     * @dev Adds a key-value pair to a map, or updates the value for an existing
     * key. O(1).
     *
     * Returns true if the key was added to the map, that is if it was not
     * already present.
     */
    function _set(Map storage map, bytes32 key, bytes32 value) private returns (bool) {
        // We read and store the key's index to prevent multiple reads from the same storage slot
        uint256 keyIndex = map._indexes[key];

        if (keyIndex == 0) { // Equivalent to !contains(map, key)
            map._entries.push(MapEntry({ _key: key, _value: value }));
            // The entry is stored at length-1, but we add 1 to all indexes
            // and use 0 as a sentinel value
            map._indexes[key] = map._entries.length;
            return true;
        } else {
            map._entries[keyIndex - 1]._value = value;
            return false;
        }
    }

    /**
     * @dev Removes a key-value pair from a map. O(1).
     *
     * Returns true if the key was removed from the map, that is if it was present.
     */
    function _remove(Map storage map, bytes32 key) private returns (bool) {
        // We read and store the key's index to prevent multiple reads from the same storage slot
        uint256 keyIndex = map._indexes[key];

        if (keyIndex != 0) { // Equivalent to contains(map, key)
            // To delete a key-value pair from the _entries array in O(1), we swap the entry to delete with the last one
            // in the array, and then remove the last entry (sometimes called as 'swap and pop').
            // This modifies the order of the array, as noted in {at}.

            uint256 toDeleteIndex = keyIndex - 1;
            uint256 lastIndex = map._entries.length - 1;

            // When the entry to delete is the last one, the swap operation is unnecessary. However, since this occurs
            // so rarely, we still do the swap anyway to avoid the gas cost of adding an 'if' statement.

            MapEntry storage lastEntry = map._entries[lastIndex];

            // Move the last entry to the index where the entry to delete is
            map._entries[toDeleteIndex] = lastEntry;
            // Update the index for the moved entry
            map._indexes[lastEntry._key] = toDeleteIndex + 1; // All indexes are 1-based

            // Delete the slot where the moved entry was stored
            map._entries.pop();

            // Delete the index for the deleted slot
            delete map._indexes[key];

            return true;
        } else {
            return false;
        }
    }

    /**
     * @dev Returns true if the key is in the map. O(1).
     */
    function _contains(Map storage map, bytes32 key) private view returns (bool) {
        return map._indexes[key] != 0;
    }

    /**
     * @dev Returns the number of key-value pairs in the map. O(1).
     */
    function _length(Map storage map) private view returns (uint256) {
        return map._entries.length;
    }

   /**
    * @dev Returns the key-value pair stored at position `index` in the map. O(1).
    *
    * Note that there are no guarantees on the ordering of entries inside the
    * array, and it may change when more entries are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function _at(Map storage map, uint256 index) private view returns (bytes32, bytes32) {
        require(map._entries.length > index, "EnumerableMap: index out of bounds");

        MapEntry storage entry = map._entries[index];
        return (entry._key, entry._value);
    }

    /**
     * @dev Tries to returns the value associated with `key`.  O(1).
     * Does not revert if `key` is not in the map.
     */
    function _tryGet(Map storage map, bytes32 key) private view returns (bool, bytes32) {
        uint256 keyIndex = map._indexes[key];
        if (keyIndex == 0) return (false, 0); // Equivalent to contains(map, key)
        return (true, map._entries[keyIndex - 1]._value); // All indexes are 1-based
    }

    /**
     * @dev Returns the value associated with `key`.  O(1).
     *
     * Requirements:
     *
     * - `key` must be in the map.
     */
    function _get(Map storage map, bytes32 key) private view returns (bytes32) {
        uint256 keyIndex = map._indexes[key];
        require(keyIndex != 0, "EnumerableMap: nonexistent key"); // Equivalent to contains(map, key)
        return map._entries[keyIndex - 1]._value; // All indexes are 1-based
    }

    /**
     * @dev Same as {_get}, with a custom error message when `key` is not in the map.
     *
     * CAUTION: This function is deprecated because it requires allocating memory for the error
     * message unnecessarily. For custom revert reasons use {_tryGet}.
     */
    function _get(Map storage map, bytes32 key, string memory errorMessage) private view returns (bytes32) {
        uint256 keyIndex = map._indexes[key];
        require(keyIndex != 0, errorMessage); // Equivalent to contains(map, key)
        return map._entries[keyIndex - 1]._value; // All indexes are 1-based
    }

    // AddressToAddressMap

    struct AddressToAddressMap {
        Map _inner;
    }

    /**
     * @dev Adds a key-value pair to a map, or updates the value for an existing
     * key. O(1).
     *
     * Returns true if the key was added to the map, that is if it was not
     * already present.
     */
    function set(AddressToAddressMap storage map, address key, address value) internal returns (bool) {
        return _set(map._inner, bytes32(uint256(uint160(key))), bytes32(uint256(uint160(value))));
    }

    /**
     * @dev Removes a value from a set. O(1).
     *
     * Returns true if the key was removed from the map, that is if it was present.
     */
    function remove(AddressToAddressMap storage map, address key) internal returns (bool) {
        return _remove(map._inner, bytes32(uint256(uint160(key))));
    }

    /**
     * @dev Returns true if the key is in the map. O(1).
     */
    function contains(AddressToAddressMap storage map, address key) internal view returns (bool) {
        return _contains(map._inner, bytes32(uint256(uint160(key))));
    }

    /**
     * @dev Returns the number of elements in the map. O(1).
     */
    function length(AddressToAddressMap storage map) internal view returns (uint256) {
        return _length(map._inner);
    }

   /**
    * @dev Returns the element stored at position `index` in the set. O(1).
    * Note that there are no guarantees on the ordering of values inside the
    * array, and it may change when more values are added or removed.
    *
    * Requirements:
    *
    * - `index` must be strictly less than {length}.
    */
    function at(AddressToAddressMap storage map, uint256 index) internal view returns (address, address) {
        (bytes32 key, bytes32 value) = _at(map._inner, index);
        return (address(uint160(uint256(key))), address(uint160(uint256(value))));
    }

    /**
     * @dev Tries to returns the value associated with `key`.  O(1).
     * Does not revert if `key` is not in the map.
     *
     * _Available since v3.4._
     */
    function tryGet(AddressToAddressMap storage map, address key) internal view returns (bool, address) {
        (bool success, bytes32 value) = _tryGet(map._inner, bytes32(uint256(uint160(key))));
        return (success, address(uint160(uint256(value))));
    }

    /**
     * @dev Returns the value associated with `key`.  O(1).
     *
     * Requirements:
     *
     * - `key` must be in the map.
     */
    function get(AddressToAddressMap storage map, address key) internal view returns (address) {
        return address(uint160(uint256(_get(map._inner, bytes32(uint256(uint160(key)))))));
    }
}


contract Context {
    function _msgSender() internal view virtual returns (address payable) {
        return msg.sender;
    }

    function _msgData() internal view virtual returns (bytes memory) {
        this; // silence state mutability warning without generating bytecode - see https://github.com/ethereum/solidity/issues/2691
        return msg.data;
    }
}


contract Ownable is Context {
    address private _owner;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);

    /**
     * @dev Initializes the contract setting the deployer as the initial owner.
     */
    constructor () {
        address msgSender = _msgSender();
        _owner = msgSender;
        emit OwnershipTransferred(address(0), msgSender);
    }

    /**
     * @dev Returns the address of the current owner.
     */
    function owner() public view virtual returns (address) {
        return _owner;
    }

    /**
     * @dev Throws if called by any account other than the owner.
     */
    modifier onlyOwner() {
        require(owner() == _msgSender(), "Ownable: caller is not the owner");
        _;
    }

    /**
     * @dev Leaves the contract without owner. It will not be possible to call
     * `onlyOwner` functions anymore. Can only be called by the current owner.
     *
     * NOTE: Renouncing ownership will leave the contract without an owner,
     * thereby removing any functionality that is only available to the owner.
     */
    function renounceOwnership() public virtual onlyOwner {
        emit OwnershipTransferred(_owner, address(0));
        _owner = address(0);
    }

    /**
     * @dev Transfers ownership of the contract to a new account (`newOwner`).
     * Can only be called by the current owner.
     */
    function transferOwnership(address newOwner) public virtual onlyOwner {
        require(newOwner != address(0), "Ownable: new owner is the zero address");
        emit OwnershipTransferred(_owner, newOwner);
        _owner = newOwner;
    }
}


contract Relationship {
    using EnumerableSet for EnumerableSet.AddressSet;
    using EnumerableMap for EnumerableMap.AddressToAddressMap;

    mapping (address => EnumerableSet.AddressSet) inviteeSet;
    EnumerableMap.AddressToAddressMap relationship;
    
    function _addInviter(address inviter) internal {
        require(inviter != address(0), "Inviter is zero");
        require(inviter != msg.sender, "Inviter is oneself");

        address invitee = msg.sender;

        // 相互邀请
        require(!inviteeSet[invitee].contains(inviter), "Mutual invitation");

        require(relationship.set(invitee, inviter), "Inviter already exists");

        require(inviteeSet[inviter].add(invitee), "Add invitation more than once");
    }

    function getInviter(address invitee) public view returns (bool, address) {
        return relationship.tryGet(invitee);
    }

    function getInviterByIndex(uint idx) public view returns (address, address) {
        return relationship.at(idx);
    }

    function inviteeLength() public view returns (uint) {
        return relationship.length();
    }

    function getInviteeListLength(address inviter) public view returns (uint256) {
        return inviteeSet[inviter].length();
    }

    function at(address inviter, uint256 index) public view returns (address) {
        require(inviter != address(0), "Inviter is the zero address");
        return inviteeSet[inviter].at(index);
    }

    function getInviteeList(
        uint beginIndex, 
        uint endIndex, 
        address inviter
        )
        public
        view
        returns(address[] memory)
    {
        require(inviter != address(0), "Inviter is the zero address");
        require(
            beginIndex < endIndex 
            && endIndex <= inviteeSet[inviter].length()
            , "out_of_range");
        
        address[] memory inviteeLisr = new address[](endIndex - beginIndex);
        for (uint i = 0; beginIndex < endIndex; ++beginIndex) {
            inviteeLisr[i] = inviteeSet[inviter].at(beginIndex);
            ++i;
        }
        return inviteeLisr;
    }

    function getInviteeListV2(
        uint beginIndex, 
        uint endIndex, 
        address inviter
        )
        public
        view
        returns(address[] memory)
    {
        require(inviter != address(0), "Inviter is the zero address");
        
        if (inviteeSet[inviter].length() == 0) {
            return new address[](0);
        }

        require(beginIndex < endIndex && 
            beginIndex < inviteeSet[inviter].length(), "out_of_range");
        
        uint len = endIndex < inviteeSet[inviter].length()
            ? endIndex : inviteeSet[inviter].length();
        address[] memory inviteeLisr = new address[](len - beginIndex);
        for (uint i = 0; beginIndex < len; ++beginIndex) {
            inviteeLisr[i] = inviteeSet[inviter].at(beginIndex);
            ++i;
        }
        return inviteeLisr;
    }
}


contract InviteStake is Relationship, Ownable {
    using SafeMath for uint;

    struct UserInfo {
        bool isInvidteeReward;
        uint stakedAmount;
        uint invitationReward;
        uint firstStakeAmount;
        uint lastBlock; // 到期块
        uint profitDebt;
    }

    bool public paused = false;

    // 空投
    address public airdropOwner;

    // 紧急赎回地址
    address public emergencyWithdrawAddress;

    uint public constant PROFIT_BASE = 1e18;

    uint public stakeUpLimit; // 单笔质押上限
    uint public stakeLowerLimit; // 单笔质押下限
    
    uint public lastRewardBlock;
    uint public expire = 28800 * 400; // 400 day
    uint public AITDPerBlock; // 每个块的aitd量
    uint public accPerShare; // 累计的收益率
    uint public totalAmount;   // 池子当前数量

    // uint public userStakeTotalAmount;

    uint public invitationRewardAccumulate;
    uint public profitAccumulate;

    uint public invitationRewardTotalAmount;
    uint public profitTotalAmount;

    mapping(address => uint) public userOfPid;
    UserInfo[] public userInfoSet;
    
    mapping (address => uint) public inviterToInviteeReward;

    constructor(
        uint _stakeLowerLimit,
        uint _stakeUpLimit,
        uint _AITDPerBlock,
        uint _invitationRewardTotalAmount,
        uint _profitTotalAmount
    )
    {
        require(
            _stakeLowerLimit > 0
            , "The minimum limit should be greater than 0"
        );
        require(
            _stakeUpLimit > _stakeLowerLimit
            , "The upper limit is less than the lower limit"
        );

        stakeLowerLimit = _stakeLowerLimit;
        stakeUpLimit = _stakeUpLimit;
        AITDPerBlock = _AITDPerBlock;
        lastRewardBlock = block.number;
        invitationRewardTotalAmount = _invitationRewardTotalAmount;
        profitTotalAmount = _profitTotalAmount;
    }

    event Deposit(address indexed user, uint amount);
    event AirdropDeposit(address indexed src, address indexed des, uint amount);
    event Withdraw(address indexed user, uint amount);
    event InviterReward(address indexed user, uint amount);
    event WithdrawProfit(address indexed user, uint amount);
    event EmergencyWithdraw(address indexed user, uint amount);
    event InviteShip(address indexed inviter, address indexed invitee);

    modifier notPause() {
        require(paused == false, "Stake has been suspended");
        _;
    }

    receive() external payable onlyOwner {}

    function  onlyOwnerWithdraw(uint _amount) public onlyOwner{
        require(_amount >  0, "Incorrect transfer amount");
        msg.sender.transfer(_amount);
    }

    function onlyOwnerBalance() public view onlyOwner returns (uint){
        return address(this).balance;
    }

    function _airdropDeposit(address addr, uint amount) private {
        require(addr !=  address(0), "Zero address");
        require(amount >  0, "Wrong number of air drops");

        uint pid = userOfPid[addr];
        // 新账户
        if (pid == 0) {

            uint profitDebt = amount.mul(accPerShare).div(PROFIT_BASE);

            userInfoSet.push(UserInfo({
                stakedAmount : amount,
                invitationReward : 0,
                isInvidteeReward : false,
                firstStakeAmount : 0,
                lastBlock :  block.number.add(expire),
                profitDebt : profitDebt
            }));
            userOfPid[addr] = userInfoSet.length;
        } else {
            UserInfo storage user = userInfoSet[pid - 1];

            if (user.stakedAmount > 0) {

                uint256 profit = user.stakedAmount.mul(accPerShare)
                    .div(PROFIT_BASE).sub(user.profitDebt);
                if (profit > 0) {

                    if (0 == profitTotalAmount) {
                        profit = 0;
                    } else if (profit > profitTotalAmount) {
                        profit = profitTotalAmount;
                    }
                    
                    profitStatistical(profit);

                    if (profit > 0) {
                        payable(addr).transfer(profit);
                        emit WithdrawProfit(addr, profit);
                    }
                }
            }
            user.stakedAmount = user.stakedAmount.add(amount);
            user.lastBlock = block.number.add(expire);

            user.profitDebt = user.stakedAmount
                .mul(accPerShare).div(PROFIT_BASE);
        }
        totalAmount = totalAmount.add(amount);
    }

    function rewardStatistical(uint reware) private {
        invitationRewardTotalAmount = invitationRewardTotalAmount.sub(
            reware
        );

        invitationRewardAccumulate = invitationRewardAccumulate.add(
            reware
        );
    }

    function profitStatistical(uint profit) private {
        profitTotalAmount = profitTotalAmount.sub(profit);

        profitAccumulate = profitAccumulate.add(profit);
    }

    function _deposit() private {
        uint pid = userOfPid[msg.sender];

        uint amount = msg.value;
        // 新账户
        if (pid == 0) {
        
            uint profitDebt = amount.mul(accPerShare).div(PROFIT_BASE);

            userInfoSet.push(UserInfo({
                isInvidteeReward : false,
                stakedAmount : amount,
                invitationReward : 0,
                firstStakeAmount : amount,
                lastBlock :  block.number.add(expire),
                profitDebt : profitDebt
            }));
            userOfPid[msg.sender] = userInfoSet.length;
        } else {

            UserInfo storage user = userInfoSet[pid - 1];

            if (user.stakedAmount > 0) {

                uint256 profit = user.stakedAmount.mul(accPerShare)
                    .div(PROFIT_BASE).sub(user.profitDebt);
                if (profit > 0) {

                    if (0 == profitTotalAmount) {
                        profit = 0;
                    } else if (profit > profitTotalAmount) {
                        profit = profitTotalAmount;
                    }
                    profitStatistical(profit);
                    if (profit > 0) {
                        msg.sender.transfer(profit);
                        emit WithdrawProfit(msg.sender, profit);
                    }  
                }
            }

            user.stakedAmount = user.stakedAmount.add(amount);
            user.lastBlock = block.number.add(expire);

            user.profitDebt = user.stakedAmount
                .mul(accPerShare).div(PROFIT_BASE);
        }
        totalAmount = totalAmount.add(amount);
    }

    // 下级质押
    function _inviteeDeposit() private {
        // 被邀请人质押
        (bool exist, address inviter) = getInviter(msg.sender);
        // 没有邀请人
        if (!exist) {
            return ;
        }
        
        require(inviter != msg.sender, "Inviter is oneself");
        
        uint inviterPid = userOfPid[inviter];

        uint inviteePid = userOfPid[msg.sender]; 
        UserInfo storage inviteeUser = userInfoSet[inviteePid - 1];

        // 上级没质押
        if (inviterPid == 0) {

            if (
                inviteeUser.isInvidteeReward 
                || inviteeUser.firstStakeAmount == 0
            )
            {
                return ;
            }

            inviterToInviteeReward[inviter] = inviterToInviteeReward[inviter]
                .add(inviteeUser.firstStakeAmount.mul(150).div(1e3));
            
            inviteeUser.isInvidteeReward = true;

            return ;
        }

        // 邀请人已经领取了邀请奖励 || 空投质押
        if (inviteeUser.isInvidteeReward || inviteeUser.firstStakeAmount == 0) {
            return ;
        }
        
        UserInfo storage inviterUser = userInfoSet[inviterPid - 1];

        if (inviterUser.stakedAmount > 0) {

            uint256 profit = inviterUser.stakedAmount.mul(accPerShare)
                .div(PROFIT_BASE).sub(inviterUser.profitDebt);
            if (profit > 0) {

                if (0 == profitTotalAmount) {
                    profit = 0;
                } else if (profit > profitTotalAmount) {
                    profit = profitTotalAmount;
                }
                profitStatistical(profit);

                if (profit > 0) {
                    payable(inviter).transfer(profit);
                    emit WithdrawProfit(inviter, profit);
                }
            }
        }
        
        // 把邀请奖励转给池子 (记录为邀请人)
        uint inviteReward = inviteeUser.firstStakeAmount.mul(150).div(1e3);

        if (0 == invitationRewardTotalAmount) {
            inviteReward = 0;
        } else if (inviteReward > invitationRewardTotalAmount) {
            inviteReward = invitationRewardTotalAmount;
        }

        rewardStatistical(inviteReward);

        inviterUser.stakedAmount = inviterUser.stakedAmount.add(inviteReward);
        inviterUser.invitationReward = inviterUser
            .invitationReward.add(inviteReward);
        // inviterUser.lastBlock = block.number.add(expire);
        
        inviteeUser.isInvidteeReward = true;
        inviteeUser.lastBlock = block.number.add(expire);

        totalAmount = totalAmount.add(inviteReward);

        inviterUser.profitDebt = inviterUser.stakedAmount
                .mul(accPerShare).div(PROFIT_BASE);

        if (inviteReward > 0) {
            emit InviterReward(inviter, inviteReward);
        }
    }

    // 上级质押, 下级也质押过
    function _inviterDeposit() private {
        uint pid = userOfPid[msg.sender];
        // 没有邀请人
        if (pid == 0) {
            return ;
        }
        UserInfo storage inviterUser = userInfoSet[pid - 1];

        uint len = getInviteeListLength(msg.sender);
        uint inviteReward = inviterToInviteeReward[msg.sender];

        if(len > 0 && inviteReward > 0) {

            if (0 == invitationRewardTotalAmount) {
                inviteReward = 0;
            } else if (inviteReward > invitationRewardTotalAmount) {
                inviteReward = invitationRewardTotalAmount;
            }

            rewardStatistical(inviteReward);

            inviterUser.stakedAmount = inviterUser.stakedAmount.add(inviteReward);
            inviterUser.invitationReward = inviterUser.invitationReward.add(inviteReward);
            inviterUser.lastBlock = block.number.add(expire);

            inviterUser.profitDebt = inviterUser.stakedAmount
                .mul(accPerShare).div(PROFIT_BASE);

            totalAmount = totalAmount.add(inviteReward);

            inviterToInviteeReward[msg.sender] = 0;

            if (inviteReward > 0) {
                emit InviterReward(msg.sender, inviteReward);
            }
        }
    }

    function _updateAccPerShare() private {
        if (block.number > lastRewardBlock && totalAmount > 0) {
            accPerShare = accPerShare.add(block.number
            .sub(lastRewardBlock).mul(AITDPerBlock)
            .mul(PROFIT_BASE).div(totalAmount));
            lastRewardBlock = block.number;
        }
    }

    function setPause() public onlyOwner {
        paused = !paused;
    }

    function setStakeLimit(uint lowerLimit, uint upLimit) public onlyOwner {
        require(lowerLimit > 0, "The minimum limit should be greater than 0");
        require(
            upLimit > lowerLimit
            , "The upper limit is less than the lower limit"
        );
       
        stakeLowerLimit = lowerLimit;
        stakeUpLimit = upLimit;
    }

    function setExpire(uint exp) public onlyOwner {
        expire = exp;
    }

    function setAITDPerBlock(uint amount) public onlyOwner {
        _updateAccPerShare();
        AITDPerBlock = amount;
    }

    function setAirdropOwner(address _airdropOwner) public onlyOwner {
        require(_airdropOwner != address(0), "Zero address");

        airdropOwner = _airdropOwner;
    }

    function setEmergencyWithdrawAddress(address addr) public onlyOwner {
        require(addr != address(0), "Zero address");

        emergencyWithdrawAddress = addr;
    }

    function setInvitationRewardTotalAmount(uint amount) public onlyOwner {
        invitationRewardTotalAmount = amount;
    }

    function setProfitTotalAmount(uint amount) public onlyOwner {
        profitTotalAmount = amount;
    }

    function userLength() public view returns (uint) {
        return userInfoSet.length;
    }

    function getProfit(address addr) view public returns(uint) {
        require(addr !=  address(0), "Zero address");
        uint pid = userOfPid[addr];
        // user 没有抵押
        if (pid == 0) {
            return 0;
        }

        UserInfo storage user = userInfoSet[pid - 1];
        if (block.number > lastRewardBlock) {
            uint perShare = block.number.sub(
                lastRewardBlock).mul(AITDPerBlock).mul(
                PROFIT_BASE).div(totalAmount);

            perShare = perShare.add(accPerShare);
            return user.stakedAmount.mul(perShare).div(PROFIT_BASE)
                .sub(user.profitDebt);
        } else {
            return user.stakedAmount.mul(accPerShare).div(PROFIT_BASE)
                .sub(user.profitDebt);
        }
    }

    function getStakedInvitation(
        address addr
        ) 
        public 
        view 
        returns (uint, uint) 
    { 
        uint pid = userOfPid[addr];
        // user 没有抵押
        if (pid == 0) {
            return (0, 0);
        }
        UserInfo storage user = userInfoSet[pid - 1];
        return (user.stakedAmount, user.invitationReward);
    }

    function getRemainderExpire(address addr) view public returns(bool, uint) {
        require(addr !=  address(0), "Zero address");

        uint pid = userOfPid[addr];
        // 没有抵押
        if (pid == 0) {
            return (false, 0);
        }
        UserInfo storage user = userInfoSet[pid - 1];
        if (user.lastBlock >  block.number) {
            return (true, user.lastBlock - block.number);
        } else {
            return (true, 0);
        }
    }

    function getUserInfo(address addr) 
        view 
        public 
        returns (bool, bool, uint, uint, uint, uint, uint) 
    {
        require(addr !=  address(0), "Zero address");

        uint pid = userOfPid[addr];
        // 没有抵押
        if (pid == 0) {
            return (false, false, 0, 0, 0, 0, 0);
        }
        UserInfo storage user = userInfoSet[pid - 1];
        return (true, user.isInvidteeReward, user.stakedAmount, 
            user.invitationReward, user.firstStakeAmount, 
            user.lastBlock, user.profitDebt);
    }

    function isDeposit(address addr) view public returns(bool) {
        require(addr !=  address(0), "Zero address");

        uint pid = userOfPid[addr];
        // 没有抵押
        if (pid == 0) {
            return false;
        }
        UserInfo storage user = userInfoSet[pid - 1];
        if (user.firstStakeAmount > 0) {
            return true;
        }
        return false;
    }

    function isExpire(address addr) view public returns(bool) {
        require(addr !=  address(0), "Zero address");

        uint pid = userOfPid[addr];
        // 没有抵押
        if (pid == 0) {
            return false;
        }
        UserInfo storage user = userInfoSet[pid - 1];
        if (block.number < user.lastBlock) {
            return false;
        }
        return true;
    }

    function addInviter(address inviter) public notPause {
        _addInviter(inviter);

        emit InviteShip(inviter, msg.sender);

        uint inviterPid = userOfPid[inviter];
        if (inviterPid == 0) {
            return ;
        }

        UserInfo storage inviterUser = userInfoSet[inviterPid - 1];
        if (inviterUser.firstStakeAmount == 0) {
            return ;
        }

        uint inviteePid = userOfPid[msg.sender]; 
        if (inviteePid == 0) {
             return ;
        }

        UserInfo storage inviteeUser = userInfoSet[inviteePid - 1];
        if (inviteeUser.isInvidteeReward || inviteeUser.firstStakeAmount == 0) {
            return ;
        }

        _updateAccPerShare();

        if (inviterUser.stakedAmount > 0) {

            uint256 profit = inviterUser.stakedAmount.mul(accPerShare)
                .div(PROFIT_BASE).sub(inviterUser.profitDebt);
            if (profit > 0) {

                if (0 == profitTotalAmount) {
                    profit = 0;
                } else if (profit > profitTotalAmount) {
                    profit = profitTotalAmount;
                }
                profitStatistical(profit);
                if (profit > 0) {
                    payable(inviter).transfer(profit);
                    emit WithdrawProfit(inviter, profit);
                }
                
            }
        }

        // 把邀请奖励转给池子 (记录为邀请人)
        uint inviteReward = inviteeUser.firstStakeAmount.mul(150).div(1e3);

        if (0 == invitationRewardTotalAmount) {
            inviteReward = 0;
        }
        else if (inviteReward > invitationRewardTotalAmount) {
            inviteReward = invitationRewardTotalAmount;
        }

        rewardStatistical(inviteReward);

        inviterUser.stakedAmount = inviterUser.stakedAmount.add(inviteReward);
        inviterUser.invitationReward = inviterUser
            .invitationReward.add(inviteReward);
        
        inviteeUser.isInvidteeReward = true;

        inviterUser.profitDebt = inviterUser.stakedAmount
                .mul(accPerShare).div(PROFIT_BASE);

        totalAmount = totalAmount.add(inviteReward);

        if (inviteReward > 0) {
            emit InviterReward(inviter, inviteReward);
        }
    }

    function deposit() public payable notPause {
        require(
            msg.value > 0 
            && msg.value >= stakeLowerLimit 
            && msg.value <= stakeUpLimit
            , "deposit amount error"
        );

        if(userInfoSet.length == 0) {
            lastRewardBlock = block.number;
        }

        _updateAccPerShare();

        _deposit();
        // 被邀请人质押
        _inviteeDeposit();
        // 邀请人后质押, 有被邀请人
        _inviterDeposit();

        emit Deposit(msg.sender, msg.value);
    }

    function withdraw(uint _amount) public notPause {
        /*
            1. 质押数量 + 收益 
            2. 收益
        */
        
        _updateAccPerShare();

        uint pid = userOfPid[msg.sender];
        // 没有抵押
        if (pid == 0) {
            return ;
        }

        UserInfo storage user =  userInfoSet[pid - 1];
        require(
            user.stakedAmount > 0 
            && user.stakedAmount >= _amount, 
            "There's not enough staked"
        );

        uint profit = user.stakedAmount
            .mul(accPerShare).div(PROFIT_BASE).sub(user.profitDebt);
        
        if (0 == profitTotalAmount) {
            profit = 0;
        } else if (profit > profitTotalAmount) {
            profit = profitTotalAmount;
        }

        profitStatistical(profit);

        if (_amount > 0) {
            // 赎回一定的质押数量 + 收益
            // 到期才可以赎回质押额
            require(block.number >= user.lastBlock, "It's still under staking");
            
            user.stakedAmount = user.stakedAmount.sub(_amount);
            totalAmount = totalAmount.sub(_amount);

            msg.sender.transfer(_amount);
            emit Withdraw(msg.sender, _amount);
            
            if (profit > 0) {
                msg.sender.transfer(profit);
                user.profitDebt = user.stakedAmount
                    .mul(accPerShare).div(PROFIT_BASE);
                emit WithdrawProfit(msg.sender, profit);
            }
        } else {
            // 赎回收益
            if (profit > 0) {
                msg.sender.transfer(profit);
                user.profitDebt = user.stakedAmount
                    .mul(accPerShare).div(PROFIT_BASE);
                emit WithdrawProfit(msg.sender, profit);
            }
        }
    }

    function emergencyWithdraw(address addr) public {
        require(
            emergencyWithdrawAddress !=  address(0), 
            "emergencyWithdrawAddress is zero address"
        );

        require(
            emergencyWithdrawAddress ==  msg.sender
            , "Insufficient permissions"
        );
        require(addr !=  address(0), "Zero address");

        uint pid = userOfPid[addr];
        if (pid == 0){
            return ;
        }
        UserInfo storage user = userInfoSet[pid - 1];
        if (user.stakedAmount > 0) {
            uint amount = user.stakedAmount;
            uint reward = user.invitationReward;

            if (amount > reward) {
                amount = amount.sub(reward);

                user.stakedAmount = 0;
                user.invitationReward = 0;
                user.profitDebt = 0;
                payable(addr).transfer(amount);

                totalAmount = totalAmount.sub(amount);
                if (reward > 0) {
                    totalAmount = totalAmount.sub(reward);
                }
                emit EmergencyWithdraw(addr, amount);

            } else if (reward > amount) {
                user.stakedAmount = 0;
                user.invitationReward = 0;
                user.profitDebt = 0;

                totalAmount = totalAmount.sub(amount);
            }
        }
    }

    function airdrop(address[] memory batch, uint amount) public notPause {
        require(airdropOwner !=  address(0), "airdropOwner is zero address");
        require(airdropOwner == msg.sender, "Insufficient permissions");
    
        if(userInfoSet.length == 0) {
            lastRewardBlock = block.number;
        }

        _updateAccPerShare();

        for(uint i = 0; i < batch.length; ++i) {        
            _airdropDeposit(batch[i], amount);

            if (amount > 0) {
                emit AirdropDeposit(msg.sender, batch[i], amount);
            }
        }
    }
}


