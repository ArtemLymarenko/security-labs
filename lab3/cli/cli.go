package cli

import (
	"bis/lab3/rsa"
	"fmt"
	"log"
)

type CLI struct{}

func New() *CLI {
	return &CLI{}
}

func (cli *CLI) displayResults(publicKey, privateKey, n string) {
	fmt.Println("Generated Public Key: ", publicKey)
	fmt.Println("Generated Private Key: ", privateKey)
	fmt.Println("N: ", n)
}

func (cli *CLI) ReadInOutFiles() (in, out string) {
	fmt.Println("Enter the input file path for encryption: ")
	fmt.Scanln(&in)
	fmt.Println("Enter the output file path for encrypted data: ")
	fmt.Scanln(&out)

	return in, out
}

func (cli *CLI) Run() {
	var bitSize int
	fmt.Println("Enter amount of bits: ")
	_, err := fmt.Scanln(&bitSize)
	if err != nil {
		log.Fatal("Error reading input:", err)
	}

	public, private, err := rsa.GenerateKeys(bitSize)
	if err != nil {
		log.Fatal(err)
	}

	cli.displayResults(public.E.String(), private.D.String(), public.N.String())

	const (
		OptionEncrypt = iota
		OptionDecrypt
		OptionExit
	)
	for {
		var operation int
		fmt.Println("\nChoose an operation (encrypt - 0, decrypt - 1, exit - 2): ")
		_, err := fmt.Scanln(&operation)
		if err != nil {
			log.Fatal("Error reading input:", err)
		}

		const filePrefix = "lab3/resources/"
		switch operation {
		case OptionEncrypt:
			var filepath, encryptedFilepath = cli.ReadInOutFiles()
			err := public.EncryptFile(filePrefix+filepath, filePrefix+encryptedFilepath)
			if err != nil {
				log.Println("Error encrypting file:", err)
				break
			}
			fmt.Println("File encrypted successfully.")

		case OptionDecrypt:
			var filepath, decryptedFilepath = cli.ReadInOutFiles()
			err := private.DecryptFile(filePrefix+filepath, filePrefix+decryptedFilepath)
			if err != nil {
				log.Println("Error decrypting file:", err)
				break
			}
			fmt.Println("File decrypted successfully.")

		case OptionExit:
			fmt.Println("Exiting the program.")
			return

		default:
			fmt.Println("Choose valid operation.")
		}
	}
}
