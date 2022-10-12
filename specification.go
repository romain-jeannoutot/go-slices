package goslices

type Specification interface {
	IsSatisfied(v any) bool
}
