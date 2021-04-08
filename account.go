package main

type Account struct {
	balance int
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) Credit(amount int) {
	account.balance = account.balance + amount
}
