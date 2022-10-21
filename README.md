# Go slices

Functions for slices using Go generics and (sometimes) specification pattern.

- [Go slices](#go-slices)
  - [Documentation](#documentation)
  - [Install](#install)
  - [Usage](#usage)
    - [func Every](#func-every)
    - [func Filter](#func-filter)
    - [func Find](#func-find)
    - [func FindIndex](#func-findindex)
    - [func Includes](#func-includes)
    - [func Map](#func-map)
    - [func Rand](#func-rand)
    - [func Some](#func-some)

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

The `Every()` func returns `true` if all items satisfy specification. Else, return `false`.

*Example :*

```go
goslices.Every(users, NewHasFirstnameSpecification()) // true
goslices.Every(users, NewFirstnameSpecification("John")) // false
```

### func Filter

```go
func Filter[T any](items []T, specification Specification) []T
```

The `Filter()` func returns a new slice with items satisfying specification.

*Example :*

```go
goslices.Filter(users, NewFirstnameSpecification("John")) // John Doe, John Travis
goslices.Filter(users, NewFirstnameSpecification("Mike")) // []
```

### func Find

```go
func Find[T any](items []T, specification Specification) T
```

The `Find()` func returns the first item in slice satisfying specification. Else, return empty value.

*Example :*

```go
goslices.Find(users, NewFirstnameSpecification("John")) // John Doe
goslices.Find(users, NewFirstnameSpecification("Mike")) // Empty User
```

### func FindIndex

```go
func FindIndex[T any](items []T, specification Specification) int
```

The `FindIndex()` func returns index of the first item in slice satisfying specification. Else, return -1.

*Example :*

```go
goslices.FindIndex(users, NewFirstnameSpecification("John")) // 0
goslices.FindIndex(users, NewFirstnameSpecification("Mike")) // -1
```

### func Includes

```go
func Includes[T comparable](items []T, searchElement T) bool
```

The `Includes()` func returns `true` if at least one item in slice is equal to `searchElement`. Else, return `false`.

*Example :*

```go
goslices.Includes(jobs, "developer") // true
goslices.Includes(jobs, "postman") // false
```

### func Map

```go
func Map[T any, S any](items []T, callback func(item T, idx int, items []T) S) []S
```

The `Map()` func creates a new slice populated with elements returned by function called on each slice element.

*Example :*

```go
goslices.Map(users, func(user User, _ int, _ []User) Employee {
  return NewEmployee(user.firstname, user.lastname, goslices.Rand(jobs))
}) // []Employee
```

### func Rand

```go
func Rand[T any](items []T) T
```

The `Rand()` func returns a random element from slice. Empty value if slice is empty.

*Example :*

```go
goslices.Rand(jobs) // string
```

### func Some

```go
func Some[T any](items []T, specification Specification) bool
```

The `Some()` func returns `true` if at least one item satisfy specification. Else, return `false`.

*Example :*

```go
goslices.Some(users, NewFirstnameSpecification("John")) // true
goslices.Some(users, NewFirstnameSpecification("Mike")) // false
```
