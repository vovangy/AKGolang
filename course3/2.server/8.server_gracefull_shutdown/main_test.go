package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/jwtauth"
)

func TestHashPassword(t *testing.T) {
	byteString := []byte("Hello world")
	firstHash := hashPassword(byteString)
	secondHash := hashPassword(byteString)

	if firstHash != secondHash {
		t.Errorf("Should return similar hash on similar input")
	}

	newByteString := []byte("world Hello")
	thirdHash := hashPassword(newByteString)

	if thirdHash == firstHash {
		t.Errorf("Should return different hash on different input")
	}

}

func TestLogin(t *testing.T) {
	user := User{Username: "testuser", Password: "testpassword"}
	Users["testuser"] = hashPassword([]byte("testpassword"))
	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/login", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	loginUser(w, req)
	body, _ := io.ReadAll(w.Body)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if string(body) == "" {
		t.Errorf("expected a token but got an empty string")
	}
}

func TestRegister(t *testing.T) {
	user := User{Username: "test", Password: "testpassword"}

	reqBody, _ := json.Marshal(user)
	req := httptest.NewRequest(http.MethodPost, "/register", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	registerUser(w, req)

	if status := w.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}

	if _, ok := Users["test"]; !ok {
		t.Errorf("register func didnt add user to Users map")
	}
}

func TestGeoCodeAnswer(t *testing.T) {
	address := Address{
		House_number: "1600",
		Road:         "Amphitheatre Parkway",
		Suburb:       "Mountain View",
		City:         "CA",
		State:        "",
		Country:      "",
	}

	reqBody, _ := json.Marshal(address)
	req := httptest.NewRequest(http.MethodPost, "/geocode", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	geocodeAnswer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	result, _ := io.ReadAll(w.Body)

	expected := "Your lattitude = 37.4217636; Your longitude = -122.084614"
	if string(result) != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestSearchAnswer(t *testing.T) {
	coords := RequestAddressSearch{
		Lat: 37.42,
		Lng: -122.08,
	}

	reqBody, _ := json.Marshal(coords)
	req := httptest.NewRequest(http.MethodPost, "/search", strings.NewReader(string(reqBody)))
	w := httptest.NewRecorder()

	searchAnswer(w, req)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	result, _ := io.ReadAll(w.Body)
	expected := "you are in Mountain View"

	if string(result) != expected {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
