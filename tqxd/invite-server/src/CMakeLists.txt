cmake_minimum_required(VERSION 2.8)

project(utxosever)

include_directories(./)
option(RUN_TEST "run test" OFF)

set(compat_src  compat/glibc_compat.cpp
                compat/glibc_sanity.cpp
                compat/glibcxx_sanity.cpp
                compat/strnlen.cpp)

set(crypto_src  #crypto/aes.cpp
                crypto/chacha20.cpp
                crypto/hmac_sha256.cpp
                crypto/hmac_sha512.cpp
                crypto/poly1305.c
                crypto/ripemd160.cpp
                crypto/sha1.cpp
                crypto/sha256.cpp
                crypto/sha256_sse4.cpp
                crypto/sha512.cpp
                #crypto/ctaes/ctaes.c
                )

set(rpc_src rpc/protocol.cpp
            rpc/server.cpp
            rpc/ka2ch.cpp)

set(support_src support/cleanse.cpp
                support/lockedpool.cpp)

set(database_src database/db_mysql.cpp)


set(base_src    
                base58.cpp
                bech32.cpp
                fs.cpp
                hash.cpp
                httprpc.cpp
                httpserver.cpp
                netaddress.cpp
                netbase.cpp           
                random.cpp
                sync.cpp
                scheduler.cpp
                threadinterrupt.cpp
                uint256.cpp
                util.cpp
                utilstrencodings.cpp
                utiltime.cpp
                warnings.cpp
                timedata.cpp
                title.cpp)

set(main_link_lib c rt pthread event event_pthreads curl  mysqlclient boost_system boost_chrono boost_program_options boost_filesystem boost_thread ssl crypto )   
if(RUN_TEST)
    set(test_src 
            test/mk_kafkaTest.cpp)

    find_package(GTest REQUIRED) 
    set(basic_link_lib ${main_link_lib} gtest_main gtest )
    set(main_src    ${compat_src}
                    ${crypto_src}
                    ${rpc_src}
                    ${support_src}
                    ${database_src}
                    ${base_src}
                    ${test_src}
                    config.cpp
                    test.cpp)    
    enable_testing()
    add_executable(UTTEST ${main_src})
    target_link_libraries(UTTEST ${basic_link_lib})
    add_test(
        Name ${test_src}
        COMMAND ${test_src}
    )
else()
    find_package(OpenSSL)
    set(basic_link_lib ${main_link_lib})
    set(main_src    
                    config.cpp
                    ${compat_src}
                    ${crypto_src}
                    ${rpc_src}
                    ${support_src}
                    ${database_src}
                    ${warehouse_src}
                    ${base_src}
                    main.cpp)
    add_executable(data_center ${main_src})
    target_link_libraries(data_center ${basic_link_lib})
endif(RUN_TEST)
