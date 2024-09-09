package main

import (
	"fmt"
)

type Bank struct {
	queue []string
}

func (b *Bank) AddClient(client string) {
	b.queue = append(b.queue, client)
}

func (b *Bank) ServeNextClient() string {
	if len(b.queue) == 0 {
		return "No clients in the queue"
	}

	nextClient := b.queue[0]

	b.queue = b.queue[1:]

	return nextClient
}

func main() {
	bank := Bank{}

	bank.AddClient("Client 1")
	bank.AddClient("Client 2")
	bank.AddClient("Client 3")

	fmt.Println(bank.ServeNextClient())
	fmt.Println(bank.ServeNextClient())
	fmt.Println(bank.ServeNextClient())
	fmt.Println(bank.ServeNextClient())
}
