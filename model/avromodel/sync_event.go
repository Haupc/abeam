package avromodel

import (
	"abeam/libs/common"
	"abeam/libs/model"
	"fmt"
	"math/big"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	syncEventSchema = `{
		"type": "record",
		"name": "syncEvent",
		"fields": [
			{ "name": "blockHash", "type": "string" },
			{ "name": "blockNumber", "type": "long" },
			{ "name": "blockTime", "type": "long" },
			{ "name": "logIndex", "type": "int" },
			{ "name": "txHash", "type": "string" },
			{ "name": "poolAddress", "type": "string" },
			{ "name": "token0Reserve", "type": "string" },
			{ "name": "token1Reserve", "type": "string" }
		]
	}`
)

type SyncEvent struct {
	BlockHash     string `avro:"blockHash"`
	BlockNumber   int64  `avro:"blockNumber"`
	BlockTime     int64  `avro:"blockTime"`
	LogIndex      int32  `avro:"logIndex"`
	TxHash        string `avro:"txHash"`
	PoolAddress   string `avro:"poolAddress"`
	Token0Reserve string `avro:"token0Reserve"`
	Token1Reserve string `avro:"token1Reserve"`
}

type SyncModeler[U model.UniPairSyncEvent, V SyncEvent] struct {
}

func (a SyncModeler[U, V]) FromModel(u U) (V, error) {
	event := model.UniPairSyncEvent(u)
	return V(SyncEvent{
		BlockHash:     event.BlockHash.Hex(),
		BlockNumber:   int64(event.BlockNumber),
		BlockTime:     event.BlockTime,
		LogIndex:      int32(event.LogIndex),
		TxHash:        event.TxHash.Hex(),
		PoolAddress:   event.PoolAddress.Hex(),
		Token0Reserve: event.Token0Reserve.Int().String(),
		Token1Reserve: event.Token1Reserve.Int().String(),
	}), nil
}

func (s SyncModeler[U, V]) ToModel(v V) (U, error) {
	a := SyncEvent(v)
	event := model.UniPairSyncEvent{
		BaseEvent: &model.BaseEvent{
			BlockHash:   etherCommon.HexToHash(a.BlockHash),
			BlockNumber: uint64(a.BlockNumber),
			BlockTime:   a.BlockTime,
			LogIndex:    uint(a.LogIndex),
			TxHash:      etherCommon.HexToHash(a.TxHash),
		},
		PoolAddress:   etherCommon.HexToAddress(a.PoolAddress),
		Token0Reserve: &common.Int{},
		Token1Reserve: &common.Int{},
	}
	token0Reserve, ok := new(big.Int).SetString(a.Token0Reserve, 10)
	if !ok {
		return U(event), fmt.Errorf("incorrect token0Reserve format")
	}
	token1Reserve, ok := new(big.Int).SetString(a.Token1Reserve, 10)
	if !ok {
		return U(event), fmt.Errorf("incorrect token1Reserve format")
	}
	event.Token0Reserve = (*common.Int)(token0Reserve)
	event.Token1Reserve = (*common.Int)(token1Reserve)
	return U(event), nil
}
