package main

import (
	"encoding/json"
	"errors"
	"os"
	"time"
)

type Commit struct {
	Message string `json:"message"`
	UUID    string `json:"uuid"`
	Date    string `json:"date"`
}

type Node struct {
	data *Commit
	prev *Node
	next *Node
}

type DoubleLinkedList struct {
	head *Node
	tail *Node
	curr *Node
	len  int
}

type LinkedLister interface {
	LoadData(path string) error
	Init(c []Commit)
	Len() int
	SetCurrent(n *Node) error
	Current() *Node
	Next() *Node
	Prev() *Node
	Insert(n int, c Commit) error
	Push(c Commit) error
	Delete(n int) error
	DeleteCurrent() error
	Index() (int, error)
	GetByIndex(n int) (*Node, error)
	Pop() *Node
	Shift() *Node
	SearchUUID(uuID string) *Node
	Search(message string) *Node
	Reverse() *DoubleLinkedList
}

func (d *DoubleLinkedList) LoadData(path string) error {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	var commits []Commit
	if err := json.Unmarshal(bytes, &commits); err != nil {
		return err
	}

	QuickSort(commits)

	d.Init(commits)
	return nil
}

func (d *DoubleLinkedList) Init(c []Commit) {
	for _, commit := range c {
		d.Push(commit)
	}
}

func (d *DoubleLinkedList) Len() int {
	return d.len
}

func (d *DoubleLinkedList) SetCurrent(n *Node) error {
	if n == nil {
		return errors.New("node is nil")
	}
	d.curr = n
	return nil
}

func (d *DoubleLinkedList) Current() *Node {
	return d.curr
}

func (d *DoubleLinkedList) Next() *Node {
	if d.curr != nil {
		return d.curr.next
	}
	return nil
}

func (d *DoubleLinkedList) Prev() *Node {
	if d.curr != nil {
		return d.curr.prev
	}
	return nil
}

func (d *DoubleLinkedList) Insert(n int, c Commit) error {
	if n < 0 || n > d.len {
		return errors.New("index out of bounds")
	}

	newNode := &Node{data: &c}

	if n == 0 {
		if d.head == nil {
			d.head = newNode
			d.tail = newNode
		} else {
			newNode.next = d.head
			d.head.prev = newNode
			d.head = newNode
		}
	} else if n == d.len {
		newNode.prev = d.tail
		if d.tail != nil {
			d.tail.next = newNode
		}
		d.tail = newNode
	} else {
		current := d.head
		for i := 0; i < n; i++ {
			current = current.next
		}
		newNode.next = current
		newNode.prev = current.prev
		if current.prev != nil {
			current.prev.next = newNode
		}
		current.prev = newNode
	}

	d.len++
	return nil
}

func (d *DoubleLinkedList) Push(c Commit) error {
	return d.Insert(d.len, c)
}

func (d *DoubleLinkedList) Delete(n int) error {
	if n < 0 || n >= d.len {
		return errors.New("index out of bounds")
	}

	var nodeToDelete *Node
	if n == 0 {
		nodeToDelete = d.head
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		}
	} else if n == d.len-1 {
		nodeToDelete = d.tail
		d.tail = d.tail.prev
		if d.tail != nil {
			d.tail.next = nil
		}
	} else {
		nodeToDelete = d.head
		for i := 0; i < n; i++ {
			nodeToDelete = nodeToDelete.next
		}
		nodeToDelete.prev.next = nodeToDelete.next
		nodeToDelete.next.prev = nodeToDelete.prev
	}

	d.len--
	return nil
}

func (d *DoubleLinkedList) DeleteCurrent() error {
	if d.curr == nil {
		return errors.New("current node is nil")
	}

	if d.curr == d.head {
		d.head = d.head.next
		if d.head != nil {
			d.head.prev = nil
		}
	} else if d.curr == d.tail {
		d.tail = d.tail.prev
		if d.tail != nil {
			d.tail.next = nil
		}
	} else {
		d.curr.prev.next = d.curr.next
		d.curr.next.prev = d.curr.prev
	}

	d.curr = nil
	d.len--
	return nil
}

func (d *DoubleLinkedList) Index() (int, error) {
	current := d.head
	index := 0
	for current != nil {
		if current == d.curr {
			return index, nil
		}
		current = current.next
		index++
	}
	return -1, errors.New("current node not found")
}

func (d *DoubleLinkedList) GetByIndex(n int) (*Node, error) {
	if n < 0 || n >= d.len {
		return nil, errors.New("index out of bounds")
	}

	current := d.head
	for i := 0; i < n; i++ {
		current = current.next
	}

	return current, nil
}

func (d *DoubleLinkedList) Pop() *Node {
	if d.tail == nil {
		return nil
	}

	nodeToPop := d.tail
	d.tail = d.tail.prev
	if d.tail != nil {
		d.tail.next = nil
	} else {
		d.head = nil
	}
	d.len--
	return nodeToPop
}

func (d *DoubleLinkedList) Shift() *Node {
	if d.head == nil {
		return nil
	}

	nodeToShift := d.head
	d.head = d.head.next
	if d.head != nil {
		d.head.prev = nil
	} else {
		d.tail = nil
	}
	d.len--
	return nodeToShift
}

func (d *DoubleLinkedList) SearchUUID(uuID string) *Node {
	current := d.head
	for current != nil {
		if current.data.UUID == uuID {
			return current
		}
		current = current.next
	}
	return nil
}

func (d *DoubleLinkedList) Search(message string) *Node {
	current := d.head
	for current != nil {
		if current.data.Message == message {
			return current
		}
		current = current.next
	}
	return nil
}

func (d *DoubleLinkedList) Reverse() *DoubleLinkedList {
	newList := &DoubleLinkedList{}
	current := d.tail
	for current != nil {
		newList.Push(*current.data)
		current = current.prev
	}
	return newList
}

func QuickSort(commits []Commit) {
	if len(commits) < 2 {
		return
	}
	quickSort(commits, 0, len(commits)-1)
}

func quickSort(commits []Commit, low, high int) {
	if low < high {
		pivotIndex := partition(commits, low, high)
		quickSort(commits, low, pivotIndex-1)
		quickSort(commits, pivotIndex+1, high)
	}
}

func partition(commits []Commit, low, high int) int {
	pivot := commits[high]
	pivotDate, _ := time.Parse("2006-01-02", pivot.Date)

	i := low
	for j := low; j < high; j++ {
		dateJ, _ := time.Parse("2006-01-02", commits[j].Date)
		if dateJ.Before(pivotDate) {
			commits[i], commits[j] = commits[j], commits[i]
			i++
		}
	}
	commits[i], commits[high] = commits[high], commits[i]
	return i
}
