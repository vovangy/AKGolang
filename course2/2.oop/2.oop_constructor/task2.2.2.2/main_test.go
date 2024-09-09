package main

import "testing"

func TestNewUser_Default(t *testing.T) {
	user := NewUser(1)

	if user.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", user.ID)
	}

	if user.Username != "" {
		t.Errorf("Expected Username to be empty, got %s", user.Username)
	}

	if user.Email != "" {
		t.Errorf("Expected Email to be empty, got %s", user.Email)
	}

	if user.Role != "" {
		t.Errorf("Expected Role to be empty, got %s", user.Role)
	}
}

func TestNewUser_WithUsername(t *testing.T) {
	user := NewUser(1, WithUsername("testuser"))

	if user.Username != "testuser" {
		t.Errorf("Expected Username to be 'testuser', got %s", user.Username)
	}
}

func TestNewUser_WithEmail(t *testing.T) {
	user := NewUser(1, WithEmail("testuser@example.com"))

	if user.Email != "testuser@example.com" {
		t.Errorf("Expected Email to be 'testuser@example.com', got %s", user.Email)
	}
}

func TestNewUser_WithRole(t *testing.T) {
	user := NewUser(1, WithRole("admin"))

	if user.Role != "admin" {
		t.Errorf("Expected Role to be 'admin', got %s", user.Role)
	}
}

func TestNewUser_AllOptions(t *testing.T) {
	user := NewUser(1,
		WithUsername("testuser"),
		WithEmail("testuser@example.com"),
		WithRole("admin"),
	)

	if user.ID != 1 {
		t.Errorf("Expected ID to be 1, got %d", user.ID)
	}

	if user.Username != "testuser" {
		t.Errorf("Expected Username to be 'testuser', got %s", user.Username)
	}

	if user.Email != "testuser@example.com" {
		t.Errorf("Expected Email to be 'testuser@example.com', got %s", user.Email)
	}

	if user.Role != "admin" {
		t.Errorf("Expected Role to be 'admin', got %s", user.Role)
	}
}
