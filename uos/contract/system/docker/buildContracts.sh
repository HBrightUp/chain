#!/bin/bash
cd /uosio.contracts
./build.sh
cd build
tar -pczf /artifacts/contracts.tar.gz *
