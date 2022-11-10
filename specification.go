package goslices

// Specification represents a rule implementation.
type Specification interface {
	IsSatisfied(v any) bool
}
