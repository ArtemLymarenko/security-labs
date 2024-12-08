package rsa

import (
	"fmt"
	"io/ioutil"
	"math/big"
)

func (pubKey *PublicKey) Encrypt(data []byte) []byte {
	var encryptedData []byte

	blockSize := len(pubKey.N.Bytes()) - 11

	for i := 0; i < len(data); i += blockSize {
		block := data[i:min(i+blockSize, len(data))]
		blockInt := new(big.Int).SetBytes(block)
		encryptedBlock := new(big.Int).Exp(blockInt, pubKey.E, pubKey.N)
		encryptedData = append(encryptedData, encryptedBlock.Bytes()...)
	}

	return encryptedData
}

func (prtKey *PrivateKey) Decrypt(data []byte) []byte {
	var decryptedData []byte

	blockSize := len(prtKey.N.Bytes())

	for i := 0; i < len(data); i += blockSize {
		encryptedBlock := new(big.Int).SetBytes(data[i:min(i+blockSize, len(data))])
		decryptedBlock := new(big.Int).Exp(encryptedBlock, prtKey.D, prtKey.N)
		decryptedData = append(decryptedData, decryptedBlock.Bytes()...)
	}

	return decryptedData
}

func (pubKey *PublicKey) EncryptFile(from, to string) error {
	data, err := ioutil.ReadFile(from)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	encryptedData := pubKey.Encrypt(data)

	err = ioutil.WriteFile(to, encryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}

func (prtKey *PrivateKey) DecryptFile(from, to string) error {
	data, err := ioutil.ReadFile(from)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	decryptedData := prtKey.Decrypt(data)

	err = ioutil.WriteFile(to, decryptedData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %v", err)
	}

	return nil
}
