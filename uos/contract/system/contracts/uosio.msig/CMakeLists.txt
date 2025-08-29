add_contract(uosio.msig uosio.msig ${CMAKE_CURRENT_SOURCE_DIR}/src/uosio.msig.cpp)

target_include_directories(uosio.msig
   PUBLIC
   ${CMAKE_CURRENT_SOURCE_DIR}/include)

set_target_properties(uosio.msig
   PROPERTIES
   RUNTIME_OUTPUT_DIRECTORY "${CMAKE_CURRENT_BINARY_DIR}")

configure_file( ${CMAKE_CURRENT_SOURCE_DIR}/ricardian/uosio.msig.contracts.md.in ${CMAKE_CURRENT_BINARY_DIR}/ricardian/uosio.msig.contracts.md @ONLY )

target_compile_options( uosio.msig PUBLIC -R${CMAKE_CURRENT_SOURCE_DIR}/ricardian -R${CMAKE_CURRENT_BINARY_DIR}/ricardian )
