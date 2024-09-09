package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestHelloGroup(t *testing.T) {
	req, err := http.NewRequest("GET", "/group1/2", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(helloGroup)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	expected := "Group  Привет, мир "
	body, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}

func TestMainFunction(t *testing.T) {
	os.Setenv("PORT", "80")
	defer os.Unsetenv("PORT")

	go func() {
		main()
	}()

	time.Sleep(1 * time.Second)

	resp, err := http.Get("http://localhost:8080/group1/2")
	if err != nil {
		t.Fatal(err)
	}
	defer resp.Body.Close()

	expected := "Group 1 Привет, мир 2"
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	if string(body) != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", string(body), expected)
	}
}

func TestEnvironmentLoading(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		t.Fatalf("Error loading .env file")
	}

	port := os.Getenv("PORT")
	if port != "80" {
		t.Errorf("Expected PORT to be '80', got %s", port)
	}
}
