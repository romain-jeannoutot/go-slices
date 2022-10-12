package main

type FirstnameSpecification struct {
	firstname string
}

func NewFirstnameSpecification(firstname string) *FirstnameSpecification {
	return &FirstnameSpecification{firstname}
}

func (s *FirstnameSpecification) IsSatisfied(v any) bool {
	return v.(User).firstname == s.firstname
}
