package mydes

import (
	"encoding/binary"
	"math/rand"
	"time"
)

func generateInitialVectorCBC() uint64 {
	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	return random.Uint64()
}

func (c *cypher) EncryptCBC(data []byte) (result []byte, sign [8]byte) {
	data = applyPadding(data)

	result = make([]byte, 0, len(data))
	chunk := make([]byte, BlockSize)

	var (
		compressedChunk uint64
		encBlock        uint64
		encChunk        [8]byte
		endIdx          int
		prevCypherBlock = c.initialVector
	)

	for i := 0; i < len(data); i += BlockSize {
		endIdx = i + BlockSize
		if endIdx > len(data) {
			endIdx = len(data)
		}
		copy(chunk, data[i:endIdx])

		compressedChunk = binary.BigEndian.Uint64(chunk[:])
		compressedChunk ^= prevCypherBlock

		encBlock = feistel(compressedChunk, c.keys)
		prevCypherBlock = encBlock

		binary.BigEndian.PutUint64(encChunk[:], encBlock)
		result = append(result, encChunk[:]...)
	}

	return result, encChunk
}

func (c *cypher) DecryptCBC(data []byte) []byte {
	result := make([]byte, 0, len(data))
	chunk := make([]byte, BlockSize)

	var (
		compressedChunk uint64
		decBlock        uint64
		decChunk        [8]byte
		endIdx          int
		prevCypherBlock = c.initialVector
	)
	for i := 0; i < len(data); i += BlockSize {
		endIdx = i + BlockSize
		if endIdx > len(data) {
			endIdx = len(data)
		}
		copy(chunk, data[i:endIdx])

		compressedChunk = binary.BigEndian.Uint64(chunk[:])
		decBlock = feistel(compressedChunk, c.reversedKeys)

		decBlock ^= prevCypherBlock
		prevCypherBlock = compressedChunk

		binary.BigEndian.PutUint64(decChunk[:], decBlock)
		result = append(result, decChunk[:]...)
	}

	return removePadding(result)
}
