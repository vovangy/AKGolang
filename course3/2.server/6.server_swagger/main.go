package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/swaggo/http-swagger" // swagger docs
	// Add for Swagger middleware generation
)

// @title Address Search API
// @version 1.0
// @description This is an API for searching addresses and geocoding.
// @host localhost:8080
// @BasePath /api

// ResponseAddress contains the details of an address
type ResponseAddress struct {
	Address Address `json:"address"`
}

// Address holds the structure of address fields
type Address struct {
	Road        string `json:"road"`
	Town        string `json:"town"`
	County      string `json:"county"`
	State       string `json:"state"`
	Postcode    string `json:"postcode"`
	Country     string `json:"country"`
	CountryCode string `json:"country_code"`
}

// RequestAddressGeocode represents coordinates for geocoding
type RequestAddressGeocode struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// ResponseAddressGeocode represents geocoding response
type ResponseAddressGeocode struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

// @Summary Search address by coordinates
// @Description Search address based on latitude and longitude.
// @Accept json
// @Produce json
// @Param address body RequestAddressGeocode true "Lat and Lng"
// @Success 200 {string} string "You are in Country: Country, Town: Town, Road: Road"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/search [post]
func handlerSearch(w http.ResponseWriter, r *http.Request) {
	var coord RequestAddressGeocode
	err := json.NewDecoder(r.Body).Decode(&coord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coord.Lat, coord.Lng)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", url, http.NoBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	address := &ResponseAddress{}
	err = json.Unmarshal(body, address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("You are in Country:%s, Town:%s, Road:%s", address.Address.Country, address.Address.Town, address.Address.Road)
	_, _ = w.Write([]byte(response))
}

// @Summary Geocode address to coordinates
// @Description Converts an address into geographic coordinates.
// @Accept json
// @Produce json
// @Param address body ResponseAddress true "Address Data"
// @Success 200 {string} string "Coordinate: Lon: Longitude, Lat: Latitude"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/address/geocode [post]
func handlerGeocode(w http.ResponseWriter, r *http.Request) {
	var address ResponseAddress
	err := json.NewDecoder(r.Body).Decode(&address)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	q := GetQuery(address)
	request := fmt.Sprintf("https://nominatim.openstreetmap.org/search?q=%s&format=json", q)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, "GET", request, http.NoBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	client := http.DefaultClient
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	coord := []ResponseAddressGeocode{}
	err = json.Unmarshal(body, &coord)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	response := fmt.Sprintf("Coordinate: Lon:%v, Lat:%v", coord[0].Lon, coord[0].Lat)
	_, _ = w.Write([]byte(response))
}

// GetQuery concatenates query parts for geocoding
func GetQuery(address ResponseAddress) string {
	parts := []string{}
	parts = append(parts, strings.Split(address.Address.Road, " ")...)
	parts = append(parts, strings.Split(address.Address.Town, " ")...)
	parts = append(parts, strings.Split(address.Address.State, " ")...)
	parts = append(parts, strings.Split(address.Address.Country, " ")...)

	var sb strings.Builder
	for _, i := range parts {
		if i != "" {
			sb.WriteString("+")
			sb.WriteString(i)
		}
	}
	return strings.Trim(sb.String(), "+")
}

func handlerNot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}

func NewRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Post("/api/address/search", handlerSearch)
	router.Post("/api/address/geocode", handlerGeocode)
	router.NotFound(handlerNot)
	return router
}

func main() {
	_ = http.ListenAndServe(":8080", NewRouter())
}
