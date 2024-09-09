package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"sync"
)

type Client struct {
	conn net.Conn
	name string
	ch   chan<- string
}

type Chat struct {
	entering chan Client
	leaving  chan Client
	msg      chan string
}

func NewChat() *Chat {
	return &Chat{
		leaving:  make(chan Client),
		msg:      make(chan string),
		entering: make(chan Client),
	}
}

func broadcaster(chat *Chat) {
	clients := make(map[Client]bool)

	var mu sync.Mutex
	go func() {
		for enter := range chat.entering {
			mu.Lock()
			clients[enter] = true
			mu.Unlock()
		}
	}()

	go func() {
		for leav := range chat.leaving {
			mu.Lock()
			delete(clients, leav)
			mu.Unlock()
		}
	}()

	for msg := range chat.msg {
		mu.Lock()
		for client := range clients {
			client.ch <- msg
		}
		mu.Unlock()
	}
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for text := range ch {
		_, err := conn.Write([]byte(text))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func handleConn(conn net.Conn, chat *Chat) {
	ch := make(chan string)
	go clientWriter(conn, ch)
	who := conn.RemoteAddr().String()
	cli := Client{conn: conn, name: who, ch: ch}
	ch <- "You are " + who
	chat.msg <- who + " has arrived"
	chat.entering <- cli
	input := bufio.NewScanner(conn)
	for input.Scan() {
		chat.msg <- who + ": " + input.Text()
	}
	chat.leaving <- cli
	chat.msg <- who + " has left"
	conn.Close()
	close(ch)
}

func clientReader(conn net.Conn) {
	for {
		capacityBuff := 1024
		buf := make([]byte, capacityBuff)
		res := ""
		n, err := conn.Read(buf)
		if err != nil {
			log.Printf("read error: %v\n", err)
		}
		for n >= capacityBuff {
			res += string(buf[:n])
			n, err = conn.Read(buf)
			if err != nil {
				log.Printf("read error: %v\n", err)
			}
		}
		res += string(buf[:n])
		fmt.Println(res)
	}

}

func StartClient(conn net.Conn) {
	go clientReader(conn)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := conn.Write([]byte(line + "\r\n"))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func StartServer(listener net.Listener) {
	log.Println("Start server")
	chat := NewChat()
	go broadcaster(chat)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		go handleConn(conn, chat)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":8000")
	if err != nil {
		panic(err)
	}
	defer listener.Close()
	StartServer(listener)
}
