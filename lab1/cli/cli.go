package bigIntCli

import (
	"bis/lab1/fme"
	"bufio"
	"fmt"
	"math/big"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

type Option string

const (
	opAdd    Option = "1"
	opMul    Option = "2"
	opSqr    Option = "3"
	opMod    Option = "4"
	opModExp Option = "5"
	opExit   Option = "6"
)

type Cli struct {
	prevRes *big.Int
	a       *big.Int
	b       *big.Int
}

func New() *Cli {
	return &Cli{
		prevRes: big.NewInt(0),
		a:       big.NewInt(0),
		b:       big.NewInt(0),
	}
}

func randomBigNumber(length int) string {
	builder := strings.Builder{}

	if rand.Intn(2) == 0 {
		builder.WriteString("-")
	}

	for length > 0 {
		digit := 1 + rand.Intn(9)
		builder.WriteString(strconv.Itoa(digit))
		length--
	}

	return builder.String()
}

func (c *Cli) readBigInt(prompt string) (*big.Int, error) {
	fmt.Print(prompt)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()

	num := new(big.Int)
	num, ok := num.SetString(input, 10)
	if !ok {
		return nil, fmt.Errorf("invalid input: %s", input)
	}
	return num, nil
}

func (c *Cli) input(name string) (*big.Int, error) {
	fmt.Printf("Do you want to enter a number (e) or generate a random number (r)? ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	choice := strings.TrimSpace(scanner.Text())

	var num *big.Int
	var err error

	switch choice {
	case "e":
		num, err = c.readBigInt(fmt.Sprintf("Enter %s number: ", name))
		if err != nil {
			return nil, err
		}
	case "r":
		fmt.Printf("Enter length of number: ")
		scanner.Scan()
		lengthStr := strings.TrimSpace(scanner.Text())
		length, err := strconv.Atoi(lengthStr)
		if err != nil {
			return nil, fmt.Errorf("invalid length: %s", lengthStr)
		}

		randomNumber := randomBigNumber(length)
		num = new(big.Int)
		num.SetString(randomNumber, 10)

		fmt.Printf("%s: %s\n", name, randomNumber)
	default:
		return nil, fmt.Errorf("invalid choice: %s", choice)
	}

	return num, nil
}

func (c *Cli) Start() {
	for {
		fmt.Println()
		fmt.Println(opAdd, "-", "Add")
		fmt.Println(opMul, "-", "Mul")
		fmt.Println(opSqr, "-", "Sqr")
		fmt.Println(opMod, "-", "Mod")
		fmt.Println(opModExp, "-", "ModExp")
		fmt.Println(opExit, "-", "Exit")

		fmt.Print("Choose an option: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		option := Option(scanner.Text())

		switch option {
		case opAdd:
			a, err := c.input("first")
			if err != nil {
				fmt.Println(err)
				continue
			}

			b, err := c.input("second")
			if err != nil {
				fmt.Println(err)
				continue
			}

			c.prevRes = new(big.Int).Add(a, b)
			fmt.Printf("Result: %v\n", c.prevRes)

		case opMul:
			a, err := c.input("first")
			if err != nil {
				fmt.Println(err)
				continue
			}

			b, err := c.input("second")
			if err != nil {
				fmt.Println(err)
				continue
			}

			c.prevRes = new(big.Int).Mul(a, b)
			fmt.Printf("Result: %v\n", c.prevRes)

		case opSqr:
			a, err := c.input("first")
			if err != nil {
				fmt.Println(err)
				continue
			}

			c.prevRes = new(big.Int).Mul(a, a)
			fmt.Printf("Result: %v\n", c.prevRes)

		case opMod:
			a, err := c.input("base")
			if err != nil {
				fmt.Println(err)
				continue
			}

			b, err := c.input("mod")
			if err != nil {
				fmt.Println(err)
				continue
			}
			c.prevRes = new(big.Int).Mod(a, b)
			fmt.Printf("Result: %v\n", c.prevRes)

		case opModExp:
			a, err := c.input("base")
			if err != nil {
				fmt.Println(err)
				continue
			}

			b, err := c.input("mod")
			if err != nil {
				fmt.Println(err)
				continue
			}

			m, err := c.readBigInt("Enter modulus: ")
			if err != nil {
				fmt.Println(err)
				continue
			}

			c.prevRes = fme.ModExp(a, b, m)
			fmt.Printf("Result: %v\n", c.prevRes)

		case opExit:
			fmt.Println("Exiting...")
			return

		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}
