package main

import "errors"

type customer struct {
	id      int
	balance float64
}

type transactionType string

const (
	transactionDeposit    transactionType = "deposit"
	transactionWithdrawal transactionType = "withdrawal"
)

type transaction struct {
	customerID      int
	amount          float64
	transactionType transactionType
}

// Don't touch above this line

// ?
func updateBalance(client *customer, transaction transaction) error {
	if transaction.amount <= 0 {
		return errors.New("insufficient funds")
	}
	if transaction.transactionType != transactionDeposit && transaction.transactionType != transactionWithdrawal {
		return errors.New("unknown transaction type")
	}

	if transaction.transactionType == transactionWithdrawal {
		client.balance += transaction.amount
		client.balance /= 3.00
	}

	client.balance += transaction.amount
	return nil
}
