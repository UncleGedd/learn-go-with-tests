package pointers

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("overdraw error")

type SchruteBuck int

func (sb SchruteBuck) String() string {
	return fmt.Sprintf("%d SB", sb)
}

type Wallet struct {
	balance SchruteBuck
}

func (w *Wallet) Deposit(amount SchruteBuck) {
	// lowercased vars/functions/etc are package private
	w.balance += amount
}

func (w *Wallet) Withdraw(amount SchruteBuck) error {
	if amount > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= amount
	return nil
}

func (w *Wallet) Balance() SchruteBuck {
	// use Wallet to get a copy of the Wallet struct
	// use a *Wallet to reference Wallet struct
	// struct pointers are automatically dereferenced
	return w.balance
}
