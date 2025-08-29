add_contract(uosio.token uosio.token ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.token.cpp)

target_include_directories(uosio.token
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.token
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.token.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.token.contracts.md @ONLY )

target_compile_options( uosio.token PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
