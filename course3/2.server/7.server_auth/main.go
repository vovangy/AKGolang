package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt/v5"
)

// @title GeoService
// @version 1.0
// @description Simple GeoService.

// @host localhost:8080
// @BasePath /

const (
	adress = "localhost:8080"
)

var Users = make(map[string]string)

var tokenAuth *jwtauth.JWTAuth

func main() {
	r := makeRouter()
	http.ListenAndServe(adress, r)
}

func makeRouter() *chi.Mux {
	r := chi.NewRouter()
	tokenAuth = jwtauth.New("HS256", []byte("mysecretkey"), nil)
	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Post("/api/address/search", searchAnswer)
		r.Post("/api/address/geocode", geocodeAnswer)
	})

	r.Post("/api/register", registerUser)
	r.Post("/api/login", loginUser)
	r.NotFound(usualAnswer)
	return r
}

// registerUser handle POST-requests on api/register
// @Summary Register
// @Tags Login
// @Description Register user
// @Accept  json
// @Produce  json
// @Param  input  body  User true  "username and password"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/register [post]
func registerUser(w http.ResponseWriter, r *http.Request) {
	var newUser User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	if _, ok := Users[newUser.Username]; ok {
		newErrorResponce(w, fmt.Errorf("username already exist"), http.StatusInternalServerError)
		return
	}

	passwordHash := hashPassword([]byte(newUser.Password))
	Users[newUser.Username] = passwordHash
	message := fmt.Sprintf("User %s sucessfully created", newUser.Username)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(message))
}

// loginUser handle POST-requests on api/login
// @Summary Login
// @Tags Login
// @Description Login user
// @Accept  json
// @Produce  json
// @Param  input  body  User true  "username and password"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/login [post]
func loginUser(w http.ResponseWriter, r *http.Request) {
	var user User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	if _, ok := Users[user.Username]; !ok {
		newErrorResponce(w, fmt.Errorf("user dont exist"), http.StatusForbidden)
		return
	}

	passwordHash := hashPassword([]byte(user.Password))
	if passwordHash != Users[user.Username] {
		newErrorResponce(w, fmt.Errorf("invalid password"), http.StatusForbidden)
		return
	}
	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      jwtauth.ExpireIn(time.Hour),
	}
	_, tokenString, _ := tokenAuth.Encode(claims)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(tokenString))
}

// searchAnswer handle POST-requests on api/address/search
// @Summary SearchCity
// @Tags Search
// @Description Search city Name by coords
// @Accept  json
// @Produce  json
// @Param  coordinates  body  RequestAddressSearch true  "Lattitude and Longitude"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/address/search [post]
func searchAnswer(w http.ResponseWriter, r *http.Request) {
	var coordinates RequestAddressSearch
	json.NewDecoder(r.Body).Decode(&coordinates)
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)

	resp, err := http.Get(url)

	if err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	var address ResponseAddress

	err = json.Unmarshal(body, &address)
	if err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("you are in " + address.Address.City))
}

// geocodeAnswer handle POST-requests on api/address/geocode
// @Summary SearchCoords
// @Tags Search
// @Description Search coords by address
// @Accept  json
// @Produce  json
// @Param  coordinates  body  Address true  "House number, road, suburb, city, state, country"
// @Success 200 {object} string
// @Failure 400 {object} errorResponce
// @Failure 500 {object} errorResponce
// @Router /api/address/search [post]
func geocodeAnswer(w http.ResponseWriter, r *http.Request) {
	var address Address
	json.NewDecoder(r.Body).Decode(&address)

	parts := []string{}
	parts = append(parts, strings.Split(address.House_number, " ")...)
	parts = append(parts, strings.Split(address.Road, " ")...)
	parts = append(parts, strings.Split(address.Suburb, " ")...)
	parts = append(parts, strings.Split(address.City, " ")...)
	parts = append(parts, strings.Split(address.State, " ")...)
	parts = append(parts, strings.Split(address.Country, " ")...)

	var sb strings.Builder
	for _, i := range parts {
		if i != "" {
			sb.WriteString("+")
			sb.WriteString(i)
		}
	}

	request := "https://nominatim.openstreetmap.org/search?q=" + strings.Trim(sb.String(), "+") + "&format=json"

	resp, err := http.Get(request)
	if err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	var coords []GetCoords

	err = json.Unmarshal(answer, &coords)
	if err != nil {
		fmt.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Your lattitude = " + coords[0].Lat + "; Your longitude = " + coords[0].Lon))
}

func usualAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

type ResponseAddress struct {
	Address Address `json:"address"`
}

type Address struct {
	House_number string `json:"house_number"`
	Road         string `json:"road"`
	Suburb       string `json:"suburb"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
}

type RequestAddressSearch struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

type GetCoords struct {
	Lat string `json:"lat"`
	Lon string `json:"lon"`
}

type errorResponce struct {
	Message string `json:"message"`
}

func hashPassword(password []byte) string {
	hash := sha256.New()
	hash.Write(password)
	return hex.EncodeToString(hash.Sum(nil))
}
func newErrorResponce(w http.ResponseWriter, err error, responce int) {
	errResponce := errorResponce{Message: err.Error()}
	http.Error(w, errResponce.Message, responce)
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
