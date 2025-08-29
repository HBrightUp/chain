# CMake generated Testfile for 
# Source directory: /home/uos/UOS/contracts/uosio.union
# Build directory: /home/uos/UOS/build/contracts/uosio.union
# 
# This file includes the relevant testing commands required for 
# testing this directory and lists subdirectories to be tested as well.
add_test(validate_uosio.union_abi "/home/uos/UOS/build/scripts/abi_is_json.py" "/home/uos/UOS/contracts/uosio.union/uosio.union.abi")
