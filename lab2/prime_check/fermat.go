package primecheck

import (
	bigint "bis/lab2/big"
	"strconv"
)

func Fermat(num *bigint.Int, target int) bool {
	corner, _ := strconv.Atoi(num.Number)
	if corner <= 1 || corner == 4 {
		return false
	}

	if corner <= 3 {
		return true
	}

	for range target {
		random := num.Sub(four).GenerateRandomInRange().Add(two).Get() // [2;n-2]

		//(rand^(n-1))*n
		check := random.ModExp(num.Sub(one).Get(), num).Get()
		if check.Number != one.Number {
			return false
		}
	}

	return true
}
