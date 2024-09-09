package entities

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	AuthorID int    `json:"authorid"`
	TakenBy  int    `json:"takenby"`
}

type User struct {
	Id         int    `json:"id"`
	UserName   string `json:"username"`
	BooksTaken []Book `json:"bookstaken"`
}

type Author struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Books []Book `json:"books"`
}

type ErrorResponce struct {
	Message string `json:"message"`
}
