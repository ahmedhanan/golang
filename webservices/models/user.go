package models

import "errors"

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

func GetUsers() []*User {
	return users
}

func AddNewUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, errors.New("failed to add a new user")
}