package avromodel

import (
	"abeam/libs/common"
	"abeam/libs/model"
	"errors"
	"math/big"

	etherCommon "github.com/ethereum/go-ethereum/common"
)

const (
	transferEventSchema = `{
		"type": "record",
		"name": "transferEvent",
		"fields": [
			{ "name": "blockHash", "type": "string" },
			{ "name": "blockNumber", "type": "long" },
			{ "name": "blockTime", "type": "long" },
			{ "name": "logIndex", "type": "int" },
			{ "name": "eventType", "type": "string" },
			{ "name": "txHash", "type": "string" },
			{ "name": "token", "type": "string" },
			{ "name": "from", "type": "string" },
			{ "name": "to", "type": "string" },
			{ "name": "tokenId", "type": "string" },
			{ "name": "amount", "type": "string" },
			{ "name": "decimals", "type": "int" }
		]
	}`
)

type TransferEvent struct {
	BlockHash   string `avro:"blockHash"`
	BlockNumber int64  `avro:"blockNumber"`
	BlockTime   int64  `avro:"blockTime"`
	LogIndex    int32  `avro:"logIndex"`
	EventType   string `avro:"eventType"`
	TxHash      string `avro:"txHash"`
	Token       string `avro:"token"`
	From        string `avro:"from"`
	To          string `avro:"to"`
	TokenID     string `avro:"tokenId"`
	Amount      string `avro:"amount"`
	Decimals    int32  `avro:"decimals"`
}

type TransferModeler[U model.TransferEvent, V TransferEvent] struct {
}

func (a TransferModeler[U, V]) FromModel(u U) (V, error) {
	event := model.TransferEvent(u)
	return V(TransferEvent{
		BlockHash:   event.BlockHash.Hex(),
		BlockNumber: int64(event.BlockNumber),
		BlockTime:   event.BlockTime,
		LogIndex:    int32(event.LogIndex),
		EventType:   event.EventType,
		TxHash:      event.TxHash.Hex(),
		Token:       event.Token.Hex(),
		From:        event.From.Hex(),
		To:          event.To.Hex(),
		TokenID:     event.TokenID.Int().String(),
		Amount:      event.Amount.Int().String(),
		Decimals:    int32(event.Decimals),
	}), nil
}

func (s TransferModeler[U, V]) ToModel(v V) (U, error) {
	a := TransferEvent(v)
	event := model.TransferEvent{
		BaseEvent: &model.BaseEvent{
			BlockHash:   etherCommon.HexToHash(a.BlockHash),
			BlockNumber: uint64(a.BlockNumber),
			LogIndex:    uint(a.LogIndex),
			TxHash:      etherCommon.HexToHash(a.TxHash),
			BlockTime:   a.BlockTime,
		},
		EventType: a.EventType,
		Token:     etherCommon.HexToAddress(a.Token),
		From:      etherCommon.HexToAddress(a.From),
		To:        etherCommon.HexToAddress(a.To),
		Decimals:  uint8(a.Decimals),
	}

	tokenID, ok := new(big.Int).SetString(a.TokenID, 10)
	if !ok {
		return U(event), errors.New("incorrect tokenID string format")
	}
	amount, ok := new(big.Int).SetString(a.Amount, 10)
	if !ok {
		return U(event), errors.New("incorrect amount string format")
	}
	event.TokenID = (*common.Int)(tokenID)
	event.Amount = (*common.Int)(amount)
	return U(event), nil
}
