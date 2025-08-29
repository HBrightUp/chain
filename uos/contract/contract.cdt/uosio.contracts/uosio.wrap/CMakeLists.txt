add_contract(uosio.wrap uosio.wrap ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.wrap.cpp)
target_include_directories(uosio.wrap.wasm
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.wrap.wasm
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")
