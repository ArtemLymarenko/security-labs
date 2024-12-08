package rsa

import (
	cryptorand "crypto/rand"
	"errors"
	"math/big"
)

type PublicKey struct {
	E *big.Int
	N *big.Int
}

type PrivateKey struct {
	D *big.Int
	N *big.Int
}

func GeneratePrimes(bitLength int) (p *big.Int, q *big.Int) {
	halfBits := bitLength / 2
	return GenerateRandomPrimeByBits(halfBits), GenerateRandomPrimeByBits(halfBits)
}

func EulerValue(p, q *big.Int) *big.Int {
	p1 := new(big.Int).Sub(p, one)
	q1 := new(big.Int).Sub(q, one)
	return new(big.Int).Mul(p1, q1)
}

func GCD(a, b *big.Int) *big.Int {
	if a.Cmp(zero) == 0 {
		return b
	}

	return GCD(new(big.Int).Mod(b, a), a)
}

func ExtendedEuclideanAlg(a, b *big.Int) (*big.Int, *big.Int) {
	if a.Cmp(zero) == 0 {
		return zero, one
	}

	x1, y1 := ExtendedEuclideanAlg(new(big.Int).Mod(b, a), a)

	x := new(big.Int).Sub(y1, new(big.Int).Mul(new(big.Int).Div(b, a), x1))
	y := x1

	return x, y
}

func ModInverse(num, mod *big.Int) *big.Int {
	x, _ := ExtendedEuclideanAlg(num, mod)
	return new(big.Int).Mod(x, mod)
}

func GenerateKeys(bitLength int) (public *PublicKey, private *PrivateKey, err error) {
	p, q := GeneratePrimes(bitLength)
	n := new(big.Int).Mul(p, q)
	phi := EulerValue(p, q)

	var e *big.Int
	for {
		phiSub4 := new(big.Int).Sub(phi, four)
		randomNum, err := cryptorand.Int(cryptorand.Reader, phiSub4)
		if err != nil {
			return nil, nil, errors.New("failed to generate keys")
		}

		e = new(big.Int).Add(randomNum, three) // [3; n-1]
		if GCD(e, phi).Cmp(one) == 0 {
			break
		}
	}

	public = &PublicKey{
		E: e,
		N: n,
	}

	d := ModInverse(e, phi)

	private = &PrivateKey{
		D: d,
		N: n,
	}

	return public, private, nil
}
