package common

import (
	"strings"

	"abeam/libs/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

var (
	chainMap = map[string]int{
		"eth":         1,
		"cronos":      25,
		"bsc":         56,
		"polygon":     137,
		"fantom":      250,
		"xdai":        100,
		"avax":        43114,
		"solana":      101,
		"moonriver":   1285,
		"huobi-token": 256,
		"arbitrum":    42161,
		"aurora":      1313161554,
	}
)

var (
	// map chainid to address
	MulticallAddress = map[int]common.Address{
		1:          common.HexToAddress("0x5ba1e12693dc8f9c48aad8770482f4739beed696"), // eth
		56:         common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // bsc
		137:        common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // polygon
		250:        common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // fantom
		43114:      common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // avax
		42161:      common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // arbitrum
		1313161554: common.HexToAddress("0xed386Fe855C1EFf2f843B910923Dd8846E45C5A4"), // aurora
	}
)

func GetChainID(chain string) int {
	return chainMap[chain]
}

const (
	ETH_NETWORK     = "eth"
	BSC_NETWORK     = "bsc"
	POLYGON_NETWORK = "polygon"
)

const (
	MAX_BLOCK_STEP = 500
	BLOCK_DELAY    = 12
	MAX_LOG_SIZE   = 9000

	TOKEN_TYPE_ERC20   = "ERC20"
	TOKEN_TYPE_ERC721  = "ERC721"
	TOKEN_TYPE_ERC1155 = "ERC1155"
)

var (
	ERC20ABI                       abi.ABI
	ERC20_TRANSFER_EVENT_NAME      = "Transfer"
	ERC20_TRANSFER_EVENT_SIGNATURE common.Hash
	ERC20_TOPIC_FILTER             [][]common.Hash

	ERC721ABI                       abi.ABI
	ERC721_TRANSFER_EVENT_NAME      = "Transfer"
	ERC721_TRANSFER_EVENT_SIGNATURE common.Hash
	ERC721_TOPIC_FILTER             [][]common.Hash

	ERC1155ABI                              abi.ABI
	ERC1155_SINGLE_TRANSFER_EVENT_NAME      = "TransferSingle"
	ERC1155_BATCH_TRANSFER_EVENT_NAME       = "TransferBatch"
	ERC1155_SINGLE_TRANSFER_EVENT_SIGNATURE common.Hash
	ERC1155_BATCH_TRANSFER_EVENT_SIGNATURE  common.Hash
	ERC1155_TOPIC_FILTER                    [][]common.Hash

	WETHABI, _                      = contracts.WethMetaData.GetAbi()
	WETH_DEPOSIT_EVENT_SIGNATURE    = WETHABI.Events["Deposit"].ID
	WETH_WITHDRAWAL_EVENT_SIGNATURE = WETHABI.Events["Withdrawal"].ID

	UNIV2ABI, _                = contracts.UniV2PairMetaData.GetAbi()
	UNIV2_SYNC_EVENT_NAME      = "Sync"
	UNIV2_SWAP_EVENT_NAME      = "Swap"
	UNIV2_SWAP_EVENT_SIGNATURE common.Hash
	// UNIV2_SYNC_EVENT_NAME = UNIV2ABI.Events["Sync"].ID
	MULTICALLABI abi.ABI

	// offchianOracle
	OFFCHAIN_ORACLE_ABI, _      = contracts.OffchainOracleMetaData.GetAbi()
	OFFCHAIN_ORACLE_BSC_ADDRESS = common.HexToAddress(" ")
)

var (
	BNBNativeAddress     = common.HexToAddress("0xbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb")
	ETHNativeAddress     = common.HexToAddress("0xeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee")
	POLYGONNativeAddress = common.HexToAddress("0xcccccccccccccccccccccccccccccccccccccccc")
)

func init() {
	ERC20ABI, _ = abi.JSON(strings.NewReader(contracts.ERC20MetaData.ABI))
	ERC20_TRANSFER_EVENT_SIGNATURE = ERC20ABI.Events[ERC20_TRANSFER_EVENT_NAME].ID
	ERC20_TOPIC_FILTER = [][]common.Hash{{ERC20_TRANSFER_EVENT_SIGNATURE}}

	UNIV2_SWAP_EVENT_SIGNATURE = UNIV2ABI.Events[UNIV2_SWAP_EVENT_NAME].ID

	ERC721ABI, _ = abi.JSON(strings.NewReader(contracts.ERC721MetaData.ABI))
	ERC721_TRANSFER_EVENT_SIGNATURE = ERC721ABI.Events[ERC20_TRANSFER_EVENT_NAME].ID
	ERC721_TOPIC_FILTER = [][]common.Hash{{ERC721_TRANSFER_EVENT_SIGNATURE}}

	ERC1155ABI, _ = abi.JSON(strings.NewReader(contracts.ERC1155MetaData.ABI))
	ERC1155_SINGLE_TRANSFER_EVENT_SIGNATURE = ERC1155ABI.Events[ERC1155_SINGLE_TRANSFER_EVENT_NAME].ID
	ERC1155_BATCH_TRANSFER_EVENT_SIGNATURE = ERC1155ABI.Events[ERC1155_BATCH_TRANSFER_EVENT_NAME].ID
	ERC1155_TOPIC_FILTER = [][]common.Hash{{ERC1155_SINGLE_TRANSFER_EVENT_SIGNATURE, ERC1155_BATCH_TRANSFER_EVENT_SIGNATURE}}

	MULTICALLABI, _ = abi.JSON(strings.NewReader(contracts.MulticallMetaData.ABI))
}

var (
	LogTypeTransferEvents = "transferEvent"
	LogTypeSyncEvents     = "syncEvent"
	LogTypeSwapEvents     = "swapEvent"
)
