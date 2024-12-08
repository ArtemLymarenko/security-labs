package primecheck

import (
	bigint "bis/lab2/big"
	"fmt"
	"strconv"
	"testing"
)

func TestEqual(t *testing.T) {
	for i := 2; i < 1000; i++ {
		n := strconv.Itoa(i)
		bigInt := bigint.New(n)

		millerRabin := MillerRabin(bigInt, i)
		fermat := Fermat(bigInt, i)
		t.Logf("Miller-Rabbin %v %v", n, fermat)
		t.Logf("Fermat %v %v", n, fermat)

		if millerRabin != fermat {
			t.Error(millerRabin != fermat)
		}
	}
}

func TestPrimeCheck(t *testing.T) {
	test1 := bigint.New("531137992816767098689588206552468627329593117727031923199444138200403559860852242739162502265229285668889329486246501015346579337652707239409519978766587351943831270835393219031728127")
	test2 := bigint.New("1531137992816767098689588206552468627329593117727031923199444138200403559860852242739162502265229285668889329486246501015346579337652707239409519978766587351943831270835393219031728127")
	target := 512

	result1MillerRabin := MillerRabin(test1, target)
	if !result1MillerRabin {
		t.Errorf("Test1 should be prime (MillerRabin)")
	}

	result1Fermat := Fermat(test1, target)
	if !result1Fermat {
		t.Errorf("Test1 should be prime (Fermat)")
	}

	result2MillerRabin := MillerRabin(test2, target)
	if result2MillerRabin {
		t.Errorf("Test2 should not be prime (MillerRabin)")
	}

	result2Fermat := Fermat(test2, target)
	if result2Fermat {
		t.Errorf("Test2 should not be prime (Fermat)")
	}

	fmt.Printf("Test1 is prime(MillerRabin): %v\n", result1MillerRabin)
	fmt.Printf("Test1 is prime(Fermat): %v\n", result1Fermat)
	fmt.Printf("Test2 is prime(MillerRabin): %v\n", result2MillerRabin)
	fmt.Printf("Test2 is prime(Fermat): %v\n", result2Fermat)
}
