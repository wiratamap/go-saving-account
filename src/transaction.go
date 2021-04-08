package main

type TransactionType string

const(
	CREDIT TransactionType = "CREDIT"
	DEBIT TransactionType = "DEBIT"
)

type Transaction struct {
	amount	int
	transactionType TransactionType 
}