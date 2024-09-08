package main

type User struct {
	ID       int
	Username string
	Email    string
	Role     string
}

type UserOption func(*User)

func NewUser(id int, opts ...UserOption) *User {
	user := &User{
		ID: id,
	}
	for _, opt := range opts {
		opt(user)
	}
	return user
}

func WithUsername(username string) UserOption {
	return func(u *User) {
		u.Username = username
	}
}

func WithEmail(email string) UserOption {
	return func(u *User) {
		u.Email = email
	}
}

func WithRole(role string) UserOption {
	return func(u *User) {
		u.Role = role
	}
}
