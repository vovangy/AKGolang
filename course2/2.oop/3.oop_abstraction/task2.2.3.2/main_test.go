package main

import (
	"testing"
)

func TestSQLiteGenerator_CreateTableSQL(t *testing.T) {
	sqlGen := &SQLiteGenerator{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The method did not panic")
		}
	}()

	sqlGen.CreateTableSQL(nil)
}

func TestSQLiteGenerator_CreateInsertSQL(t *testing.T) {
	sqlGen := &SQLiteGenerator{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The method did not panic")
		}
	}()

	sqlGen.CreateInsertSQL(nil)
}

func TestGoFakeitGenerator_GenerateFakeUser(t *testing.T) {
	fakeGen := &GoFakeitGenerator{}
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The method did not panic")
		}
	}()

	fakeGen.GenerateFakeUser()
}
