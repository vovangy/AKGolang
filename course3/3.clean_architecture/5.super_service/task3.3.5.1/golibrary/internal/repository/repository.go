package repository

import (
	"database/sql"
	"fmt"
	entities "golibrary/entities"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type LibraryService struct {
	Base *sql.DB
}

type Librarer interface {
	AddAuthor(authorName string) error
	AddBook(book entities.Book) error
	AddUser(user entities.User) error
	TakeBook(userID, bookID int) error
	ReturnBook(book entities.Book) error
	GetBookInfoByID(id int) (entities.Book, error)
	GetAuthorByID(id int) (entities.Author, error)
	GetUserByID(id int) (entities.User, error)
	GetAllUsers() ([]entities.User, error)
	GetAllBooksTakenByUser(userID int) ([]entities.Book, error)
	GetAllAuthors() ([]entities.Author, error)
	GetAllAuthorBooks(authorID int) ([]entities.Book, error)
	HowManyAuthorsExist() (int, error)
	HowManyBooksExist() (int, error)
	HowManyUsersExist() (int, error)
}

func NewLibraryService() (*LibraryService, error) {
	err := godotenv.Load()

	if err != nil {
		return nil, err
	}

	service := &LibraryService{}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	connStr := "user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " host=" + dbHost + " port=" + dbPort
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return service, err
	}

	service.Base = db
	err = service.CreateNewUserTable()
	if err != nil {
		return service, err
	}

	err = service.CreateNewBookTable()
	if err != nil {
		return service, err
	}

	err = service.CreateAuthorsTable()
	if err != nil {
		return service, err
	}

	return service, err
}

func (ls *LibraryService) CreateNewUserTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(100) NOT NULL UNIQUE
	);`

	_, err := ls.Base.Exec(newTableString)
	return err
}

func (ls *LibraryService) CreateNewBookTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS books (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL,
		authorid INTEGER NOT NULL,
		usertakenid INTEGER DEFAULT 0,
		istaken BOOL DEFAULT false
	);`

	_, err := ls.Base.Exec(newTableString)
	return err
}

func (ls *LibraryService) CreateAuthorsTable() error {
	newTableString := `CREATE TABLE IF NOT EXISTS authors (
		id SERIAL PRIMARY KEY,
		name VARCHAR(100) NOT NULL UNIQUE
	);`

	_, err := ls.Base.Exec(newTableString)
	return err
}

