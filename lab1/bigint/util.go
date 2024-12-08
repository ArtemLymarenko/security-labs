package bigint

import (
	"fmt"
	"strconv"
	"strings"
)

func swapBuf(buf []byte) {
	for i, j := 0, len(buf)-1; i < j; i, j = i+1, j-1 {
		buf[i], buf[j] = buf[j], buf[i]
	}
}

func (u *UInt) fillFrom(buf []byte) error {
	swapBuf(buf)

	atoi, err := strconv.Atoi(string(buf))
	if err != nil {
		return err
	}

	u.value = append(u.value, Word(atoi))
	return nil
}

func (u *UInt) SetFromString(input string) error {
	input = strings.Trim(input, " ")

	if len(input) == 0 {
		return ErrInvalidString
	}

	if input[0] == '-' {
		return ErrLessThanZero
	}

	buf := make([]byte, 0)
	for i := len(input) - 1; i >= 0; i-- {
		if input[i] == ' ' {
			return ErrInvalidString
		}

		buf = append(buf, input[i])
		if len(buf) == CHUNK {
			err := u.fillFrom(buf)
			if err != nil {
				return err
			}
			buf = buf[:0]
		}
	}

	err := u.fillFrom(buf)
	if err != nil {
		return err
	}

	return nil
}

func (u *UInt) Print() {
	if len(u.value) == 0 {
		fmt.Println("")
		return
	}

	i := len(u.value) - 1
	for i > 0 && u.value[i] == 0 {
		i--
	}

	fmt.Printf("%v", u.value[i])

	for i--; i >= 0; i-- {
		fmt.Printf("%0*d", CHUNK, u.value[i])
	}
	fmt.Println()
}

func (u *UInt) ToString() string {
	if len(u.value) == 0 {
		return ""
	}

	var result strings.Builder
	i := len(u.value) - 1
	for i > 0 && u.value[i] == 0 {
		i--
	}

	result.WriteString(fmt.Sprintf("%v", u.value[i]))

	for i--; i >= 0; i-- {
		result.WriteString(fmt.Sprintf("%0*d", CHUNK, u.value[i]))
	}

	return result.String()
}

func (u *UInt) IsEqual(c *UInt) bool {
	if len(u.value) != len(c.value) {
		return false
	}

	for i := range u.value {
		if u.value[i] != c.value[i] {
			return false
		}
	}
	return true
}
