package main

import (
	"testing"
)

func TestInsertAndSearch(t *testing.T) {
	bt := NewBTree(2)

	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 2, Name: "Bob", Age: 25},
		{ID: 3, Name: "Charlie", Age: 35},
	}

	for _, user := range users {
		bt.Insert(user)
	}

	tests := []struct {
		id       int
		expected User
		found    bool
	}{
		{1, User{ID: 1, Name: "Alice", Age: 30}, true},
		{2, User{ID: 2, Name: "Bob", Age: 25}, true},
		{3, User{ID: 3, Name: "Charlie", Age: 35}, true},
		{4, User{}, false},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			user := bt.Search(tt.id)
			if (user != nil) != tt.found {
				t.Errorf("Search(%d) found = %v, want found = %v", tt.id, user != nil, tt.found)
			}
			if user != nil && *user != tt.expected {
				t.Errorf("Search(%d) = %v, want %v", tt.id, *user, tt.expected)
			}
		})
	}
}

func TestEmptyTree(t *testing.T) {
	bt := NewBTree(2)

	if user := bt.Search(1); user != nil {
		t.Errorf("Search(1) = %v, want nil", user)
	}
}

func TestInsertDuplicate(t *testing.T) {
	bt := NewBTree(2)

	users := []User{
		{ID: 1, Name: "Alice", Age: 30},
		{ID: 1, Name: "Alice", Age: 30},
	}

	for _, user := range users {
		bt.Insert(user)
	}

	user := bt.Search(1)
	if user == nil {
		t.Errorf("Search(1) = nil, want User")
	}
	if user.Name != "Alice" || user.Age != 30 {
		t.Errorf("Search(1) = %v, want {ID: 1, Name: Alice, Age: 30}", user)
	}
}
