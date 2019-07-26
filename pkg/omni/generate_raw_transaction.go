package main

import (
	"fmt"

	"github.com/atotto/clipboard"
)

func generateRawTransaction(to string, amount float64) (string, error) {
	satoshi, err := amountToSatoshi(amount)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	hex := decimalToHex(satoshi)

	paddedHex, err := paddingHex(hex)
	if err != nil {
		fmt.Println(err.Error())
		return "", err
	}

	res := conductRawTransaction(to, paddedHex)
	return res, nil
}

func main() {
	var to string
	fmt.Println("Please input recipient address: ")
	fmt.Scan(&to)

	var amount float64
	fmt.Println("Please input USDT amount: ")
	fmt.Scan(&amount)

	tx, err := generateRawTransaction(to, amount)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("************* Copied to clipboard *************")
	fmt.Println(tx)
	fmt.Println("***********************************************")

	err = clipboard.WriteAll(tx)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Succeed!")
}
