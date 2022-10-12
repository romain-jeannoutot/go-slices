package main

type User struct {
	firstname, lastname string
}

func NewUser(firstname, lastname string) User {
	return User{firstname, lastname}
}
