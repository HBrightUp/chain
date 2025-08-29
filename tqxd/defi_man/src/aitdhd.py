#!/usr/bin/env python3

from hdwallet import HDWallet
from hdwallet.utils import generate_entropy
from hdwallet.symbols import BTC as SYMBOL
from hdwallet.symbols import ETH as SYMBOL1

from typing import Optional

import json


def standard():
    # Choose strength 128, 160, 192, 224 or 256
    STRENGTH: int = 160  # Default is 128
    # Choose language english, french, italian, spanish, chinese_simplified, chinese_traditional, japanese or korean
    LANGUAGE: str = "korean"  # Default is english
    # Generate new entropy hex string
    ENTROPY: str = generate_entropy(strength=STRENGTH)
    # Secret passphrase for mnemonic
    PASSPHRASE: Optional[str] = None  # "meherett"

    # Initialize Bitcoin mainnet HDWallet
    hdwallet: HDWallet = HDWallet(symbol=SYMBOL, use_default_path=False)
    # Get Bitcoin HDWallet from entropy
    hdwallet.from_entropy(
    entropy=ENTROPY, language=LANGUAGE, passphrase=PASSPHRASE
    )

    # Derivation from path
    # hdwallet.from_path("m/44'/0'/0'/0/0")
    # Or derivation from index
    hdwallet.from_index(44, hardened=True)
    hdwallet.from_index(0, hardened=True)
    hdwallet.from_index(0, hardened=True)
    hdwallet.from_index(0)
    hdwallet.from_index(0)

    # Print all Bitcoin HDWallet information's
    print(json.dumps(hdwallet.dumps(), indent=4, ensure_ascii=False))

def pubkeyMaster():
    # Ethereum public key
    PUBLIC_KEY = "034f6922d19e8134de23eb98396921c02cdcf67e8c0ff23dfd955839cd557afd10"

    # Initialize Ethereum mainnet HDWallet
    hdwallet: HDWallet = HDWallet(symbol=SYMBOL)
    # Get Ethereum HDWallet from public key
    hdwallet.from_public_key(public_key=PUBLIC_KEY)

    # Print all Ethereum HDWallet information's
    # print(json.dumps(hdwallet.dumps(), indent=4, ensure_ascii=False))

    print("Cryptocurrency:", hdwallet.cryptocurrency())
    print("Symbol:", hdwallet.symbol())
    print("Network:", hdwallet.network())
    print("Uncompressed:", hdwallet.uncompressed())
    print("Compressed:", hdwallet.compressed())
    print("Public Key:", hdwallet.public_key())
    print("Finger Print:", hdwallet.finger_print())
    print("Hash:", hdwallet.hash())
    print("P2PKH Address:", hdwallet.p2pkh_address())
    print("P2SH Address:", hdwallet.p2sh_address())
    print("P2WPKH Address:", hdwallet.p2wpkh_address())
    print("P2WPKH In P2SH Address:", hdwallet.p2wpkh_in_p2sh_address())
    print("P2WSH Address:", hdwallet.p2wsh_address())
    print("P2WSH In P2SH Address:", hdwallet.p2wsh_in_p2sh_address())

def XpubMasterkey():
    # Strict for root xpublic key
    from hdwallet.utils import is_root_xpublic_key
    STRICT: bool = True
    # Bitcoin root xpublic key
    #XPUBLIC_KEY: str = "xpub661MyMwAqRbcEqD3v24ZWHGDMqqAfbDbmnUFJXfbpxGZaAshq7evA7fB75CHFbNHSot" \
    #                "LadDZw6M6ic4ZkdN6jQ2KMGR66Z2EybgdLFjNrpf"

    #PrivateKey(address=0x72eb0db81ef557b1be541fa8be46dc83da320b31, privKey=762424f0e2e7ec50939cde8d6691cadec7d959b349c18e361b808b4579275a8c)
    # 0x72eB0DB81ef557B1BE541Fa8be46Dc83Da320B31
    #XPUBLIC_KEY: str= "xpub661MyMwAqRbcG3CnMM6kMzaMbhFUF2BiuQ6nVgG1LKgZSooXtgL65v5xSfXyFGu4wRebEpoYieAY8xmajDjtW6Yb57NnGvikZnC3GK5aRpx"
    XPUBLIC_KEY: str= "xpub661MyMwAqRbcGKK9atkcKhfGvub5jwdCSkyfSoqU9bZibuKpBpv5qBDEpT3YNUyNa2t8fGHRqcvGoTYseDYvXdXrjWtfS7pTJsR9VfqP4wG"


    #0xab051d0e708caf394b35f8490fb030d51937ba7c
    # Bitcoin non-root xpublic key
    # XPUBLIC_KEY: str = "zpub6uxKjJ8pnanQKU2betFrDPVmcVUvVgyAhgWS74iaN7yUE8RADoRRnztyVEQtnzi9Fh1Vp" \
    #                    "6iJ8RT6mMqjGnS6AxGjud3P2DLzpMHUw2zT1n2"

    if STRICT:
    # Check root xpublic key
        assert is_root_xpublic_key(xpublic_key=XPUBLIC_KEY, symbol=SYMBOL1), "Invalid Root XPublic Key."

    # Initialize Bitcoin mainnet HDWallet
    hdwallet: HDWallet = HDWallet(symbol=SYMBOL1)
    # Get Bitcoin HDWallet from xpublic key
    hdwallet.from_xpublic_key(xpublic_key=XPUBLIC_KEY, strict=STRICT)

    # Derivation from path
    hdwallet.from_path("m/1588311")
    # Or derivation from index
    #hdwallet.from_index(0, hardened=False)
    #hdwallet.from_index(0, hardened=False)
    #hdwallet.from_index(0, hardened=False)
    #hdwallet.from_index(0)
    #dwallet.from_index(0)

    # Print all Bitcoin HDWallet information's
    # print(json.dumps(hdwallet.dumps(), indent=4, ensure_ascii=False))

    print("Cryptocurrency:", hdwallet.cryptocurrency())
    print("Symbol:", hdwallet.symbol())
    print("Network:", hdwallet.network())
    print("Root XPublic Key:", hdwallet.root_xpublic_key())
    print("XPublic Key:", hdwallet.xpublic_key())
    print("Uncompressed:", hdwallet.uncompressed())
    print("Compressed:", hdwallet.compressed())
    print("Chain Code:", hdwallet.chain_code())
    print("Public Key:", hdwallet.public_key())
    print("Finger Print:", hdwallet.finger_print())
    print("Semantic:", hdwallet.semantic())
    print("Path:", hdwallet.path())
    print("Hash:", hdwallet.hash())
    print("P2PKH Address:", hdwallet.p2pkh_address())
    print("P2SH Address:", hdwallet.p2sh_address())
    print("P2WPKH Address:", hdwallet.p2wpkh_address())
    print("P2WPKH In P2SH Address:", hdwallet.p2wpkh_in_p2sh_address())
    print("P2WSH Address:", hdwallet.p2wsh_address())
    print("P2WSH In P2SH Address:", hdwallet.p2wsh_in_p2sh_address())

