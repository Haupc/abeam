package model

import "github.com/shopspring/decimal"

type Pool struct {
	Address        string          `gorm:"primaryKey" json:"address"`
	Token0Address  string          `gorm:"index;column:token0_address" json:"token0_address"`
	Token1Address  string          `gorm:"index;column:token1_address" json:"token1_address"`
	FactoryAddress string          `gorm:"index;column:factory" json:"factory_address"`
	Token0         Token           `gorm:"-" json:"token0"`
	Token1         Token           `gorm:"-" json:"token1"`
	Rev0           decimal.Decimal `gorm:"column:rev0;type:decimal(65,0)" json:"rev0"`
	Rev1           decimal.Decimal `gorm:"column:rev1;type:decimal(65,0)" json:"rev1"`
	Tvl            float64         `gorm:"column:tvl;NOT NULL;DEFAULT:0" json:"tvl,omitempty"`
	ChainId        int             `gorm:"column:chain_id" json:"chain_id"`
	FirstTransfer  uint64          `gorm:"first_transfer;index" json:"first_transfer"`
}
