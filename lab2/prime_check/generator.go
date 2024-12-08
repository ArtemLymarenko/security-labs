package primecheck

import (
	bigint "bis/lab2/big"
)

func GenerateByDecimalLength(length int) *bigint.Int {
	num := bigint.Random(length)
	target := 10

	for !MillerRabin(num, target) && !Fermat(num, target) {
		num = bigint.Random(length)
	}

	return num
}
