add_contract(uosio.bios uosio.bios ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.bios.cpp)

target_include_directories(uosio.bios
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.bios
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.bios.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.bios.contracts.md @ONLY )

target_compile_options( uosio.bios PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
