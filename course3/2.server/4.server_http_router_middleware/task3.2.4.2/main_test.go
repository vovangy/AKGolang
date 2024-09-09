package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/go-chi/chi"
	"go.uber.org/zap"
)

func TestRoutes(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(LoggerMiddleware(logger))

	r.Get("/route1", handleRoute1)
	r.Get("/route2", handleRoute2)
	r.Get("/route3", handleRoute3)

	tests := []struct {
		url          string
		expectedBody string
		expectedCode int
	}{
		{"/route1", "This is route 1", http.StatusOK},
		{"/route2", "This is route 2", http.StatusOK},
		{"/route3", "This is route 3", http.StatusOK},
	}

	for _, tt := range tests {
		req := httptest.NewRequest(http.MethodGet, tt.url, nil)
		rr := httptest.NewRecorder()
		r.ServeHTTP(rr, req)

		if status := rr.Code; status != tt.expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
		}

		body, err := ioutil.ReadAll(rr.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != tt.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", string(body), tt.expectedBody)
		}
	}
}

func TestLoggerMiddleware(t *testing.T) {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	r := chi.NewRouter()
	r.Use(LoggerMiddleware(logger))

	r.Get("/route1", handleRoute1)

	req := httptest.NewRequest(http.MethodGet, "/route1", nil)
	rr := httptest.NewRecorder()
	r.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

}

func TestMainFunction(t *testing.T) {
	go func() {
		main()
	}()

	time.Sleep(1 * time.Second)

	tests := []struct {
		url          string
		expectedBody string
		expectedCode int
	}{
		{"/route1", "This is route 1", http.StatusOK},
		{"/route2", "This is route 2", http.StatusOK},
		{"/route3", "This is route 3", http.StatusOK},
	}

	for _, tt := range tests {
		resp, err := http.Get("http://localhost:80" + tt.url)
		if err != nil {
			t.Fatalf("Failed to make request: %v", err)
		}
		defer resp.Body.Close()

		if status := resp.StatusCode; status != tt.expectedCode {
			t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedCode)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Fatal(err)
		}
		if string(body) != tt.expectedBody {
			t.Errorf("handler returned unexpected body: got %v want %v", string(body), tt.expectedBody)
		}
	}
}
