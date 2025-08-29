#!/bin/bash
# expected places to find UOSIO CMAKE in the docker container, in ascending order of preference
[[ -e /usr/lib/uosio/lib/cmake/uosio/uosio-config.cmake ]] && export CMAKE_FRAMEWORK_PATH="/usr/lib/uosio"
[[ -e /opt/uosio/lib/cmake/uosio/uosio-config.cmake ]] && export CMAKE_FRAMEWORK_PATH="/opt/uosio"
[[ ! -z "$UOSIO_ROOT" && -e $UOSIO_ROOT/lib/cmake/uosio/uosio-config.cmake ]] && export CMAKE_FRAMEWORK_PATH="$UOSIO_ROOT"
# fail if we didn't find it
[[ -z "$CMAKE_FRAMEWORK_PATH" ]] && exit 1
# export variables
echo "" >> /uosio.contracts/docker/environment.Dockerfile # necessary if there is no '\n' at end of file
echo "ENV CMAKE_FRAMEWORK_PATH=$CMAKE_FRAMEWORK_PATH" >> /uosio.contracts/docker/environment.Dockerfile
echo "ENV UOSIO_ROOT=$CMAKE_FRAMEWORK_PATH" >> /uosio.contracts/docker/environment.Dockerfile