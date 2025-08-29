add_contract(uosio.bvpay uosio.bvpay ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.bvpay.cpp)

target_include_directories(uosio.bvpay
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include
   ${CMAKE_CURRENT_SOURCE_DIR}/../uosio.token/include
   ${CMAKE_CURRENT_SOURCE_DIR}/../uosio.system/include)

set_target_properties(uosio.bvpay
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.bvpay.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.bvpay.contracts.md @ONLY )

target_compile_options( uosio.bvpay PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
