package cli

import (
	mydes "bis/lab4/des"
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

const resourceDir = "lab4/resources/"

type Cli struct {
	from      string
	to        string
	key       string
	operation Operation
}

func New() *Cli {
	cli := &Cli{}
	cli.parseArgs()
	return cli
}

func (c *Cli) parseArgs() {
	from := flag.String("from", "", "file to take")
	to := flag.String("to", "", "name of file to store")
	key := flag.String("key", "", "key to use in cypher")
	operation := flag.String("operation", "", "type of operation")

	flag.Parse()

	c.from = strings.Trim(*from, " ")
	c.to = strings.Trim(*to, " ")
	c.key = strings.Trim(*key, " ")
	c.operation = Operation(strings.Trim(*operation, " "))
}

func ReadFile(path string) []byte {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}
	defer file.Close()

	result := make([]byte, 0)
	reader := bufio.NewReader(file)
	chunk := make([]byte, 1024)
	for {
		n, err := reader.Read(chunk)
		if err != nil && err.Error() != "EOF" {
			fmt.Println("Error reading file:", err)
			return nil
		}
		if n == 0 {
			break
		}

		result = append(result, chunk[:n]...)
	}

	return result
}

func WriteFile(path string, data []byte) {
	file, err := os.Create(path)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	chunkSize := 1024
	for i := 0; i < len(data); i += chunkSize {
		end := i + chunkSize
		if end > len(data) {
			end = len(data)
		}
		_, err := file.Write(data[i:end])
		if err != nil {
			fmt.Println("Error writing chunk to file:", err)
			return
		}
	}
}

func (c *Cli) Run() {
	if !c.operation.Valid() {
		fmt.Println("Invalid operation try again!")
		fmt.Println("Available operations:", c.operation.String())
		return
	}

	cypher := mydes.NewCypher([]byte(c.key))
	switch c.operation {
	case OperationEncrypt:
		data := ReadFile(resourceDir + c.from)
		encryptData := cypher.Encrypt(data)
		WriteFile(resourceDir+c.to, encryptData)
		break
	case OperationDecrypt:
		data := ReadFile(resourceDir + c.from)
		decryptData := cypher.Decrypt(data)
		WriteFile(resourceDir+c.to, decryptData)
		break
	case OperationEncryptCBC:
		data := ReadFile(resourceDir + c.from)
		encCBC, sign := cypher.EncryptCBC(data)
		fmt.Println("Підпис файлу:", sign)
		WriteFile(resourceDir+c.to, encCBC)
		break
	case OperationDecryptCBC:
		data := ReadFile(resourceDir + c.from)
		decCBC := cypher.DecryptCBC(data)
		WriteFile(resourceDir+c.to, decCBC)
		break
	default:
		return
	}
}
