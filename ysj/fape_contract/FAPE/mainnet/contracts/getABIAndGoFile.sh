solcjs --abi --include-path ../node_modules/ --base-path .  goverment.sol  -o ./abi
abigen --abi ./abi/goverment_sol_Goverment.abi --pkg goverment --out ./goFile/goverment.go


solcjs --abi --include-path ../node_modules/ --base-path .  ./token/FightingNFT.sol  -o ./abi
abigen --abi ./abi/token_FightingNFT_sol_FightingNFT.abi --pkg FightingNFT --out ./goFile/FightingNFT.go


solcjs --abi --include-path ../node_modules/ --base-path .  ./token/StoneNFT.sol  -o ./abi
abigen --abi ./abi/token_StoneNFT_sol_StoneNFT.abi --pkg StoneNFT --out ./goFile/StoneNFT.go
