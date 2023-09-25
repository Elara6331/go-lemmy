package types

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
)

var ErrOptionalEmpty = errors.New("optional value is empty")

type Optional[T any] struct {
	value *T
}

func NewOptional[T any](v T) Optional[T] {
	return Optional[T]{value: &v}
}

func NewOptionalPtr[T any](v *T) Optional[T] {
	return Optional[T]{value: v}
}

func NewOptionalNil[T any]() Optional[T] {
	return Optional[T]{}
}

func (o Optional[T]) Set(v T) Optional[T] {
	o.value = &v
	return o
}

func (o Optional[T]) SetPtr(v *T) Optional[T] {
	o.value = v
	return o
}

func (o Optional[T]) SetNil() Optional[T] {
	o.value = nil
	return o
}

func (o Optional[T]) IsValid() bool {
	return o.value != nil
}

func (o Optional[T]) MustValue() T {
	if o.value == nil {
		panic("optional value is nil")
	}
	return *o.value
}

func (o Optional[T]) Value() (T, error) {
	if o.value != nil {
		return *o.value, ErrOptionalEmpty
	}
	return *new(T), nil
}

func (o Optional[T]) ValueOr(fallback T) T {
	if o.value != nil {
		return *o.value
	}
	return fallback
}

func (o Optional[T]) ValueOrEmpty() T {
	if o.value != nil {
		return *o.value
	}
	var value T
	return value
}

func (o Optional[T]) MarshalJSON() ([]byte, error) {
	return json.Marshal(o.value)
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, []byte("null")) {
		o.value = nil
		return nil
	}

	o.value = new(T)
	return json.Unmarshal(b, o.value)
}

func (o Optional[T]) EncodeValues(key string, v *url.Values) error {
	s := o.String()
	if s != "<nil>" {
		v.Set(key, s)
	}
	return nil
}

func (o Optional[T]) String() string {
	if o.value == nil {
		return "<nil>"
	}
	return fmt.Sprint(*o.value)
}

func (o Optional[T]) GoString() string {
	if o.value == nil {
		return "nil"
	}
	return fmt.Sprintf("%#v", *o.value)
}
