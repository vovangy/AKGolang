package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func testRoute(t *testing.T, route string, expectedStatus int, expectedBody string) {
	req, err := http.NewRequest("GET", route, nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := NewRouter()

	router.ServeHTTP(rr, req)

	if rr.Code != expectedStatus {
		t.Errorf("handler returned wrong status code: got %v want %v",
			rr.Code, expectedStatus)
	}

	if rr.Body.String() != expectedBody {
		t.Errorf("handler returned wrong body: got %v want %v",
			rr.Body.String(), expectedBody)
	}
}

func TestHandleRoute1(t *testing.T) {
	testRoute(t, "/1", http.StatusOK, "Hello World")
}

func TestHandleRoute2(t *testing.T) {
	testRoute(t, "/2", http.StatusOK, "Hello World 2")
}

func TestHandleRoute3(t *testing.T) {
	testRoute(t, "/3", http.StatusOK, "Hello World 3")
}
