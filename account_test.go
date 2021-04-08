package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T)  {
	t.Run("GetBalance", func(t *testing.T) {
		t.Run("should return 10000 when account balance is 10000", func(t *testing.T) {
			expectedBalance := 10_000
			account := Account{10_000}

			assert.Equal(t, expectedBalance, account.GetBalance())
		})
	})

	t.Run("Credit", func(t *testing.T) {
		t.Run("should add balance to 20000 when credit 10000 balance into account", func(t *testing.T) {
			expectedBalance := 20_000
			account := Account{10_000}

			account.Credit(10_000)

			assert.Equal(t, expectedBalance, account.GetBalance())
		})
	})
}