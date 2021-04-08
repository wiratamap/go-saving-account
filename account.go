package main

import (
	"errors"
)

type Account struct {
	balance int
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) Credit(amount int) {
	account.balance = account.balance + amount
}

func (account *Account) Debit(amount int) error {
	if amount <= 0 {
		return errors.New("InvalidAmountError")
	}
	if amount > account.balance {
		return errors.New("AmountExceedingBalanceError")
	}
	account.balance = account.balance - amount
	return nil
}
