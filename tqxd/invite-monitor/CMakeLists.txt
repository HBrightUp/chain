cmake_minimum_required(VERSION 2.8)

project(moniter)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -std=c++17")
set(CMAKE_BUILD_TYPE "Debug")

include_directories(./include)


aux_source_directory(./src DIR_SRCS)

set(base_src src/common.cpp
	     src/db_mysql.cpp
	     src/rpc.cpp
	     src/syncer.cpp)
set(job_src src/job/base.cpp
	    src/job/job.cpp)

set(link_lib curl mysqlclient event boost_system boost_filesystem boost_program_options glog gmp)

add_executable(invite-monitor ${base_src} ${job_src} src/monitor.cpp)

target_link_libraries(invite-monitor ${link_lib})

#add_executable(arc-cli ${base_src} ${job_src} src/client.cpp)

#target_link_libraries(arc-cli ${link_lib})
