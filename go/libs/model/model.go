package model

import (
	"fmt"

	etherCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type BaseEvent struct {
	BlockHash   etherCommon.Hash `json:"blockHash"`
	BlockNumber uint64           `json:"blockNumber"`
	BlockTime   int64            `json:"blockTime"`
	LogIndex    uint             `json:"logIndex"`
	TxHash      etherCommon.Hash `json:"txHash"`
	Removed     bool             `json:"removed"`
}

type Event interface {
	GetBlockNumber() uint64
	GetBlockTime() int64
	GetLogIndex() uint
	SetBlockTime(int64)
	IsRemoved() bool
	SetRemoved(removed bool)
	ComparesTo(Event) int
}

func (a *BaseEvent) GetBlockNumber() uint64 {
	return a.BlockNumber
}

func (a *BaseEvent) GetBlockTime() int64 {
	return a.BlockTime
}

func (a *BaseEvent) GetLogIndex() uint {
	return a.LogIndex
}

func (a *BaseEvent) SetBlockTime(blockTime int64) {
	a.BlockTime = blockTime
}

func (a *BaseEvent) IsRemoved() bool {
	return a.Removed
}

func (a *BaseEvent) SetRemoved(removed bool) {
	a.Removed = removed
}

func (a *BaseEvent) ComparesTo(other Event) int {
	if a.GetBlockNumber() < other.GetBlockNumber() {
		return -1
	}
	if a.GetBlockNumber() > other.GetBlockNumber() {
		return 1
	}
	if a.GetLogIndex() < other.GetLogIndex() {
		return -1
	}
	if a.GetLogIndex() > other.GetLogIndex() {
		return 1
	}
	return 0
}

func EventFromEthLog[T Event](log ethtypes.Log) ([]T, error) {
	var t interface{} = *new(T)
	switch t.(type) {
	case TransferEvent:
		events, err := TransferEventFromEthLog(log)
		result := make([]T, len(events))
		for i, ev := range events {
			result[i] = (interface{})(ev).(T)
		}
		return result, err

	case UniPairSyncEvent:
		var event UniPairSyncEvent
		err := event.FormatSyncEventFromEthLog(log)
		result := make([]T, 1)
		result[0] = (interface{})(event).(T)
		return result, err

	case UniPairSwapEvent:
		var event UniPairSwapEvent
		err := event.FormatSwapEventFromLog(log)
		result := make([]T, 1)
		result[0] = (interface{})(event).(T)
		return result, err

	default:
		return nil, fmt.Errorf("invalid Event type: %T", t)
	}
}
