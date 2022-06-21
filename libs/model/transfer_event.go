package model

import (
	"errors"
	"math/big"

	"abeam/libs/common"
	"abeam/libs/contracts"

	etherCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

type TransferEvent struct {
	*BaseEvent
	EventType string              `json:"eventType"`
	Token     etherCommon.Address `json:"token"`
	From      etherCommon.Address `json:"from"`
	To        etherCommon.Address `json:"to"`
	TokenID   *common.Int         `json:"tokenId"`
	Amount    *common.Int         `json:"amount"`
	Decimals  uint8               `json:"decimals"`
}

func TransferEventFromEthLog(log ethtypes.Log) ([]TransferEvent, error) {
	if len(log.Topics) == 0 {
		return nil, errors.New("invalid topic length")
	}

	transferEvent := TransferEvent{
		BaseEvent: &BaseEvent{
			BlockHash:   log.BlockHash,
			BlockNumber: log.BlockNumber,
			LogIndex:    log.Index,
			TxHash:      log.TxHash,
		},
		Token: log.Address,
	}
	signature := log.Topics[0]
	switch signature {
	case
		common.ERC20_TRANSFER_EVENT_SIGNATURE,
		common.ERC721_TRANSFER_EVENT_SIGNATURE:
		if len(log.Topics) == 3 {
			transferEvent.EventType = common.TOKEN_TYPE_ERC20
			transferEvent.From = etherCommon.BytesToAddress(log.Topics[1].Bytes())
			transferEvent.To = etherCommon.BytesToAddress(log.Topics[2].Bytes())
			amount := new(big.Int).SetBytes(log.Data)
			transferEvent.TokenID = common.NewInt(-1)
			transferEvent.Amount = (*common.Int)(amount)
			return []TransferEvent{transferEvent}, nil
		}
		if len(log.Topics) == 4 {
			transferEvent.EventType = common.TOKEN_TYPE_ERC721
			transferEvent.From = etherCommon.BytesToAddress(log.Topics[1].Bytes())
			transferEvent.To = etherCommon.BytesToAddress(log.Topics[2].Bytes())
			tokenID := new(big.Int).SetBytes(log.Topics[3].Bytes())
			transferEvent.TokenID = (*common.Int)(tokenID)
			transferEvent.Amount = common.NewInt(1)

			return []TransferEvent{transferEvent}, nil
		}
		return nil, errors.New("invalid topic len for TransferEvent")
	case common.ERC1155_SINGLE_TRANSFER_EVENT_SIGNATURE:
		var event contracts.ERC1155TransferSingle
		if err := common.ERC1155ABI.UnpackIntoInterface(&event, "TransferSingle", log.Data); err != nil {
			return nil, err
		}
		if len(log.Topics) == 4 {
			transferEvent.EventType = common.TOKEN_TYPE_ERC1155
			transferEvent.From = etherCommon.BytesToAddress(log.Topics[2].Bytes())
			transferEvent.To = etherCommon.BytesToAddress(log.Topics[3].Bytes())
			transferEvent.TokenID = (*common.Int)(event.Id)
			transferEvent.Amount = (*common.Int)(event.Value)
			return []TransferEvent{transferEvent}, nil
		}
		return nil, errors.New("invalid topic len for ERC1155TransferSingleEvent")
	case common.ERC1155_BATCH_TRANSFER_EVENT_SIGNATURE:
		var event contracts.ERC1155TransferBatch
		if err := common.ERC1155ABI.UnpackIntoInterface(&event, "TransferBatch", log.Data); err != nil {
			return nil, err
		}
		var transferEvents []TransferEvent
		transferEvent.EventType = common.TOKEN_TYPE_ERC1155
		transferEvent.From = etherCommon.BytesToAddress(log.Topics[2].Bytes())
		transferEvent.To = etherCommon.BytesToAddress(log.Topics[3].Bytes())
		if len(event.Ids) != len(event.Values) {
			return nil, errors.New("invalid data for ERC1155TransferBatchEvent")
		}
		if len(log.Topics) == 4 {
			for i, tokenID := range event.Ids {
				amount := event.Values[i]

				transferEvent.TokenID = (*common.Int)(tokenID)
				transferEvent.Amount = (*common.Int)(amount)
				transferEvents = append(transferEvents, transferEvent)
			}
			return transferEvents, nil
		}
		return nil, errors.New("invalid topic len for ERC1155TransferSingleEvent")
	case common.WETH_DEPOSIT_EVENT_SIGNATURE:
		if len(log.Topics) == 2 {
			transferEvent.EventType = common.TOKEN_TYPE_ERC20
			transferEvent.From = etherCommon.HexToAddress("0x0")
			transferEvent.To = etherCommon.BytesToAddress(log.Topics[1].Bytes())
			transferEvent.TokenID = common.NewInt(-1)
			transferEvent.Amount = (*common.Int)(new(big.Int).SetBytes(log.Data))
			return []TransferEvent{transferEvent}, nil
		}
		return nil, errors.New("invalid topic len for WETH deposit event")
	case common.WETH_WITHDRAWAL_EVENT_SIGNATURE:
		if len(log.Topics) == 2 {
			transferEvent.EventType = common.TOKEN_TYPE_ERC20
			transferEvent.From = etherCommon.BytesToAddress(log.Topics[1].Bytes())
			transferEvent.To = etherCommon.HexToAddress("0x0")
			transferEvent.TokenID = common.NewInt(-1)
			transferEvent.Amount = (*common.Int)(new(big.Int).SetBytes(log.Data))
			return []TransferEvent{transferEvent}, nil
		}
		return nil, errors.New("invalid topic len for WETH deposit event")
	}

	return nil, errors.New("unsupported event")
}
