package model

import (
	"math/big"
	"testing"

	"abeam/libs/common"

	etherCommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/suite"
)

type TypeSuite struct {
	suite.Suite
}

func (ts *TypeSuite) TestFromEthLog() {
	testData := []struct {
		log      ethtypes.Log
		expected []TransferEvent
	}{
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x1d3437e570e93581bd94b2fd8fbf202d4a65654a"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					etherCommon.HexToHash("0x9a2a7e1a0852d2c8d78723fdec417d0a49144916"),
					etherCommon.HexToHash("0x3210a8b8206bc205840c983499ab92db0f4f3c6e"),
				},
				Data:        etherCommon.FromHex("0x0000000000000000000000000000000000000000000002544ce3ec8fd5cc0000"),
				BlockNumber: 15925684,
				TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
				TxIndex:     1,
				BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
				Index:       14,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC20,
					Token:     etherCommon.HexToAddress("0x1d3437e570e93581bd94b2fd8fbf202d4a65654a"),
					From:      etherCommon.HexToAddress("0x9a2a7e1a0852d2c8d78723fdec417d0a49144916"),
					To:        etherCommon.HexToAddress("0x3210a8b8206bc205840c983499ab92db0f4f3c6e"),
					TokenID:   common.NewInt(-1),
					Amount:    intFromString("10999800000000000000000"),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x1d3437e570e93581bd94b2fd8fbf202d4a65654a"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					etherCommon.HexToHash("0x9a2a7e1a0852d2c8d78723fdec417d0a49144916"),
					etherCommon.HexToHash("0x3210a8b8206bc205840c983499ab92db0f4f3c6e"),
					etherCommon.HexToHash("0x0000000000000000000000000000000000000000000002544ce3ec8fd5cc0000"),
				},
				Data:        []byte{},
				BlockNumber: 15925684,
				TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
				TxIndex:     1,
				BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
				Index:       14,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC721,
					Token:     etherCommon.HexToAddress("0x1d3437e570e93581bd94b2fd8fbf202d4a65654a"),
					From:      etherCommon.HexToAddress("0x9a2a7e1a0852d2c8d78723fdec417d0a49144916"),
					To:        etherCommon.HexToAddress("0x3210a8b8206bc205840c983499ab92db0f4f3c6e"),
					TokenID:   intFromString("10999800000000000000000"),
					Amount:    common.NewInt(1),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
					etherCommon.HexToHash("0x000000000000000000000000f17f5f37ff0e0152bb1d3977d5fd34da849becad"),
					etherCommon.HexToHash("0x0000000000000000000000006930940512ecf85a3a20409ac9867c39b007f0d6"),
					etherCommon.HexToHash("0x0000000000000000000000001c94b8cba96c101c7c95af4ea73017d125ffcd37"),
				},
				Data:        etherCommon.FromHex("0x6930940512ecf85a3a20409ac9867c39b007f0d60000000000000e00000000010000000000000000000000000000000000000000000000000000000000000001"),
				BlockNumber: 15925684,
				TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
				TxIndex:     1,
				BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
				Index:       14,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
					From:      etherCommon.HexToAddress("0x6930940512ecf85a3a20409ac9867c39b007f0d6"),
					To:        etherCommon.HexToAddress("0x1c94b8cba96c101c7c95af4ea73017d125ffcd37"),
					TokenID:   intFromString("47578679355604496672128132373117134144055957312424508622042117721445355225089"),
					Amount:    common.NewInt(1),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
					etherCommon.HexToHash("0x37b94141bca7000241b87b4b361f155197181002"),
					etherCommon.HexToHash("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					etherCommon.HexToHash("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
				},
				Data:        etherCommon.FromHex("0x000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000001800000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000002ee00000000000000000000000000000000000000000000000000000000000002ef00000000000000000000000000000000000000000000000000000000000002f000000000000000000000000000000000000000000000000000000000000002f400000000000000000000000000000000000000000000000000000000000002da00000000000000000000000000000000000000000000000000000000000002f100000000000000000000000000000000000000000000000000000000000002db00000000000000000000000000000000000000000000000000000000000002f500000000000000000000000000000000000000000000000000000000000002d9000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000009000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001000000000000000000000000000000000000000000000000000000000000000100000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000001"),
				BlockNumber: 15925684,
				TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
				TxIndex:     1,
				BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
				Index:       14,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("255211775190703847597530955573826158592000"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("255552057557624786060994330181257926803456"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("255892339924545724524457704788689695014912"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("257253469392229478378311203218416767860736"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("248406127852285078328263463425190794362880"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("256232622291466662987921079396121463226368"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("248746410219206016791726838032622562574336"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("257593751759150416841774577825848536072192"),
					Amount:    common.NewInt(1),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x856cab2d76c5ba117f1acf0d69bcc3efe1f44dd01e4854bbe380eededa241223"),
						BlockNumber: 15925684,
						LogIndex:    14,
						TxHash:      etherCommon.HexToHash("0x33be7e64cdafc1cb3b0a81ebe7ddfba5dd72457e3d3e55097d39640707584732"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0xc36cf0cfcb5d905b8b513860db0cfe63f6cf9f5c"),
					From:      etherCommon.HexToAddress("0x381e840f4ebe33d0153e9a312105554594a98c42"),
					To:        etherCommon.HexToAddress("0xfff9254c4db3b56a9bb999c6ab7955aca6d33ad1"),
					TokenID:   intFromString("248065845485364139864800088817759026151424"),
					Amount:    common.NewInt(1),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					etherCommon.HexToHash("0x20b7807939c774a152b396beb704acac68fa10fe"),
					etherCommon.HexToHash("0xdac17f958d2ee523a2206206994597c13d831ec7"),
				},
				Data:        etherCommon.Hex2Bytes("000000000000000000000000000000000000000000000000000000ad81c19a4a"),
				TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
				BlockNumber: uint64(14444741),
				BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
						BlockNumber: 14444741,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
					},
					EventType: common.TOKEN_TYPE_ERC20,
					Token:     etherCommon.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
					From:      etherCommon.HexToAddress("0x20b7807939c774a152b396beb704acac68fa10fe"),
					To:        etherCommon.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
					TokenID:   common.NewInt(-1),
					Amount:    intFromString("745206291018"),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x57f1887a8bf19b14fc0df6fd9b2acc9af147ea85"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef"),
					etherCommon.HexToHash("0xbcb8c532636f9b6b59ef417203ce60c1577bbe19"),
					etherCommon.HexToHash("0xdac17f958d2ee523a2206206994597c13d831ec7"),
					etherCommon.HexToHash("0x43b23e523310ee824c4239f01d6a4f438a5fec5bfd8d9e4ef277106f9682e1d3"),
				},
				Data:        etherCommon.Hex2Bytes(""),
				TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
				BlockNumber: uint64(14444741),
				BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
						BlockNumber: 14444741,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
					},
					EventType: common.TOKEN_TYPE_ERC721,
					Token:     etherCommon.HexToAddress("0x57f1887a8bf19b14fc0df6fd9b2acc9af147ea85"),
					From:      etherCommon.HexToAddress("0xbcb8c532636f9b6b59ef417203ce60c1577bbe19"),
					To:        etherCommon.HexToAddress("0xdac17f958d2ee523a2206206994597c13d831ec7"),
					TokenID:   intFromString("0x43b23e523310ee824c4239f01d6a4f438a5fec5bfd8d9e4ef277106f9682e1d3"),
					Amount:    common.NewInt(1),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62"),
					etherCommon.HexToHash("0x00000000000000000000000005a11a432c57a4470d985a0880d329a4a7c07d65"),
					etherCommon.HexToHash("0x00000000000000000000000005a11a432c57a4470d985a0880d329a4a7c07d65"),
					etherCommon.HexToHash("0x000000000000000000000000f216ea62a2f5fce8369f347cafa9cd73a019ec5c"),
				},
				Data:        etherCommon.Hex2Bytes("05a11a432c57a4470d985a0880d329a4a7c07d650000000000004200000000010000000000000000000000000000000000000000000000000000000000000001"),
				TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
				BlockNumber: uint64(14444741),
				BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
						BlockNumber: 14444741,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
					From:      etherCommon.HexToAddress("0x05a11a432c57a4470d985a0880d329a4a7c07d65"),
					To:        etherCommon.HexToAddress("0xf216ea62a2f5fce8369f347cafa9cd73a019ec5c"),
					TokenID:   intFromString("0x5a11a432c57a4470d985a0880d329a4a7c07d65000000000000420000000001"),
					Amount:    common.NewInt(1),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb"),
					etherCommon.HexToHash("0x00000000000000000000000005a11a432c57a4470d985a0880d329a4a7c07d65"),
					etherCommon.HexToHash("0x00000000000000000000000005a11a432c57a4470d985a0880d329a4a7c07d65"),
					etherCommon.HexToHash("0x000000000000000000000000f216ea62a2f5fce8369f347cafa9cd73a019ec5c"),
				},
				Data:        etherCommon.Hex2Bytes("000000000000000000000000000000000000000000000000000000000000004000000000000000000000000000000000000000000000000000000000000000a0000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000010000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000050000000000000000000000000000000000000000000000000000000000000002"),
				TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
				BlockNumber: uint64(14444741),
				BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
						BlockNumber: 14444741,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
					From:      etherCommon.HexToAddress("0x05a11a432c57a4470d985a0880d329a4a7c07d65"),
					To:        etherCommon.HexToAddress("0xf216ea62a2f5fce8369f347cafa9cd73a019ec5c"),
					TokenID:   common.NewInt(1),
					Amount:    common.NewInt(5),
				},
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x134c3c614aec433a7c8d7e2fbf24bc57c8c983b9013aaf80cbd736f9886b1571"),
						BlockNumber: 14444741,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0xb7617fbf0851bcfc6d0bd2fe755cd4b4ba93043c2d93022bd663e6d34d449be6"),
					},
					EventType: common.TOKEN_TYPE_ERC1155,
					Token:     etherCommon.HexToAddress("0x495f947276749ce646f68ac8c248420045cb7b5e"),
					From:      etherCommon.HexToAddress("0x05a11a432c57a4470d985a0880d329a4a7c07d65"),
					To:        etherCommon.HexToAddress("0xf216ea62a2f5fce8369f347cafa9cd73a019ec5c"),
					TokenID:   (*common.Int)(new(big.Int).SetBytes([]byte{0})),
					Amount:    common.NewInt(2),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0x7fcf532c15f0a6db0bd6d0e038bea71d30d808c7d98cb3bf7268a95bf5081b65"),
					etherCommon.HexToHash("0xdad36a77358ce31324b9c51687fee6ace5ac60ef"),
				},
				Data:        etherCommon.Hex2Bytes("0000000000000000000000000000000000000000000000255185929b7721bdf5"),
				BlockNumber: 11111111,
				TxHash:      etherCommon.HexToHash("0x336d418938c28007894651bf1bc8b0081cc2d23054e4bf19f9e53064149992b2"),
				TxIndex:     0,
				BlockHash:   etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
				Index:       0,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
						BlockNumber: 11111111,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0x336d418938c28007894651bf1bc8b0081cc2d23054e4bf19f9e53064149992b2"),
					},
					EventType: common.TOKEN_TYPE_ERC20,
					Token:     etherCommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					From:      etherCommon.HexToAddress("0xdad36a77358ce31324b9c51687fee6ace5ac60ef"),
					To:        etherCommon.HexToAddress("0x0"),
					TokenID:   common.NewInt(-1),
					Amount:    intFromString("688403793212644376053"),
				},
			},
		},
		{
			log: ethtypes.Log{
				Address: etherCommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
				Topics: []etherCommon.Hash{
					etherCommon.HexToHash("0xe1fffcc4923d04b559f4d29a8bfc6cda04eb5b0d3c460751c2402c5c5cc9109c"),
					etherCommon.HexToHash("0x0b215d59861b997c99785f761bc58fe4dec8ae60"),
				},
				Data:        etherCommon.Hex2Bytes("000000000000000000000000000000000000000000000000013fbe85edc90000"),
				BlockNumber: 11111111,
				TxHash:      etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
				TxIndex:     0,
				BlockHash:   etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
				Index:       0,
				Removed:     false,
			},
			expected: []TransferEvent{
				{
					BaseEvent: &BaseEvent{
						BlockHash:   etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
						BlockNumber: 11111111,
						LogIndex:    0,
						TxHash:      etherCommon.HexToHash("0x16ada8e4a7c527b6f226a065624f6d7cb7648e263f45750d4fc40d3bac9d8b95"),
					},
					EventType: common.TOKEN_TYPE_ERC20,
					Token:     etherCommon.HexToAddress("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2"),
					From:      etherCommon.HexToAddress("0x0"),
					To:        etherCommon.HexToAddress("0x0b215d59861b997c99785f761bc58fe4dec8ae60"),
					TokenID:   common.NewInt(-1),
					Amount:    common.NewInt(90000000000000000),
				},
			},
		},
	}

	for _, td := range testData {
		event, err := EventFromEthLog[TransferEvent](td.log)
		ts.Assert().NoError(err)
		ts.Assert().EqualValues(td.expected, event)
	}
}

func intFromString(input string) *common.Int {
	bn, _ := new(big.Int).SetString(input, 0)
	return (*common.Int)(bn)
}

func TestTypeSuite(t *testing.T) {
	suite.Run(t, new(TypeSuite))
}
