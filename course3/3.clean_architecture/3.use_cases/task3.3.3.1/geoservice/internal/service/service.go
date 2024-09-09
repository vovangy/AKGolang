package controller

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	models "geoservice/models"

	"github.com/go-chi/jwtauth"
	"github.com/golang-jwt/jwt"
)

type ServiceOption func(*GeoService)

type GeoService struct {
	Users map[string]string
	Token *jwtauth.JWTAuth
}

func WithToken(token *jwtauth.JWTAuth) ServiceOption {
	return func(c *GeoService) {
		c.Token = token
	}
}

func NewGeoService(options ...ServiceOption) *GeoService {
	Users := make(map[string]string)
	service := &GeoService{Users: Users}

	for _, option := range options {
		option(service)
	}

	return service
}

type GeoServicer interface {
	RegisterUser(user models.User) (error, int)
	LoginUser(user models.User) (error, int, string)
	SearchAnswer(coordinates models.RequestAddressSearch) (error, int, models.ResponseAddress)
	GeocodeAnswer(address models.Address) (error, int, []models.GetCoords)
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
func (c *GeoService) RegisterUser(user models.User) (error, int) {
	if _, ok := c.Users[user.Username]; ok {
		return fmt.Errorf("username already exist"), http.StatusInternalServerError
	}

	passwordHash := hashPassword([]byte(user.Password))
	c.Users[user.Username] = passwordHash
	return nil, http.StatusCreated
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
func (c *GeoService) LoginUser(user models.User) (error, int, string) {
	if _, ok := c.Users[user.Username]; !ok {
		return fmt.Errorf("user dont exist"), http.StatusForbidden, ""
	}

	passwordHash := hashPassword([]byte(user.Password))
	if passwordHash != c.Users[user.Username] {
		return fmt.Errorf("invalid password"), http.StatusForbidden, ""
	}

	claims := jwt.MapClaims{
		"username": user.Username,
		"exp":      jwtauth.ExpireIn(time.Hour),
	}
	_, tokenString, _ := c.Token.Encode(claims)

	return nil, http.StatusOK, tokenString
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
func (c *GeoService) SearchAnswer(coordinates models.RequestAddressSearch) (error, int, models.ResponseAddress) {
	var address models.ResponseAddress
	url := fmt.Sprintf("https://nominatim.openstreetmap.org/reverse?format=json&lat=%f&lon=%f", coordinates.Lat, coordinates.Lng)
	resp, err := http.Get(url)

	if err != nil {
		return err, http.StatusInternalServerError, address
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, http.StatusInternalServerError, address
	}

	err = json.Unmarshal(body, &address)
	if err != nil {
		return err, http.StatusInternalServerError, address
	}

	return nil, http.StatusOK, address
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
func (c *GeoService) GeocodeAnswer(address models.Address) (error, int, []models.GetCoords) {
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
	var coords []models.GetCoords

	resp, err := http.Get(request)
	if err != nil {
		return err, http.StatusInternalServerError, coords
	}

	answer, err := io.ReadAll(resp.Body)
	if err != nil {
		return err, http.StatusInternalServerError, coords
	}

	err = json.Unmarshal(answer, &coords)
	if err != nil {
		return err, http.StatusInternalServerError, coords
	}

	return nil, http.StatusOK, coords
}

func hashPassword(password []byte) string {
	hash := sha256.New()
	hash.Write(password)
	return hex.EncodeToString(hash.Sum(nil))
}
