package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBankAccount_Withdraw(t *testing.T) {
	t.Run("Balance: 100, Withdraw: 50", func(t *testing.T) {
		ba := BankAccount{100.0}
		err := ba.Withdraw(50.0)

		assert.NoError(t, err)
		assert.Equal(t, ba.balance, 50.0)
	})

	t.Run("Balance: 100, Withdraw: 200", func(t *testing.T) {
		ba := BankAccount{100.0}
		err := ba.Withdraw(200.0)

		assert.ErrorIs(t, err, &InsufficientFunds{})
	})
}
