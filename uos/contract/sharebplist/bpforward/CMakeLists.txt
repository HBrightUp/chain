cmake_minimum_required(VERSION 3.5)

PROJECT(bpforward)

#set(CMAKE_BUILD_TYPE DEBUG)
set(CMAKE_BUILD_TYPE Release)

if(NOT CMAKE_BUILD_TYPE)
  set(CMAKE_BUILD_TYPE "Release")
endif()

message("CMAKE_BUILD_TYPE = ${CMAKE_BUILD_TYPE}")
IF(CMAKE_BUILD_TYPE MATCHES Debug)
  set(CMAKE_CXX_FLAGS_DEBUG "-O0")
  message("Debug build.")
ELSEIF(CMAKE_BUILD_TYPE MATCHES Release)
  set(CMAKE_CXX_FLAGS_RELEASE "-O2 -finline-limit=1000 -DNDEBUG")
  message("Release build.")
ELSE()
  message("Some other build type.")
ENDIF()


include_directories(${PROJECT_SOURCE_DIR})
include_directories(/usr/local/boost/include)


add_definitions(-std=c++11)
link_directories( /usr/local/boost/lib)


ADD_DEFINITIONS(-DBOOST_ALL_DYN_LINK)

set(CXX_FLAGS
 -g
 -DCHECK_PTHREAD_RETURN_VALUE
 -D_FILE_OFFSET_BITS=64
 -Wall
 -Wextra
 -Wsign-compare
# -Werror
# -Wconversion
 -Wunused-variable
 -Wno-unused-parameter
# -Wold-style-cast
 -Woverloaded-virtual
 -Wpointer-arith
# -Wshadow
 -Wwrite-strings
 -march=native
 # -MMD
 -std=c++11
 -lrt 

 -rdynamic
 -pthread
# -fPIC
 -ljson
 -lssl
 -lcrypto 
 -lboost_system
 -lboost_program_options

 )

 if(CMAKE_BUILD_BITS EQUAL 32)
 list(APPEND CXX_FLAGS "-m32")
endif()

string(REPLACE ";" " " CMAKE_CXX_FLAGS "${CXX_FLAGS}")


set(IMPORTLIB 
   boost_system 
   boost_log_setup
   boost_log
   boost_thread
   boost_filesystem
   libjson.a
   libcrypto.a
   libboost_program_options.so
)

set(CMAKE_CXX_FLAGS_DEBUG   "-O0" )             # 调试包不优化
set(CMAKE_CXX_FLAGS_RELEASE "-O2 -DNDEBUG " )   # release包优化

add_executable(bpforward  main.cpp)
target_link_libraries(bpforward ${IMPORTLIB})


