package service

import (
	"golibrary/entities"
	repository "golibrary/internal/repository"
	"math/rand"
	"time"

	gofakeit "github.com/brianvoe/gofakeit/v6"
)

type LibraryFacade struct {
	Service repository.Librarer
}

type Servicer interface {
	StartProgram() error
	TakeBook(userID, bookID int) error
	ReturnBook(book entities.Book) error
	AllUsersInfo() ([]entities.User, error)
	AllAuthorsInfo() ([]entities.Author, error)
	AddBook(book entities.Book) error
	GetAllBooks() ([]entities.Book, error)
}

func NewLibraryFacade() (*LibraryFacade, error) {
	service, err := repository.NewLibraryService()
	return &LibraryFacade{Service: service}, err
}

func (lf *LibraryFacade) StartProgram() error {
	authorsNumber, err := lf.Service.HowManyAuthorsExist()
	if err != nil {
		return err
	}

	if authorsNumber < 10 {
		for i := 0; i < (10 - authorsNumber); i++ {
			lf.Service.AddAuthor(gofakeit.BookAuthor())
		}
	}
	if authorsNumber < 10 {
		authorsNumber = 10
	}

	booksNumber, err := lf.Service.HowManyBooksExist()
	if err != nil {
		return err
	}

	if booksNumber < 100 {
		for i := 0; i < (100 - booksNumber); i++ {
			rand.Seed(time.Now().UnixNano())
			authorID := rand.Intn(authorsNumber) + 1

			book := entities.Book{
				Name:     gofakeit.BookTitle(),
				AuthorID: authorID,
			}

			err = lf.Service.AddBook(book)
			if err != nil {
				return err
			}
		}
	}

	usersNumber, err := lf.Service.HowManyUsersExist()
	if err != nil {
		return err
	}

	if usersNumber < 50 {
		for i := 0; i < (50 - usersNumber); i++ {
			user := entities.User{
				UserName: gofakeit.Name(),
			}
			lf.Service.AddUser(user)
		}
	}

	return nil
}

func (lf *LibraryFacade) TakeBook(userID, bookID int) error {
	return lf.Service.TakeBook(userID, bookID)
}

func (lf *LibraryFacade) ReturnBook(book entities.Book) error {
	return lf.Service.ReturnBook(book)
}

func (lf *LibraryFacade) AllUsersInfo() ([]entities.User, error) {
	users, err := lf.Service.GetAllUsers()
	if err != nil {
		return nil, err
	}

	for _, user := range users {
		userBooks, err := lf.Service.GetAllBooksTakenByUser(user.Id)
		if err != nil {
			return nil, err
		}

		user.BooksTaken = userBooks
	}

	return users, nil
}

func (lf *LibraryFacade) AllAuthorsInfo() ([]entities.Author, error) {
	authors, err := lf.Service.GetAllAuthors()
	if err != nil {
		return nil, err
	}

	for _, author := range authors {
		authorBooks, err := lf.Service.GetAllAuthorBooks(author.Id)
		if err != nil {
			return nil, err
		}

		author.Books = authorBooks
	}

	return authors, nil
}

func (lf *LibraryFacade) AddBook(book entities.Book) error {
	return lf.Service.AddBook(book)
}

func (lf *LibraryFacade) GetAllBooks() ([]entities.Book, error) {
	booksQuantity, err := lf.Service.HowManyBooksExist()
	if err != nil {
		return nil, err
	}

	books := []entities.Book{}
	for i := 1; i <= booksQuantity; i++ {
		book, err := lf.Service.GetBookInfoByID(i)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}
