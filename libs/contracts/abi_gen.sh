#!/bin/bash
abigen -abi Erc1155.abi -pkg contracts -out Erc1155.go -type ERC1155
abigen -abi weth.abi -pkg contracts -out weth.go -type Weth
abigen -abi UniV2Pair.abi -pkg contracts -out univ2_pair.go -type UniV2Pair
