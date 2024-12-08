package mydes

func expansion32Bits(block uint32) uint64 {
	var result uint64

	for idx, shift := range expansionFunction {
		toShift := 31 - shift
		resultBit := (block >> toShift) & 1
		result |= uint64(resultBit) << (47 - idx)
	}

	return result
}

func substitution(block uint64) uint32 {
	var result uint32

	for i := 0; i < 8; i++ {
		shiftedKey := (block >> (6 * uint(i))) & 0x3F

		row := ((shiftedKey >> 4) & 2) | shiftedKey&1
		column := (shiftedKey >> 1) & 0x0F

		sBoxValue := sBoxes[i][row][column]
		shifted := uint32(sBoxValue) << (4 * (7 - i))
		result |= shifted
	}

	return result
}

func permutationInt64(block uint64, table []byte) uint64 {
	var result uint64

	for idx, shift := range table {
		toShift := 63 - shift
		resultBit := (block >> toShift) & 1
		result |= resultBit << (len(table) - idx - 1)
	}

	return result
}

func shuffle(block uint32) uint32 {
	var result uint32

	for idx, shift := range permutationFunction {
		resultBit := (block >> (31 - shift)) & 1
		result |= resultBit << (31 - idx)
	}

	return result
}

func leftRotate28Bits(input uint32, shifts uint32) uint32 {
	return (input << shifts) | (input >> (28 - shifts))
}

func concat(left uint32, right uint32, bits int) uint64 {
	return (uint64(left) << bits) | uint64(right)
}
