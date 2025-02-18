package option

import (
	"errors"
	"fmt"
)

// Option represents an optional value that may or may not be present.
type Option[T any] struct {
	value *T
}

// Some creates an Option with a present value.
func Some[T any](v T) Option[T] {
	return Option[T]{value: &v}
}

// None creates an Option without a value (absent).
func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

// IsSome returns true if the Option contains a value.
func (o Option[T]) IsSome() bool {
	return o.value != nil
}

// IsNone returns true if the Option does not contain a value.
func (o Option[T]) IsNone() bool {
	return o.value == nil
}

// Unwrap returns the value or panics if the Option is None.
func (o Option[T]) Unwrap() T {
	if o.value == nil {
		panic("called `Unwrap()` on a `None` value")
	}
	return *o.value
}

// UnwrapOr returns the value or a default value if the Option is None.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.value == nil {
		return defaultValue
	}
	return *o.value
}

// UnwrapOrElse returns the value or calls a fallback function to generate a value.
func (o Option[T]) UnwrapOrElse(f func() T) T {
	if o.value == nil {
		return f()
	}
	return *o.value
}

// Expect returns the value or a custom error message if the Option is None.
func (o Option[T]) Expect(errMsg string) (T, error) {
	if o.value == nil {
		return *new(T), errors.New(errMsg)
	}
	return *o.value, nil
}

// Map applies a function to the contained value (if present) and returns a new Option with the result.
func Map[T, U any](opt Option[T], f func(T) U) Option[U] {
	if opt.IsNone() {
		return None[U]()
	}
	return Some(f(*opt.value))
}

// And returns None if the first Option is None, otherwise it returns the second Option.
func And[T, U any](opt Option[T], other Option[U]) Option[U] {
	if opt.IsNone() {
		return None[U]()
	}
	return other
}

// Or returns the first Option if it's Some, otherwise it returns the second Option.
func (o Option[T]) Or(opt Option[T]) Option[T] {
	if o.IsSome() {
		return o
	}
	return opt
}

// Filter returns the Option if the value satisfies the predicate, otherwise returns None.
func (o Option[T]) Filter(predicate func(T) bool) Option[T] {
	if o.IsSome() && predicate(*o.value) {
		return o
	}
	return None[T]()
}

// String returns a string representation of the Option.
func (o Option[T]) String() string {
	if o.IsSome() {
		return fmt.Sprintf("Some(%v)", *o.value)
	}
	return "None"
}
