package main

import (
	"fmt"
	"./oop/transaction"
)

func main() {
	fmt.Println("Starting Program")

	test := transaction.New("Bob", "Phil", 57000)
	fmt.Println(test.From)
}