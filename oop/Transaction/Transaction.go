package transaction

import (
	"fmt"
)

type transaction struct {
	From	string
	To		string
	Amount	float64
}

// Creates a transaction between two people and how much was sent
func New(from string, to string, amount float64) transaction {
	fmt.Println("Adding a Transaction...")
	t := transaction {from, to, amount}
	return t
}