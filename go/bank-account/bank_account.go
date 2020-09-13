// Package account implements simple routines to simulate a bank account
package account

import "sync"

type accountStatus int

const (
	active accountStatus = iota
	closed
)

// Account sturct to hold balance and current status
type Account struct {
	balance int64
	mux     sync.Mutex
	status  accountStatus
}

// Open returns a new Account
func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{balance: initialDeposit, status: active}
}

// Close sets the status to closed and returns closing balance
func (a *Account) Close() (payout int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.status == closed {
		return 0, false
	}

	payout, a.status, a.balance = a.balance, closed, 0
	return payout, true
}

// Balance returns current balance of the account
func (a *Account) Balance() (balance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.status == closed {
		return 0, false
	}
	return a.balance, true
}

// Deposit updates accounts balance and returns new balance
func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mux.Lock()
	defer a.mux.Unlock()

	if a.status == closed || a.balance+amount < 0 {
		return 0, false
	}
	a.balance += amount
	return a.balance, true
}