def XprivMasterkey():
    # Strict for root xpublic key
    from hdwallet.utils import is_root_xprivate_key
    STRICT: bool = True
    # Bitcoin root xprivate key
    XPRIVATE_KEY: str = "xprv9s21ZrQH143K3Z8KFKZjzrdd3fQyqZTsYBBBhHrPmz9aa1UPM91qY7mUbQ14WuZXZWsbJq2uHPDP96m6pMLbbqDUPFzSwZVMkSdVNjd2NsM"
    # Bitcoin non-root xprivate key
    # XPRIVATE_KEY: str = "yprvAMZNWbcSVmxMiVoKgQuKmemTpEz8dJs3v8hmgkRVUjncqkXsgoxyqZ8rDb" \
    #                     "eXzMqRQZEsTcB4T5iQQx7WazLyy3KiHZrdcHo6DmGAibeMxQV"

    if STRICT:
        # Check root xprivate key
        assert is_root_xprivate_key(xprivate_key=XPRIVATE_KEY, symbol=SYMBOL1), "Invalid Root XPrivate Key."

    # Initialize Bitcoin mainnet HDWallet
    hdwallet: HDWallet = HDWallet(symbol=SYMBOL1)
    # Get Bitcoin HDWallet from xprivate key
    hdwallet.from_xprivate_key(xprivate_key=XPRIVATE_KEY, strict=STRICT)

    # Derivation from path
    hdwallet.from_path("m/0")
    # Or derivation from index
    #hdwallet.from_index(44, hardened=True)
    #hdwallet.from_index(0, hardened=True)
    #hdwallet.from_index(0, hardened=True)
    #hdwallet.from_index(0)
    #hdwallet.from_index(0)

    # Print all Bitcoin HDWallet information's
    # print(json.dumps(hdwallet.dumps(), indent=4, ensure_ascii=False))
    
    print("Cryptocurrency:", hdwallet.cryptocurrency())
    print("Symbol:", hdwallet.symbol())
    print("Network:", hdwallet.network())
    print("Root XPrivate Key:", hdwallet.root_xprivate_key())
    print("Root XPublic Key:", hdwallet.root_xpublic_key())
    print("XPrivate Key:", hdwallet.xprivate_key())
    print("XPublic Key:", hdwallet.xpublic_key())
    print("Uncompressed:", hdwallet.uncompressed())
    print("Compressed:", hdwallet.compressed())
    print("Chain Code:", hdwallet.chain_code())
    print("Private Key:", hdwallet.private_key())
    print("Public Key:", hdwallet.public_key())
    print("Wallet Important Format:", hdwallet.wif())
    print("Finger Print:", hdwallet.finger_print())
    print("Semantic:", hdwallet.semantic())
    print("Path:", hdwallet.path())
    print("Hash:", hdwallet.hash())
    print("P2PKH Address:", hdwallet.p2pkh_address())
    print("P2SH Address:", hdwallet.p2sh_address())
    print("P2WPKH Address:", hdwallet.p2wpkh_address())
    print("P2WPKH In P2SH Address:", hdwallet.p2wpkh_in_p2sh_address())
    print("P2WSH Address:", hdwallet.p2wsh_address())




XpubMasterkey()
#XprivMasterkey()