package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type Logger interface {
	Log(message string) error
}

type ConsoleLogger struct {
	Writer io.Writer
}

func (c *ConsoleLogger) Log(message string) error {
	_, err := fmt.Fprintln(c.Writer, message)
	return err
}

type FileLogger struct {
	File *os.File
}

func (f *FileLogger) Log(message string) error {
	_, err := fmt.Fprintln(f.File, message)
	return err
}

type RemoteLogger struct {
	Address string
}

func (r *RemoteLogger) Log(message string) error {
	_, err := fmt.Printf("Logging remotely to %s: %s\n", r.Address, message)
	return err
}

func LogAll(loggers []Logger, message string) {
	for _, logger := range loggers {
		err := logger.Log(message)
		if err != nil {
			log.Println("Failed to log message:", err)
		}
	}
}

func main() {
	consoleLogger := &ConsoleLogger{Writer: os.Stdout}
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatal("Failed to create log file:", err)
	}
	defer file.Close()
	fileLogger := &FileLogger{File: file}
	remoteLogger := &RemoteLogger{Address: "http://example.com/log"}

	loggers := []Logger{consoleLogger, fileLogger, remoteLogger}

	LogAll(loggers, "\nThis is a test log message.")
}
