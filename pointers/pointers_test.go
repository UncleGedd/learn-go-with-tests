package pointers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(SchruteBuck(10))

		got := wallet.Balance()
		want := SchruteBuck(10)

		fmt.Println("demonstrating package private balance", wallet.balance)
		assert.Equal(t, got, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(SchruteBuck(10))

		got := wallet.Balance()
		want := SchruteBuck(10)

		assert.Equal(t, got, want)
	})

	t.Run("overdraw attempt", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		err := wallet.Withdraw(SchruteBuck(11))

		got := wallet.Balance()
		want := SchruteBuck(10)
		assert.NotNil(t, err)
		assert.Equal(t, got, want)
	})
}
