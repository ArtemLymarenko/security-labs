package mydes

import (
	"encoding/binary"
)

const BlockSize = 8

type cypher struct {
	keys          [16]uint64
	reversedKeys  [16]uint64
	initialVector uint64
}

func NewCypher(key []byte) *cypher {
	keys := generateSubKeys(key)
	revKeys := reverseKeysOrder(keys)
	initialVector := generateInitialVectorCBC()
	return &cypher{keys, revKeys, initialVector}
}

func (c *cypher) GetInitialVector() uint64 {
	return c.initialVector
}

func reverseKeysOrder(keys [16]uint64) [16]uint64 {
	reversed := [16]uint64{}
	for i := range keys {
		reversed[i] = keys[15-i]
	}

	return reversed
}

func generateSubKeys(key []byte) [16]uint64 {
	subKeys := [16]uint64{}

	block := binary.BigEndian.Uint64(key)
	pc1Key := permutationInt64(block, permutedChoice1[:])
	left := uint32(pc1Key >> 28)
	right := uint32(pc1Key)

	for i := range 16 {
		left = leftRotate28Bits(left, ksRotations[i])
		right = leftRotate28Bits(right, ksRotations[i])
		ctd := concat(left, right, 28)
		subKeys[i] = permutationInt64(ctd, permutedChoice2[:])
	}

	return subKeys
}

func feistel(chunk uint64, keys [16]uint64) uint64 {
	chunkUint64 := permutationInt64(chunk, initialPermutation[:])
	left := uint32(chunkUint64 >> 32)
	right := uint32(chunkUint64)
	tmp := uint32(0)

	for i := 0; i < 16; i++ {
		tmp = right
		expandedRight := expansion32Bits(right)
		expandedRight ^= keys[i]
		right = substitution(expandedRight)
		right = shuffle(right)
		right ^= left
		left = tmp
	}

	cipherChunk := concat(right, left, 32)
	cipherChunk = permutationInt64(cipherChunk, finalPermutation[:])

	return cipherChunk
}

func applyPadding(data []byte) []byte {
	paddingLength := BlockSize - len(data)%BlockSize
	if paddingLength == BlockSize {
		return data
	}
	padding := make([]byte, paddingLength)
	return append(data, padding...)
}

func removePadding(data []byte) []byte {
	for i := len(data) - 1; i >= 0 && data[i] == 0; i-- {
		data = data[:i]
	}
	return data
}

func (c *cypher) Encrypt(data []byte) []byte {
	data = applyPadding(data)

	result := make([]byte, 0, len(data))
	chunk := make([]byte, BlockSize)

	var (
		compressedChunk uint64
		cipherBlock     uint64
		cipherChunk     [8]byte
		endIdx          int
	)
	for i := 0; i < len(data); i += BlockSize {
		endIdx = i + BlockSize
		if endIdx > len(data) {
			endIdx = len(data)
		}
		copy(chunk, data[i:endIdx])

		compressedChunk = binary.BigEndian.Uint64(chunk[:])
		cipherBlock = feistel(compressedChunk, c.keys)

		binary.BigEndian.PutUint64(cipherChunk[:], cipherBlock)
		result = append(result, cipherChunk[:]...)
	}

	return result
}

func (c *cypher) Decrypt(data []byte) []byte {
	result := make([]byte, 0, len(data))
	chunk := make([]byte, BlockSize)

	var (
		compressedChunk uint64
		cipherBlock     uint64
		cipherChunk     [8]byte
		endIdx          int
	)
	for i := 0; i < len(data); i += BlockSize {
		endIdx = i + BlockSize
		if endIdx > len(data) {
			endIdx = len(data)
		}
		copy(chunk, data[i:endIdx])

		compressedChunk = binary.BigEndian.Uint64(chunk[:])
		cipherBlock = feistel(compressedChunk, c.reversedKeys)

		binary.BigEndian.PutUint64(cipherChunk[:], cipherBlock)
		result = append(result, cipherChunk[:]...)
	}

	return removePadding(result)
}
