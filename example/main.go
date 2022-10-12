package main

import (
	"log"

	goslices "github.com/romain-jeannoutot/go-slices"
)

func main() {
	users := []User{
		NewUser("John", "Doe"),
		NewUser("Aliyah", "Decker"),
		NewUser("Arnas", "Willis"),
		NewUser("Marcos", "Christian"),
		NewUser("John", "Travis"),
	}

	log.Println(goslices.Filter(users, NewFirstnameSpecification("John")))
}
