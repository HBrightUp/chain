#!/usr/bin/env python3

from hdwallet import HDWallet
from hdwallet.utils import generate_entropy
from hdwallet.symbols import BTC, ETH
from typing import Optional
import json

def generateAcount():
    # Choose strength 128, 160, 192, 224 or 256
    STRENGTH: int = 160  # Default is 128
    # Choose language english, french, italian, spanish, chinese_simplified, chinese_traditional, japanese or korean
    LANGUAGE: str = "korean"  # Default is english
    # Generate new entropy hex string
    ENTROPY: str = generate_entropy(strength=STRENGTH)
    # Secret passphrase for mnemonic
    PASSPHRASE: Optional[str] = None  # "meherett"

    # Initialize Bitcoin mainnet HDWallet
    hdwallet: HDWallet = HDWallet(symbol=ETH)
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

def createAddressByPublicKey(master_pubkey, symbol, index):
    hdwallet:HDWallet = HDWallet(symbol=symbol)
# Get Bitcoin HDWallet from xpublic 
    STRICT = True
    hdwallet.from_xpublic_key(xpublic_key=master_pubkey, strict=STRICT)

# Derivation from path
# hdwallet.from_path("m/44/0/0/0/0")
# Or derivation from index
    '''hdwallet.from_index(44, hardened=False)
    hdwallet.from_index(60, hardened=False)
    hdwallet.from_index(0, hardened=False)
    hdwallet.from_index(0)
    hdwallet.from_index(0)'''

    hdwallet.from_path("m/44/1/0/0/0")

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

if __name__ == '__main__':
    master_private_key = "xprv9s21ZrQH143K2DRhQRG94gYztZzfS9oB6U7U9BY54nG23tmM76UdXL8gSp4QschTJX4ECpwNQKDprtpG3te7AnFTtuMXnAEB5UYqXLSyS7V"
    master_public_key = "xpub661MyMwAqRbcGtALCW8ymr6L1suqTcFHBUaHXoEMY3tA8xdLrYhgDKwEjybp2mqPCWoZU2KhfgwrrnDreA9uZYdVpMqMCFf6dTGY1y5FfHA"
    symbol = ETH
    createAddressByPublicKey(master_public_key, symbol, 1)