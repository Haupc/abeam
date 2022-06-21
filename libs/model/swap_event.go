package model

import (
	"errors"
	"math/big"

	"abeam/libs/common"
	"abeam/libs/contracts"

	etherCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type UniPairSwapEvent struct {
	*BaseEvent

	Sender        etherCommon.Address `json:"sender"`
	Receiver      etherCommon.Address `json:"receiver"`
	PoolAddress   etherCommon.Address `json:"pool_address"`
	Amount0Change *common.Int         `json:"amount0_change"`
	Amount1Change *common.Int         `json:"amount1_change"`
}

func (u *UniPairSwapEvent) FormatSwapEventFromLog(log ethtypes.Log) error {
	if len(log.Topics) == 0 {
		return errors.New("invalid topic length")
	}
	event := UniPairSwapEvent{
		BaseEvent: &BaseEvent{
			BlockHash:   log.BlockHash,
			BlockNumber: log.BlockNumber,
			LogIndex:    log.Index,
			TxHash:      log.TxHash,
			Removed:     log.Removed,
		},
		PoolAddress: log.Address,
		Sender:      etherCommon.HexToAddress(log.Topics[1].String()),
		Receiver:    etherCommon.HexToAddress(log.Topics[2].String()),
	}

	var swapEvent contracts.UniV2PairSwap
	err := common.UNIV2ABI.UnpackIntoInterface(&swapEvent, common.UNIV2_SWAP_EVENT_NAME, log.Data)
	if err != nil {
		return err
	}

	event.Amount0Change = (*common.Int)(big.NewInt(0).Sub(swapEvent.Amount0In, swapEvent.Amount0Out))
	event.Amount1Change = (*common.Int)(big.NewInt(0).Sub(swapEvent.Amount1In, swapEvent.Amount1Out))

	*u = event
	return nil
}
