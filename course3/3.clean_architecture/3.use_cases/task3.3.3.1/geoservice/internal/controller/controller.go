package controller

import (
	"encoding/json"
	"fmt"
	service "geoservice/internal/service"
	models "geoservice/models"
	"net/http"

	"github.com/go-chi/jwtauth"
)

type ControllerOption func(*Controller)

type Controller struct {
	Servicer service.GeoServicer
}

func NewController(token *jwtauth.JWTAuth, options ...ControllerOption) *Controller {
	service := service.NewGeoService(service.WithToken(token))
	controller := &Controller{Servicer: service}
	return controller
}

type Responder interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	LoginUser(w http.ResponseWriter, r *http.Request)
	SearchAnswer(w http.ResponseWriter, r *http.Request)
	GeocodeAnswer(w http.ResponseWriter, r *http.Request)
	NotFoundAnswer(w http.ResponseWriter, r *http.Request)
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
func (c *Controller) RegisterUser(w http.ResponseWriter, r *http.Request) {
	var newUser models.User

	if err := json.NewDecoder(r.Body).Decode(&newUser); err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	err, status := c.Servicer.RegisterUser(newUser)
	if err != nil {
		newErrorResponce(w, err, status)
		return
	}

	message := fmt.Sprintf("User %s sucessfully created", newUser.Username)
	w.WriteHeader(status)
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
func (c *Controller) LoginUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	err, status, tokenString := c.Servicer.LoginUser(user)
	if err != nil {
		newErrorResponce(w, err, status)
		return
	}

	w.WriteHeader(status)
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
func (c *Controller) SearchAnswer(w http.ResponseWriter, r *http.Request) {
	var coordinates models.RequestAddressSearch
	json.NewDecoder(r.Body).Decode(&coordinates)

	err, status, address := c.Servicer.SearchAnswer(coordinates)
	if err != nil {
		newErrorResponce(w, err, status)
		return
	}

	w.WriteHeader(status)
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
func (c *Controller) GeocodeAnswer(w http.ResponseWriter, r *http.Request) {
	var address models.Address
	json.NewDecoder(r.Body).Decode(&address)

	err, status, coords := c.Servicer.GeocodeAnswer(address)
	if err != nil {
		newErrorResponce(w, err, status)
		return
	}

	w.WriteHeader(status)
	w.Write([]byte("Your lattitude = " + coords[0].Lat + "; Your longitude = " + coords[0].Lon))
}

func (c *Controller) NotFoundAnswer(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func newErrorResponce(w http.ResponseWriter, err error, responce int) {
	errResponce := models.ErrorResponce{Message: err.Error()}
	http.Error(w, errResponce.Message, responce)
}
