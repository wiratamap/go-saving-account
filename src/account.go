package main

import (
	"errors"
)

type Account struct {
	balance      int
	transactions []Transaction
}

func NewAccount(balance int) Account {
	return Account{balance: balance, transactions: []Transaction{}}
}

func (account *Account) GetBalance() int {
	return account.balance
}

func (account *Account) Credit(amount int) error {
	if amount <= 0 {
		return errors.New("InvalidAmountError")
	}
	account.balance += amount
	account.transactions = append(account.transactions, Transaction{amount, CREDIT})
	return nil
}

func (account *Account) Debit(amount int) error {
	if amount <= 0 {
		return errors.New("InvalidAmountError")
	}
	if amount > account.balance {
		return errors.New("AmountExceedingBalanceError")
	}
	account.balance -= amount
	account.transactions = append(account.transactions, Transaction{amount, DEBIT})
	return nil
}

func (account *Account) GetTransactions() []Transaction {
	return account.transactions
}

func (account *Account) RecentTransactionsByType(transactionType TransactionType) []Transaction {
	var filteredTransactions []Transaction
	for _, transaction := range account.transactions {
		if transaction.transactionType == transactionType {
			filteredTransactions = append(filteredTransactions, transaction)
		}
	}
	return filteredTransactions
}
