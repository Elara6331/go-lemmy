package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var ErrOptionalEmpty = errors.New("optional value is empty")

// Optional represents an optional value
type Optional[T any] struct {
	value *T
}

// NewOptional creates an optional with value v
func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
}

// NewOptionalPtr creates an optional with the value pointed to by v
func NewOptionalPtr[T any](v *T) Optional[T] {
	return Optional[T]{value: v}
}

// NewOptionalNil creates a new nil optional value
func NewOptionalNil[T any]() Optional[T] {
	return Optional[T]{}
}

// Set sets the value of the optional
func (o Optional[T]) Set(v T) Optional[T] {
	o.value = &v
	return o
}

// SetPtr sets the value of the optional to the value bointed to by v
func (o Optional[T]) SetPtr(v *T) Optional[T] {
	o.value = v
	return o
}

// SetNil sets the optional value to nil
func (o Optional[T]) SetNil() Optional[T] {
	o.value = nil
	return o
}

// IsValid returns true if the value of the optional is not nil
func (o Optional[T]) IsValid() bool {
	return o.value != nil
}

// MustValue returns the value in the optional. It panics if the value is nil.
func (o Optional[T]) MustValue() T {
	if o.value == nil {
		panic("optional value is nil")
	}
	return *o.value
}

// Value returns the value in the optional.
func (o Optional[T]) Value() (T, error) {
	if o.value != nil {
		return *o.value, ErrOptionalEmpty
	}
	return *new(T), nil
}

// ValueOr returns the value inside the optional if it exists, or else it returns fallback
func (o Optional[T]) ValueOr(fallback T) T {
	if o.value != nil {
		return *o.value
	}
	return fallback
}

// ValueOrEmpty returns the value inside the optional if it exists, or else it returns the zero value of T
func (o Optional[T]) ValueOrEmpty() T {
	if o.value != nil {
		return *o.value
	}
	var value T
	return value
}

// MarshalJSON encodes the optional value as JSON
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

// UnmarshalJSON decodes JSON into the optional value
func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		o.value = nil
		return nil
	}

	o.value = new(T)
	return json.Unmarshal(b, o.value)
}

// EncodeValues encodes the optional as a URL query parameter
func (o Optional[T]) EncodeValues(key string, v *url.Values) error {
	s := o.String()
	if s != "<nil>" {
		v.Set(key, s)
	}
	return nil
}

// String returns the string representation of the optional value
func (o Optional[T]) String() string {
	if o.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*o.value)
}

// GoString returns the Go representation of the optional value
func (o Optional[T]) GoString() string {
	if o.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%#v", *o.value)
}