func (ls *LibraryService) AddAuthor(authorName string) error {
	tx, err := ls.Base.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO authors (name)
		VALUES ($1)
		ON CONFLICT (name) DO NOTHING;
	`

	result, err := tx.Exec(query, authorName)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("user with username %s already exists", authorName)
	}

	err = tx.Commit()
	return err
}

func (ls *LibraryService) AddBook(book entities.Book) error {
	query := `SELECT EXISTS (SELECT name FROM authors WHERE id = $1)`
	var authorExist bool
	row := ls.Base.QueryRow(query, book.AuthorID)
	err := row.Scan(&authorExist)
	if err != nil {
		return err
	}

	if !authorExist {
		return fmt.Errorf("author with id %d doesnt exist", book.AuthorID)
	}

	tx, err := ls.Base.Begin()
	if err != nil {
		return err
	}

	query = `
		INSERT INTO books (name, authorid)
		VALUES ($1, $2)
	`

	_, err = tx.Exec(query, book.Name, book.AuthorID)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}

func (ls *LibraryService) AddUser(user entities.User) error {
	tx, err := ls.Base.Begin()
	if err != nil {
		return err
	}

	query := `
		INSERT INTO users (username)
		VALUES ($1)
	`

	_, err = tx.Exec(query, user.UserName)
	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()
	return nil
}

func (ls *LibraryService) TakeBook(userID, bookID int) error {
	query := `SELECT istaken FROM books WHERE id = $1`
	var bookIsTaken bool
	row := ls.Base.QueryRow(query, bookID)
	err := row.Scan(&bookIsTaken)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("book with id %d not found", bookID)
		}

		return err
	}

	if bookIsTaken {
		return fmt.Errorf("book with id %d is already taken", bookID)
	}

	query = `UPDATE books SET usertakenid = $1, istaken = $2 WHERE id = $3`
	_, err = ls.Base.Exec(query, userID, true, bookID)
	if err != nil {
		return err
	}

	return nil
}

func (ls *LibraryService) ReturnBook(book entities.Book) error {
	query := `SELECT istaken FROM books WHERE id = $1`
	var bookIsTaken bool
	row := ls.Base.QueryRow(query, book.Id)
	err := row.Scan(&bookIsTaken)
	if err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("book with id %d not found", book.Id)
		}

		return err
	}

	if !bookIsTaken {
		return fmt.Errorf("book with id %d is not taken", book.Id)
	}

	query = `UPDATE books SET istaken = $1 WHERE id = $2`
	_, err = ls.Base.Exec(query, false, book.Id)
	if err != nil {
		return err
	}

	return nil
}

func (ls *LibraryService) GetBookInfoByID(id int) (entities.Book, error) {
	query := `SELECT name, authorid, usertakenid, istaken FROM books WHERE id = $1`
	row := ls.Base.QueryRow(query, id)
	book := entities.Book{Id: id}
	var isTaken bool
	var takenBy int
	err := row.Scan(&book.Name, &book.AuthorID, &takenBy, &isTaken)
	if err != nil {
		return book, err
	}

	if isTaken {
		book.TakenBy = takenBy
	}

	author, err := ls.GetAuthorByID(book.AuthorID)
	if err != nil {
		return book, err
	}

	book.Author = author.Name
	return book, nil
}

func (ls *LibraryService) GetAuthorByID(id int) (entities.Author, error) {
	query := `SELECT name FROM authors WHERE id = $1`
	row := ls.Base.QueryRow(query, id)
	var author entities.Author
	err := row.Scan(&author.Name)
	if err != nil {
		return author, err
	}

	author.Id = id
	return author, nil
}

func (ls *LibraryService) GetUserByID(id int) (entities.User, error) {
	query := `SELECT username FROM users WHERE id = $1`
	row := ls.Base.QueryRow(query, id)
	var user entities.User
	err := row.Scan(&user.UserName)
	if err != nil {
		return user, err
	}

	user.Id = id
	return user, nil
}

func (ls *LibraryService) GetAllBooksTakenByUser(userID int) ([]entities.Book, error) {
	query := `SELECT id FROM books WHERE usertakenid = $1 AND istaken = $2`
	rows, err := ls.Base.Query(query, userID, true)
	books := []entities.Book{}
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var bookId int
		err = rows.Scan(&bookId)

		if err != nil {
			return nil, err
		}

		book, err := ls.GetBookInfoByID(bookId)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryService) GetAllUsers() ([]entities.User, error) {
	query := `SELECT id, username FROM users`
	rows, err := ls.Base.Query(query)
	users := []entities.User{}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user entities.User
		err = rows.Scan(&user.Id, &user.UserName)

		if err != nil {
			return nil, err
		}

		userBooks, err := ls.GetAllBooksTakenByUser(user.Id)
		if err != nil {
			return nil, err
		}

		user.BooksTaken = userBooks
		users = append(users, user)
	}

	return users, nil
}

func (ls *LibraryService) GetAllAuthorBooks(authorID int) ([]entities.Book, error) {
	query := `SELECT id FROM books WHERE authorid = $1`
	rows, err := ls.Base.Query(query, authorID)
	books := []entities.Book{}
	if err != nil {
		return books, err
	}

	for rows.Next() {
		var bookId int
		err = rows.Scan(&bookId)

		if err != nil {
			return nil, err
		}

		book, err := ls.GetBookInfoByID(bookId)
		if err != nil {
			return nil, err
		}

		books = append(books, book)
	}

	return books, nil
}

func (ls *LibraryService) GetAllAuthors() ([]entities.Author, error) {
	query := `SELECT id, name FROM authors`
	rows, err := ls.Base.Query(query)
	authors := []entities.Author{}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var author entities.Author
		err = rows.Scan(&author.Id, &author.Name)

		if err != nil {
			return nil, err
		}

		authorBooks, err := ls.GetAllAuthorBooks(author.Id)
		if err != nil {
			return nil, err
		}

		author.Books = authorBooks
		authors = append(authors, author)
	}

	return authors, nil
}

func (ls *LibraryService) HowManyAuthorsExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM authors`
	var authorsQuantity int
	err := ls.Base.QueryRow(query).Scan(&authorsQuantity)
	if err != nil {
		return 0, err
	}

	return authorsQuantity, nil
}

func (ls *LibraryService) HowManyBooksExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM books`
	var booksQuantity int
	err := ls.Base.QueryRow(query).Scan(&booksQuantity)
	if err != nil {
		return 0, err
	}

	return booksQuantity, nil
}

func (ls *LibraryService) HowManyUsersExist() (int, error) {
	query := `SELECT COUNT(*) AS row_count FROM users`
	var usersQuantity int
	err := ls.Base.QueryRow(query).Scan(&usersQuantity)
	if err != nil {
		return 0, err
	}

	return usersQuantity, nil
}
