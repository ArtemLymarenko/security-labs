package primecheck

import (
	bigint "bis/lab2/big"
	"strconv"
)

func findPow(num *bigint.Int) (t int) {
	copyNum := num.Get()
	for copyNum.Mod(two).Get().Number == zero.Number {
		copyNum = copyNum.Div(two).Get()
		t++
	}

	return t
}

func intToBig(a int) *bigint.Int {
	return bigint.New(strconv.Itoa(a))
}

func MillerRabin(num *bigint.Int, target int) bool {
	corner, _ := strconv.Atoi(num.Number)
	if corner <= 1 || corner == 4 {
		return false
	}

	if corner <= 3 {
		return true
	}

	if num.Mod(two).Get().Number == zero.Number {
		return false
	}

	numSubOne := num.Sub(one).Get() // (n-1)

	//n-1=2^r*d
	r := findPow(numSubOne)
	d := numSubOne.Div(two.Pow(intToBig(r)).Get()).Get()

	set := NewSet[string]()
	for i := 0; i < target; i++ {
		var randomNum *bigint.Int
		for {
			randomNum = num.Sub(four).GenerateRandomInRange().Add(two).Get() // [2;n-2]
			check, err := strconv.Atoi(num.Number)
			if err == nil && len(set.Keys) == check-3 {
				break
			}

			if !set.Has(randomNum.Number) {
				break
			}
		}
		set.Add(randomNum.Number)

		res := randomNum.ModExp(d, num).Get() // (rand^d)%num
		if res.Number == one.Number || res.Number == numSubOne.Number {
			continue
		}

		isPrime := false
		for j := 1; j <= r; j++ {
			res = res.ModExp(two, num).Get() //(x^2)%num
			if res.Number == numSubOne.Number {
				isPrime = true
				break
			}
		}

		if !isPrime {
			return false
		}
	}

	return true
}
