package rsa

import (
	bigint "bis/lab2/big"
	primecheck "bis/lab2/prime_check"
	cryptorand "crypto/rand"
	"math/big"
)

func GenerateRandomByBits(bits int) *big.Int {
	maxVal := new(big.Int).Lsh(big.NewInt(1), uint(bits)) // max = 2^bits
	maxVal = new(big.Int).Sub(maxVal, big.NewInt(1))      // max = 2^bits - 1
	randomNum, err := cryptorand.Int(cryptorand.Reader, maxVal)
	if err != nil {
		return nil
	}

	return randomNum
}

func GenerateRandomPrimeByBits(bits int) *big.Int {
	num := GenerateRandomByBits(bits)
	target := 10

	for !primecheck.MillerRabin(bigint.New(num.String()), target) && !primecheck.Fermat(bigint.New(num.String()), target) {
		num = GenerateRandomByBits(bits)
	}

	return num
}
