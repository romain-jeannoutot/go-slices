package main

type Employee struct {
	User
	job string
}

func NewEmployee(firstname, lastname, job string) Employee {
	return Employee{NewUser(firstname, lastname), job}
}
