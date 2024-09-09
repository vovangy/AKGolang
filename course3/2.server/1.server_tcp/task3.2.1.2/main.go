package main

import (
	"net"
	"strings"
)

func handleConnection(c net.Conn) {
	defer c.Close()
	request := make([]byte, 1024)
	_, err := c.Read(request)
	if err != nil {
		return
	}
	sl := strings.Split(string(request), " ")
	sb := strings.Builder{}
	if sl[0] == "GET" && sl[1] == "/" {
		sb.WriteString("HTTP/1.1 200 OK\n")
		sb.WriteString("Content-Type: text/html\n")
		sb.WriteString("\n")
		html := `
			<html>
			<head>
			<title>Webserver</title>
			</head>
			<body>
			<p>hello world</p>
			</body>
			</html>
		`
		sb.WriteString(html)
	} else {
		sb.WriteString("HTTP/1.1 404 Not Found\n")
	}
	_, err = c.Write([]byte(sb.String()))
	if err != nil {
		return
	}
}

func main() {
	listener, _ := net.Listen("tcp", "localhost:8080")
	defer listener.Close()
	for {
		conn, _ := listener.Accept()
		go handleConnection(conn)
	}
}
