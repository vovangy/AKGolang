package controller

import (
	"encoding/json"
	entities "golibrary/entities"
	service "golibrary/internal/service"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
)

type UserService struct {
	Service service.Servicer
}

type UserServicer interface {
	GetAllUsers(w http.ResponseWriter, r *http.Request)
	GetAllAuthors(w http.ResponseWriter, r *http.Request)
	AddBook(w http.ResponseWriter, r *http.Request)
	GetAllBooks(w http.ResponseWriter, r *http.Request)
	TakeBook(w http.ResponseWriter, r *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
	NotFound(w http.ResponseWriter, r *http.Request)
}

func StartUserService() (*UserService, error) {
	service, err := service.NewLibraryFacade()
	if err != nil {
		return nil, err
	}

	err = service.StartProgram()
	if err != nil {
		return nil, err
	}

	return &UserService{Service: service}, err
}

func (us *UserService) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := us.Service.AllUsersInfo()
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	var sb strings.Builder
	for _, user := range users {
		json.NewEncoder(&sb).Encode(user)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sb.String()))
}

func (us *UserService) GetAllAuthors(w http.ResponseWriter, r *http.Request) {
	authors, err := us.Service.AllAuthorsInfo()
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	var sb strings.Builder
	for _, author := range authors {
		json.NewEncoder(&sb).Encode(author)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sb.String()))
}

func (us *UserService) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books, err := us.Service.GetAllBooks()
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	var sb strings.Builder
	for _, book := range books {
		json.NewEncoder(&sb).Encode(book)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(sb.String()))
}

func (us *UserService) AddBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := us.Service.AddBook(book)
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	answer := "Book " + book.Name + " sucessfully added."
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func (us *UserService) TakeBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	var book entities.Book
	json.NewDecoder(r.Body).Decode(&book)

	userId, err := strconv.Atoi(id)
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	err = us.Service.TakeBook(userId, book.Id)
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	answer := "Book " + book.Name + " sucessfully added."
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func (us *UserService) ReturnBook(w http.ResponseWriter, r *http.Request) {
	var book entities.Book
	json.NewDecoder(r.Body).Decode(&book)

	err := us.Service.ReturnBook(book)
	if err != nil {
		us.newErrorResponce(w, err, http.StatusInternalServerError)
		return
	}

	answer := "Book " + book.Name + " sucessfully added."
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(answer))
}

func (us *UserService) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("Not Found"))
}

func (us *UserService) newErrorResponce(w http.ResponseWriter, err error, responce int) {
	errResponce := entities.ErrorResponce{Message: err.Error()}
	http.Error(w, errResponce.Message, responce)
}
