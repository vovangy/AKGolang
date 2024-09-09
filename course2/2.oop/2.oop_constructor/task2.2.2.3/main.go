package main

import (
	"errors"
	"sync"
)

type Account interface {
	Deposit(amount float64)
	Withdraw(amount float64) error
	Balance() float64
}

type SavingsAccount struct {
	mu      sync.Mutex
	balance float64
}

func (a *SavingsAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *SavingsAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < 1000 {
		return errors.New("недостаточно средств: баланс должен быть не менее 1000")
	}
	a.balance -= amount
	return nil
}

func (a *SavingsAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

type CheckingAccount struct {
	mu      sync.Mutex
	balance float64
}

func (a *CheckingAccount) Deposit(amount float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

func (a *CheckingAccount) Withdraw(amount float64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if a.balance < amount {
		return errors.New("недостаточно средств для снятия")
	}
	a.balance -= amount
	return nil
}

func (a *CheckingAccount) Balance() float64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

type Customer struct {
	Name    string
	Account Account
}

type CustomerOption func(*Customer)

func NewCustomer(opts ...CustomerOption) *Customer {
	customer := &Customer{}
	for _, opt := range opts {
		opt(customer)
	}
	return customer
}

func WithName(name string) CustomerOption {
	return func(c *Customer) {
		c.Name = name
	}
}

func WithAccount(account Account) CustomerOption {
	return func(c *Customer) {
		c.Account = account
	}
}
