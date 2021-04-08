package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T)  {
	t.Run("Balance", func(t *testing.T) {
		t.Run("should return 10000 when account balance is 10000", func(t *testing.T) {
			expectedBalance := 10_000
			account := Account{10_000}

			assert.Equal(t, expectedBalance, account.Balance())
		})
	})
}