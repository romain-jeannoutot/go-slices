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

	jobs := []string{
		"developer",
		"astronaut",
		"singer",
	}

	// John Doe, John Travis
	log.Println(goslices.Filter(users, NewFirstnameSpecification("John")))

	// Marcos Christian
	log.Println(goslices.Find(users, NewFirstnameSpecification("Marcos")))

	// Empty user
	log.Println(goslices.Find(users, NewFirstnameSpecification("Mike")))

	for i := 0; i < 5; i++ {
		log.Println(goslices.Rand(jobs))
	}
}
