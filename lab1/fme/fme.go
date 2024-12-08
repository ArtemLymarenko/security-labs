package fme

import (
	"math/big"
)

func ModExp(a *big.Int, b *big.Int, m *big.Int) *big.Int {
	result := big.NewInt(1)
	bigTwo := big.NewInt(2)

	for b.Sign() > 0 {
		if b.Bit(0) == 1 {
			result.Mul(result, a)
			result.Mod(result, m)
		}

		a.Mul(a, a)
		a.Mod(a, m)

		b.Div(b, bigTwo)
	}

	return result
}
