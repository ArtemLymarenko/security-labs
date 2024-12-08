package bigint

import (
	cryptorand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"strconv"
	"strings"
)

func Random(length int) *Int {
	builder := strings.Builder{}

	for length > 0 {
		digit := 1 + rand.Intn(9)
		builder.WriteString(strconv.Itoa(digit))
		length--
	}

	return New(builder.String())
}

func GenerateRandomBigInt(bits int) *Int {
	maxVal := new(big.Int).Lsh(big.NewInt(1), uint(bits)) // max = 2^bits
	maxVal = new(big.Int).Sub(maxVal, big.NewInt(1))      // max = 2^bits - 1
	randomNum, err := cryptorand.Int(cryptorand.Reader, maxVal)
	if err != nil {
		return nil
	}
	fmt.Println(randomNum.String())
	return New(randomNum.String())
}
