cmake_minimum_required(VERSION 3.5)
project(uosio_contracts VERSION 1.5.2)

set(UOSIO_CDT_VERSION_MIN "1.4")
set(UOSIO_CDT_VERSION_SOFT_MAX "1.5")
#set(UOSIO_CDT_VERSION_HARD_MAX "")

find_package(uosio.cdt)

### Check the version of uosio.cdt
set(VERSION_MATCH_ERROR_MSG "")
UOSIO_CHECK_VERSION(VERSION_OUTPUT "${UOSIO_CDT_VERSION}"
                                   "${UOSIO_CDT_VERSION_MIN}"
                                   "${UOSIO_CDT_VERSION_SOFT_MAX}"
                                   "${UOSIO_CDT_VERSION_HARD_MAX}"
                                   VERSION_MATCH_ERROR_MSG)
if(VERSION_OUTPUT STREQUAL "MATCH")
   message(STATUS "Using uosio.cdt version ${UOSIO_CDT_VERSION}")
elseif(VERSION_OUTPUT STREQUAL "WARN")
   message(WARNING "Using uosio.cdt version ${UOSIO_CDT_VERSION} even though it exceeds the maximum supported version of ${UOSIO_CDT_VERSION_SOFT_MAX}; continuing with configuration, however build may fail.\nIt is recommended to use uosio.cdt version ${UOSIO_CDT_VERSION_SOFT_MAX}.x")
else() # INVALID OR MISMATCH
   message(FATAL_ERROR "Found uosio.cdt version ${UOSIO_CDT_VERSION} but it does not satisfy version requirements: ${VERSION_MATCH_ERROR_MSG}\nPlease use uosio.cdt version ${UOSIO_CDT_VERSION_SOFT_MAX}.x")
endif(VERSION_OUTPUT STREQUAL "MATCH")

if(CMAKE_BUILD_TYPE STREQUAL "Debug")
   set(TEST_BUILD_TYPE "Debug")
   set(CMAKE_BUILD_TYPE "Release")
else()
   set(TEST_BUILD_TYPE ${CMAKE_BUILD_TYPE})
endif()


add_subdirectory(uosio.bios)
add_subdirectory(uosio.msig)
add_subdirectory(uosio.wrap)
add_subdirectory(uosio.system)
add_subdirectory(uosio.token)

if (APPLE)
   set(OPENSSL_ROOT "/usr/local/opt/openssl")
elseif (UNIX)
   set(OPENSSL_ROOT "/usr/include/openssl")
endif()
set(SECP256K1_ROOT "/usr/local")

include(UnitTestsExternalProject.txt)
