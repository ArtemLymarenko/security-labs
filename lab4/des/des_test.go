package mydes

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
	"time"
)

func generateRandomString(maxLength int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := rand.New(rand.NewSource(time.Now().UnixNano())).Intn(maxLength)

	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}

func TestCypher(t *testing.T) {
	//Can be any value but anyway will be taken first 8 bytes
	key := []byte{172, 102, 12, 91, 172, 102, 234, 91, 172, 102, 12, 91, 172, 102, 234, 91}

	c := NewCypher(key)
	var msg string
	const maxLength = 500
	for range 10 {
		for range 1000 {
			msg = generateRandomString(maxLength)
			encrypt := c.Encrypt([]byte(msg))
			decrypt := c.Decrypt(encrypt)

			if msg != string(decrypt) {
				t.Errorf("Expected %s got %s", msg, decrypt)
			}
		}
	}
}

func TestNewCypherCBC(t *testing.T) {
	//Can be any value but anyway will be taken first 8 bytes
	key := []byte{172, 102, 12, 91, 172, 102, 234, 91, 172, 102, 12, 91, 172, 102, 234, 91}

	c := NewCypher(key)
	var msg string
	const maxLength = 100
	for range 10 {
		for range 1000 {
			msg = generateRandomString(maxLength)
			encrypt, _ := c.EncryptCBC([]byte(msg))
			decrypt := c.DecryptCBC(encrypt)

			if msg != string(decrypt) {
				t.Errorf("Expected %s got %s", msg, decrypt)
			}
		}
	}
}

func TestSignCBC(t *testing.T) {
	key := []byte{172, 102, 12, 91, 172, 102, 234, 91, 172, 102, 12, 91, 172, 102, 234, 91}

	c := NewCypher(key)
	var msg string
	const maxLength = 100

	msg = generateRandomString(maxLength)
	_, sign := c.EncryptCBC([]byte(msg))
	fmt.Println("Sign: ", sign)
	if len(sign) != 8 {
		t.Errorf("Expected %v got %v", 8, len(sign))
	}
}
