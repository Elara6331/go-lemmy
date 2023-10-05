package lemmy

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

// Optional represents an optional value
type Optional[T any] struct {
	value *T
}

// NewOptional creates an optional with value v
func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
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

// SetNil sets the optional value to nil
func (o Optional[T]) SetNil() Optional[T] {
	o.value = nil
	return o
}

// IsValid returns true if the value of the optional is not nil
func (o Optional[T]) IsValid() bool {
	return o.value != nil
}

// Value returns the value in the optional.
func (o Optional[T]) Value() (T, bool) {
	if o.value != nil {
		return *o.value, true
	}
	return *new(T), false
}

// ValueOr returns the value inside the optional if it exists, or else it returns fallback
func (o Optional[T]) ValueOr(fallback T) T {
	if o.value != nil {
		return *o.value
	}
	return fallback
}

// ValueOrZero returns the value inside the optional if it exists, or else it returns the zero value of T
func (o Optional[T]) ValueOrZero() T {
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
