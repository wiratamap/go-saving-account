package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	t.Run("GetBalance", func(t *testing.T) {
		t.Run("should return 10000 when account balance is 10000", func(t *testing.T) {
			expectedBalance := 10_000
			account := NewAccount(10_000)

			assert.Equal(t, expectedBalance, account.GetBalance())
		})
	})

	t.Run("Credit", func(t *testing.T) {
		t.Run("should add balance to 20000 when credit 10000 balance into account", func(t *testing.T) {
			expectedBalance := 20_000
			account := NewAccount(10_000)

			account.Credit(10_000)

			assert.Equal(t, expectedBalance, account.GetBalance())
		})
	})

	t.Run("Debit", func(t *testing.T) {
		t.Run("should deduct balance to 40000 when debit 10000 from account", func(t *testing.T) {
			expectedBalance := 40_000
			account := NewAccount(50_000)

			account.Debit(10_000)

			assert.Equal(t, expectedBalance, account.GetBalance())
		})

		t.Run("should not deduct balance and throw error when debit amount is less than 1", func(t *testing.T) {
			expectedBalance := 50_000
			expectedError := "InvalidAmountError"
			account := NewAccount(50_000)

			error := account.Debit(-50_000)

			assert.Equal(t, expectedBalance, account.GetBalance())
			assert.Equal(t, expectedError, error.Error())
		})

		t.Run("should not deduct balance and throw error when debit amount is less than 1", func(t *testing.T) {
			balance := 10_000
			expectedError := "AmountExceedingBalanceError"
			account := NewAccount(balance)

			error := account.Debit(100_000)

			assert.Equal(t, balance, account.GetBalance())
			assert.Equal(t, expectedError, error.Error())
		})
	})

	t.Run("GetTransactions", func(t *testing.T) {
		t.Run("should return all transaction when doing credit from account", func(t *testing.T) {
			var expectedTransactions []Transaction
			creditTransaction := Transaction{10_000, CREDIT}
			expectedTransactions = append(expectedTransactions, creditTransaction)
			account := NewAccount(10_000)
			account.Credit(10_000)

			transactions := account.GetTransactions()

			assert.Equal(t, expectedTransactions, transactions)
		})

		t.Run("should return all transaction when doing debit from account", func(t *testing.T) {
			var expectedTransactions []Transaction
			creditTransaction := Transaction{10_000, CREDIT}
			debitTransaction := Transaction{5_000, DEBIT}
			expectedTransactions = append(expectedTransactions, creditTransaction)
			expectedTransactions = append(expectedTransactions, debitTransaction)
			account := NewAccount(10_000)
			account.Credit(10_000)
			account.Debit(5_000)

			transactions := account.GetTransactions()

			assert.Equal(t, expectedTransactions, transactions)
		})
	})
}
