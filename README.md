# Option

A generic `Option[T]` type for Go, inspired by Rust's `Option<T>`. This package provides a way to handle optional values safely, avoiding null pointer dereferences and enabling functional programming patterns.

## Installation

```sh
go get github.com/mexirica/option
```

## Usage

### Importing the package

```go
package main

import (
    "fmt"
    "github.com/mexirica/option"
)

func main() {
    someValue := option.Some(42)    // Creates an Option with a value
    noneValue := option.None[int]() // Creates an Option without a value

    fmt.Println(someValue.IsSome()) // true
    fmt.Println(noneValue.IsNone()) // true
}
```

---

### Creating an Option

```go
some := option.Some("Hello, World!") // Creates an Option with a value
none := option.None[string]()        // Creates an Option without a value
```

---

### Checking if an Option has a value

```go
if some.IsSome() {
    fmt.Println("Option contains a value")
}

if none.IsNone() {
    fmt.Println("Option is empty")
}
```

---

### Unwrapping values

```go
value := some.Unwrap() // Returns the value (panics if None)
defaultValue := none.UnwrapOr("Default Value") // Returns the value or a default
fallbackValue := none.UnwrapOrElse(func() string {
    return "Generated Default"
})
fmt.Println(fallbackValue) // Output: Generated Default
```

---

### Handling missing values safely

```go
value, err := none.Expect("Expected a value, but found None")
if err != nil {
    fmt.Println(err)
}
```

---

### Transforming values with `Map`

```go
some := option.Some(21)
mapped := option.Map(some, func(x int) string {
    return fmt.Sprintf("Number: %d", x)
})
fmt.Println(mapped.Unwrap()) // Output: Number: 21
```

---

### Chaining Options with `And` and `Or`

```go
some := option.Some(42)
none := option.None[int]()

// And: Returns the second Option if the first is Some, otherwise returns None
andResult := option.And(some, option.Some("Hello"))
fmt.Println(andResult) // Output: Some(Hello)

andNone := option.And(none, option.Some("World"))
fmt.Println(andNone) // Output: None

// Or: Returns the first Option if it's Some, otherwise returns the second Option
orResult := none.Or(option.Some(100))
fmt.Println(orResult) // Output: Some(100)
```

---

### Filtering values

```go
some := option.Some(42)
filtered := some.Filter(func(x int) bool {
    return x > 40
})
fmt.Println(filtered) // Output: Some(42)

filteredNone := some.Filter(func(x int) bool {
    return x > 50
})
fmt.Println(filteredNone) // Output: None
```

---

## Methods and Functions

| Function / Method                       | Description |
|------------------------------------------|-------------|
| `Some(value)`                           | Creates an Option with a value |
| `None()`                                | Creates an empty Option |
| `IsSome()`                              | Returns `true` if the Option contains a value |
| `IsNone()`                              | Returns `true` if the Option is empty |
| `Unwrap()`                              | Returns the value or panics if None |
| `UnwrapOr(defaultValue)`                 | Returns the value or a default value |
| `UnwrapOrElse(func() T)`                 | Returns the value or calls a function to generate a value |
| `Expect(errMsg)`                        | Returns the value or an error if None |
| `Map(option, func(T) U)`                 | Transforms the value if present, returning a new `Option[U]` |
| `And(option, Option[U])`                 | Returns `None` if the first Option is `None`, otherwise returns the second Option |
| `Or(Option[T])`                         | Returns the first Option if it's `Some`, otherwise returns the second Option |
| `Filter(func(T) bool)`                  | Returns the Option if the value satisfies the predicate, otherwise returns `None` |

---

## Why use `Option`?

- **Safety:** Avoids null pointer dereferences by handling optional values explicitly.
- **Clarity:** Makes it clear when a value might be absent.
- **Functional Style:** Enables functional programming patterns like `Map`, `And`, `Or`, and `Filter`.

---

## License

This project is licensed under the MIT License.

---

[![Go Reference](https://pkg.go.dev/badge/github.com/mexirica/option.svg)](https://pkg.go.dev/github.com/mexirica/option)