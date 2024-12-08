package bigint

import (
	"crypto/rand"
	"math/big"
)

type Operation int

const (
	OpAdd Operation = iota
	OpSub
	OpMul
	OpMod
	OpDiv
	OpPow
)

type BigInt interface {
	Add(b *Int) BigInt
	Sub(b *Int) BigInt
	Mul(b *Int) BigInt
	Div(b *Int) BigInt
	Mod(b *Int) BigInt
	GenerateRandomInRange() BigInt
	Get() *Int
}

type Int struct {
	Number string
}

func New(number string) *Int {
	return &Int{number}
}

func (a *Int) Get() *Int {
	return a
}

func calculate(first *Int, second *Int, op Operation) *Int {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(first.Number, 10)
	b.SetString(second.Number, 10)

	switch op {
	case OpAdd:
		a.Add(a, b)
	case OpSub:
		a.Sub(a, b)
	case OpDiv:
		a.Div(a, b)
	case OpMod:
		a.Mod(a, b)
	case OpMul:
		a.Mod(a, b)
	case OpPow:
		a.Exp(a, b, nil)
	}

	return New(a.String())
}

func (a *Int) Add(b *Int) BigInt {
	return calculate(a, b, OpAdd)
}

func (a *Int) Sub(b *Int) BigInt {
	return calculate(a, b, OpSub)
}

func (a *Int) Mul(b *Int) BigInt {
	return calculate(a, b, OpMul)
}

func (a *Int) Div(b *Int) BigInt {
	return calculate(a, b, OpDiv)
}

func (a *Int) Mod(b *Int) BigInt {
	return calculate(a, b, OpMod)
}

func (a *Int) Pow(b *Int) BigInt {
	return calculate(a, b, OpPow)
}

func (a *Int) GenerateRandomInRange() BigInt {
	N := new(big.Int)
	BigAPlusOne := a.Add(New("1")).Get()
	N.SetString(BigAPlusOne.Number, 10)
	randomNum, err := rand.Int(rand.Reader, N)
	if err != nil {
		return nil
	}

	return New(randomNum.String())
}

func (a *Int) ModExp(b *Int, m *Int) *Int {
	result := big.NewInt(1)
	bigTwo := big.NewInt(2)

	aBig := new(big.Int)
	bBig := new(big.Int)
	mBig := new(big.Int)
	bBig.SetString(b.Number, 10)
	aBig.SetString(a.Number, 10)
	mBig.SetString(m.Number, 10)

	for bBig.Sign() > 0 {
		if bBig.Bit(0) == 1 {
			result.Mul(result, aBig)
			result.Mod(result, mBig)
		}

		aBig.Mul(aBig, aBig)
		aBig.Mod(aBig, mBig)

		bBig.Div(bBig, bigTwo)
	}

	return New(result.String())
}

func (a *Int) Bytes() []byte {
	N := new(big.Int)
	N.SetString(a.Number, 10)
	return N.Bytes()
}
