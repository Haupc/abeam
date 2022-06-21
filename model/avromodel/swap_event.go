package avromodel

import (
	"abeam/libs/common"
	"abeam/libs/model"
	"fmt"
	"math/big"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	swapEventSchema = `{
		"type": "record",
		"name": "swapEvent",
		"fields": [
			{ "name": "blockHash", "type": "string" },
			{ "name": "blockNumber", "type": "long" },
			{ "name": "blockTime", "type": "long" },
			{ "name": "logIndex", "type": "int" },
			{ "name": "txHash", "type": "string" },
			{ "name": "sender", "type": "string" },
			{ "name": "to", "type": "string" },
			{ "name": "poolAddress", "type": "string" },
			{ "name": "amount0", "type": "string" },
			{ "name": "amount1", "type": "string" }
		]
	}`
)

type SwapEvent struct {
	BlockHash     string `avro:"blockHash"`
	BlockNumber   int64  `avro:"blockNumber"`
	BlockTime     int64  `avro:"blockTime"`
	LogIndex      int32  `avro:"logIndex"`
	TxHash        string `avro:"txHash"`
	Sender        string `avro:"sender"`
	Receiver      string `avro:"to"`
	PoolAddress   string `avro:"poolAddress"`
	Amount0Change string `avro:"amount0"`
	Amount1Change string `avro:"amount1"`
}

type SwapModeler[S model.UniPairSwapEvent, P SwapEvent] struct {
}

func (a SwapModeler[S, P]) FromModel(s S) (P, error) {
	event := model.UniPairSwapEvent(s)

	avroEvent := SwapEvent{
		BlockHash:     event.BlockHash.Hex(),
		BlockNumber:   int64(event.BlockNumber),
		BlockTime:     event.BlockTime,
		LogIndex:      int32(event.LogIndex),
		TxHash:        event.TxHash.Hex(),
		Sender:        event.Sender.Hex(),
		Receiver:      event.Receiver.Hex(),
		PoolAddress:   event.PoolAddress.Hex(),
		Amount0Change: event.Amount0Change.Int().String(),
		Amount1Change: event.Amount1Change.Int().String(),
	}

	return P(avroEvent), nil
}

func (s SwapModeler[S, P]) ToModel(p P) (S, error) {
	avroEvent := SwapEvent(p)

	event := model.UniPairSwapEvent{
		BaseEvent: &model.BaseEvent{
			BlockHash:   etherCommon.HexToHash(avroEvent.BlockHash),
			BlockNumber: uint64(avroEvent.BlockNumber),
			BlockTime:   avroEvent.BlockTime,
			LogIndex:    uint(avroEvent.LogIndex),
			TxHash:      etherCommon.HexToHash(avroEvent.TxHash),
		},
		Sender:        etherCommon.HexToAddress(avroEvent.Sender),
		Receiver:      etherCommon.HexToAddress(avroEvent.Receiver),
		PoolAddress:   etherCommon.HexToAddress(avroEvent.PoolAddress),
		Amount0Change: &common.Int{},
		Amount1Change: &common.Int{},
	}

	amount0Change, ok := new(big.Int).SetString(avroEvent.Amount0Change, 10)
	if !ok {
		return S(event), fmt.Errorf("incorrect amount0Change format")
	}
	amount1Change, ok := new(big.Int).SetString(avroEvent.Amount1Change, 10)
	if !ok {
		return S(event), fmt.Errorf("incorrect amount1Change format")
	}
	event.Amount0Change = (*common.Int)(amount0Change)
	event.Amount1Change = (*common.Int)(amount1Change)

	return S(event), nil
}
