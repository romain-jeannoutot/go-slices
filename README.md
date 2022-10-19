# Go slices

Functions for slices using Go generics and (sometimes) specification pattern.

[TOC]

## Documentation

Full documentation coming soon ...

## Install

```shell
$ go get github.com/romain-jeannoutot/go-slices
```

## Usage

Examples are based on [example](https://github.com/romain-jeannoutot/go-slices/tree/main/example) types, variables and functions definition.

### func Every

```go
func Every[T any](items []T, specification Specification) bool
```

The `Every()` func return `true` if all items satisfy specification. Else, return `false`.

*Example :*

```go
goslices.Every(users, NewHasFirstnameSpecification()) // true
goslices.Every(users, NewFirstnameSpecification("John")) // false
```

### func Filter

```go
func Filter[T any](items []T, specification Specification) []T
```

The `Filter()` func return a new slice with items satisfying specification.

*Example :*

```go
goslices.Filter(users, NewFirstnameSpecification("John")) // John Doe, John Travis
goslices.Filter(users, NewFirstnameSpecification("Mike")) // []
```

### func Find

```go
func Find[T any](items []T, specification Specification) T
```

The `Find()` func return the first item in slice satisfying specification. Else, return empty value.

*Example :*

```go
goslices.Find(users, NewFirstnameSpecification("John")) // John Doe
goslices.Find(users, NewFirstnameSpecification("Mike")) // Empty User
```
