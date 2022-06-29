package model

import (
	"errors"

	"abeam/libs/common"
	"abeam/libs/contracts"

	etherCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type UniPairSyncEvent struct {
	*BaseEvent

	PoolAddress   etherCommon.Address `json:"pool_address"`
	Token0Reserve *common.Int         `json:"token0_reserve"`
	Token1Reserve *common.Int         `json:"token1_reserve"`
}

func (a *UniPairSyncEvent) FormatSyncEventFromEthLog(log ethtypes.Log) error {
	if len(log.Topics) == 0 {
		return errors.New("invalid topic length")
	}
	event := UniPairSyncEvent{
		BaseEvent: &BaseEvent{
			BlockHash:   log.BlockHash,
			BlockNumber: log.BlockNumber,
			LogIndex:    log.Index,
			TxHash:      log.TxHash,
		},
		PoolAddress: log.Address,
	}

	var syncEvent contracts.UniV2PairSync
	err := common.UNIV2ABI.UnpackIntoInterface(&syncEvent, common.UNIV2_SYNC_EVENT_NAME, log.Data)
	if err != nil {
		return err
	}
	event.Token0Reserve = (*common.Int)(syncEvent.Reserve0)
	event.Token1Reserve = (*common.Int)(syncEvent.Reserve1)

	*a = event
	return nil
}
