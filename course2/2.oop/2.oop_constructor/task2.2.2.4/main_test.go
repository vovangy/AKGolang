package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

func TestFileLogger(t *testing.T) {
	file, err := ioutil.TempFile("", "logtest")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(file.Name())

	fileLogger := FileLogger{file: file}
	logSystem := NewLogSystem(WithLogger(fileLogger))

	message := "Test message"
	logSystem.Log(message)

	file.Seek(0, 0)
	content, err := ioutil.ReadAll(file)
	if err != nil {
		t.Fatalf("Failed to read log file: %v", err)
	}

	expected := message + "\n"
	if string(content) != expected {
		t.Fatalf("Expected %q, got %q", expected, string(content))
	}
}

func TestConsoleLogger(t *testing.T) {
	var buf bytes.Buffer
	consoleLogger := ConsoleLogger{out: &buf}
	logSystem := NewLogSystem(WithLogger(consoleLogger))

	message := "Test console message"
	logSystem.Log(message)

	expected := message + "\n"
	if buf.String() != expected {
		t.Fatalf("Expected %q, got %q", expected, buf.String())
	}
}
