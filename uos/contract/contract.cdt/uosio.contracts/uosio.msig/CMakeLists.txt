add_contract(uosio.msig uosio.msig ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.msig.cpp)
target_include_directories(uosio.msig.wasm
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.msig.wasm
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")
