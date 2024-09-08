package main

import (
	"testing"
)

func TestDoubleLinkedList(t *testing.T) {
	list := &DoubleLinkedList{}

	list.Push(Commit{Message: "Initial commit", UUID: "uuid1", Date: "2023-01-01"})
	list.Push(Commit{Message: "Second commit", UUID: "uuid2", Date: "2023-01-02"})
	list.Push(Commit{Message: "Third commit", UUID: "uuid3", Date: "2023-01-03"})
	if list.Len() != 3 {
		t.Fatalf("expected length 3, got %d", list.Len())
	}

	node, err := list.GetByIndex(1)
	if err != nil || node.data.Message != "Second commit" {
		t.Fatalf("expected 'Second commit', got %v", node.data.Message)
	}

	err = list.Delete(1)
	if err != nil || list.Len() != 2 {
		t.Fatalf("expected length 2 after delete, got %d", list.Len())
	}

	node = list.SearchUUID("uuid3")
	if node == nil || node.data.Message != "Third commit" {
		t.Fatalf("expected 'Third commit' for UUID 'uuid3', got %v", node)
	}

	node = list.Search("Initial commit")
	if node == nil || node.data.UUID != "uuid1" {
		t.Fatalf("expected UUID 'uuid1' for 'Initial commit', got %v", node)
	}

	reversedList := list.Reverse()
	if reversedList.Len() != 2 {
		t.Fatalf("expected length 2 for reversed list, got %d", reversedList.Len())
	}

	poppedNode := list.Pop()
	if poppedNode == nil || poppedNode.data.Message != "Third commit" {
		t.Fatalf("expected 'Third commit' for pop, got %v", poppedNode.data.Message)
	}

	shiftedNode := list.Shift()
	if shiftedNode == nil || shiftedNode.data.Message != "Initial commit" {
		t.Fatalf("expected 'Initial commit' for shift, got %v", shiftedNode.data.Message)
	}
}

func TestLoadData(t *testing.T) {
	list := &DoubleLinkedList{}
	err := list.LoadData("testdata.json")
	if err != nil {
		t.Fatalf("failed to load data: %v", err)
	}

	if list.Len() == 0 {
		t.Fatalf("expected non-empty list after loading data")
	}
}

func TestQuickSort(t *testing.T) {
	commits := []Commit{
		{Message: "Commit A", UUID: "uuidA", Date: "2023-01-03"},
		{Message: "Commit B", UUID: "uuidB", Date: "2023-01-01"},
		{Message: "Commit C", UUID: "uuidC", Date: "2023-01-02"},
	}

	QuickSort(commits)
	if commits[0].Date != "2023-01-01" || commits[1].Date != "2023-01-02" || commits[2].Date != "2023-01-03" {
		t.Fatalf("expected sorted dates, got %v", commits)
	}
}
