package cli

import "fmt"

type Operation string

const (
	OperationEncrypt    Operation = "enc"
	OperationDecrypt    Operation = "dec"
	OperationEncryptCBC Operation = "encCBC"
	OperationDecryptCBC Operation = "decCBC"
)

func (o Operation) Valid() bool {
	switch o {
	case OperationEncrypt, OperationDecrypt, OperationEncryptCBC, OperationDecryptCBC:
		return true
	default:
		return false
	}
}

func (o Operation) String() string {
	return fmt.Sprintln(OperationEncrypt, OperationDecrypt, OperationEncryptCBC, OperationDecryptCBC)
}
