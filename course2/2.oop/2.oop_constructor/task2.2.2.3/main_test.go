package main

import (
	"testing"
)

func TestCustomer_WithName(t *testing.T) {
	customer := NewCustomer(WithName("Alice"))

	if customer.Name != "Alice" {
		t.Errorf("Expected Name to be 'Alice', got %s", customer.Name)
	}
}

func TestCustomer_WithAccount(t *testing.T) {
	account := &CheckingAccount{}
	customer := NewCustomer(WithAccount(account))

	if customer.Account != account {
		t.Errorf("Expected Account to be %v, got %v", account, customer.Account)
	}
}

func TestSavingsAccount_Withdraw(t *testing.T) {
	account := &SavingsAccount{}
	account.Deposit(2000)

	err := account.Withdraw(500)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if account.Balance() != 1500 {
		t.Errorf("Expected balance to be 1500, got %v", account.Balance())
	}

	err = account.Withdraw(600)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	err = account.Withdraw(600)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}

func TestCheckingAccount_Withdraw(t *testing.T) {
	account := &CheckingAccount{}
	account.Deposit(500)

	err := account.Withdraw(100)
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}

	if account.Balance() != 400 {
		t.Errorf("Expected balance to be 400, got %v", account.Balance())
	}

	err = account.Withdraw(600)
	if err == nil {
		t.Errorf("Expected error, got nil")
	}
}
