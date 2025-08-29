add_contract(uosio.wrap uosio.wrap ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.wrap.cpp)

target_include_directories(uosio.wrap
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.wrap
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.wrap.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.wrap.contracts.md @ONLY )

target_compile_options( uosio.wrap PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
