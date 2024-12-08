package bigint

import "math/big"

const (
	CHUNK = 9
	POW10 = 1_000_000_000
)

type BigUInt interface {
	Add(b *UInt) BigUInt
	Mul(b *UInt) BigUInt
	Sqr() BigUInt
	Mod(b *UInt) BigUInt
	Get() *UInt
}

type Dword uint64
type Word uint32

type UInt struct {
	value []Word
}

func (u *UInt) Get() *UInt {
	return u
}

func New() *UInt {
	return &UInt{}
}

func (u *UInt) Add(b *UInt) BigUInt {
	if len(u.value) == 0 || len(b.value) == 0 {
		return u
	}

	result := New()
	var (
		carry       Word
		chunkResult Word
	)

	for i, j := 0, 0; i < len(u.value) || j < len(b.value); i, j = i+1, j+1 {
		if i < len(u.value) {
			chunkResult += u.value[i]
		}

		if j < len(b.value) {
			chunkResult += b.value[i]
		}

		chunkResult += carry

		if chunkResult >= POW10 {
			carry = chunkResult / POW10
			chunkResult = chunkResult % POW10
		} else {
			carry = 0
		}

		result.value = append(result.value, chunkResult)
		chunkResult = 0
	}

	return result
}

func (u *UInt) Mul(b *UInt) BigUInt {
	if len(u.value) == 0 || len(b.value) == 0 {
		return u
	}

	var (
		res         = New()
		chunkResult Dword
	)

	res.value = make([]Word, len(u.value)+len(b.value))

	for i := 0; i < len(u.value); i++ {
		carry := Dword(0)
		for j := 0; j < len(b.value); j++ {
			chunkResult = Dword(u.value[i])*Dword(b.value[j]) + Dword(res.value[i+j]) + carry
			res.value[i+j] = Word(chunkResult % POW10)
			carry = chunkResult / Dword(POW10)
		}
		res.value[i+len(b.value)] = Word(carry)
	}

	for len(res.value) > 1 && res.value[len(res.value)-1] == 0 {
		res.value = res.value[:len(res.value)-1]
	}

	return res
}

func (u *UInt) Sqr() BigUInt {
	return u.Mul(u)
}

func (u *UInt) Mod(c *UInt) BigUInt {
	a := new(big.Int)
	b := new(big.Int)
	a.SetString(u.ToString(), 10)
	b.SetString(c.ToString(), 10)

	a.Mod(a, b)

	res := New()
	res.SetFromString(a.String())

	return res
}
