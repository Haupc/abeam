package common

import (
	"fmt"
	"math/big"
)

// This custom type will marshal to quoted string-number
type Int big.Int

func NewInt(x int64) *Int {
	return (*Int)(big.NewInt(x))
}

func NewUint(x uint64) *Int {
	bn := new(big.Int)
	bn.SetUint64(x)
	return (*Int)(bn)
}

func (a *Int) MarshalJSON() ([]byte, error) {
	if a == nil {
		return []byte("null"), nil
	}
	return []byte(fmt.Sprintf(`"%s"`, (*big.Int)(a).String())), nil
}

func (a *Int) UnmarshalJSON(data []byte) error {
	if len(data) > 1 && data[0] == '"' && data[len(data)-1] == '"' {
		data = data[1 : len(data)-1]
	}
	if len(data) == 0 {
		(*big.Int)(a).SetInt64(0)
		return nil
	}

	return (*big.Int)(a).UnmarshalJSON(data)
}

func (a *Int) Int() *big.Int {
	return (*big.Int)(a)
}
