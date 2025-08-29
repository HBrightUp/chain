add_contract(uosio.system uosio.system ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.system.cpp)

target_include_directories(uosio.system
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include
   ${CMAKE_CURRENT_SOURCE_DIR}/../uosio.token/include
   ${CMAKE_CURRENT_SOURCE_DIR}/../uosio.bvpay/include)

set_target_properties(uosio.system
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

add_contract(rex.results rex.results ${CMAKE_CURRENT_SOURCE_DIR}/src/rex.results.cpp)

target_include_directories(rex.results
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(rex.results
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}/.rex")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.system.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.system.contracts.md @ONLY )

target_compile_options( uosio.system PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
