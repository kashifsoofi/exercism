// Package account implements simple routines to simulate a bank account
package account

import "sync"

// Account sturct to hold balance and current status
type Account struct {
	sync.RWMutex
	balance int64
	closed  bool
}

// Open returns a new Account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, closed: false}
}

// Close sets the status to closed and returns closing balance
func (a *Account) Close() (payout int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed {
		return 0, false
	}

	payout = a.balance
	a.closed = true
	a.balance = 0
	return payout, true
}

// Balance returns current balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	a.RLock()
	defer a.RUnlock()

	if a.closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit updates accounts balance and returns new balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.Lock()
	defer a.Unlock()

	if a.closed || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
