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

	// Marcos Christian, 3
	spec := NewFirstnameSpecification("Marcos")
	log.Println(goslices.Find(users, spec), goslices.FindIndex(users, spec))

	// Empty user, -1
	spec = NewFirstnameSpecification("Mike")
	log.Println(goslices.Find(users, spec), goslices.FindIndex(users, spec))

	// All users with random job
	log.Println(goslices.Map(users, func(user User) Employee {
		return NewEmployee(user.firstname, user.lastname, goslices.Rand(jobs))
	}))

	// true
	log.Println(goslices.Includes(jobs, "developer"))

	// false
	log.Println(goslices.Includes(jobs, "postman"))
}
