add_contract(uosio.token uosio.token ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.token.cpp)
target_include_directories(uosio.token.wasm
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.token.wasm
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")
