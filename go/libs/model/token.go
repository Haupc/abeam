package model

type Token struct {
	Address       string  `json:"address,omitempty" gorm:"primaryKey"`
	Name          string  `json:"name,omitempty" gorm:"index"`
	Decimals      uint8   `json:"decimals,omitempty"`
	Symbol        string  `json:"symbol,omitempty" gorm:"index"`
	ChainID       int     `json:"chain_id,omitempty" gorm:"-"`
	FirstTransfer uint64  `json:"first_transfer,omitempty" gorm:"first_transfer;index"`
	Price         float64 `json:"price,omitempty" gorm:"price"`
	Tvl           float64 `json:"tvl" gorm:"tvl;NOT NULL;DEFAULT:0"`
}
