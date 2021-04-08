package main

type Account struct {
	balance       int
}

func (account Account) Balance() int {
	return account.balance
}