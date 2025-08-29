add_contract(uosio.system uosio.system ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.system.cpp)
#add_executable(uosio.system.wasm ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.system.cpp)
target_include_directories(uosio.system.wasm
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include
   ${CMAKE_CURRENT_SOURCE_DIR}/../uosio.token/include)

set_target_properties(uosio.system.wasm
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")
