add_contract(uosio.bios uosio.bios ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.bios.cpp)
target_include_directories(uosio.bios.wasm
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.bios.wasm
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")
