package components

import "github.com/Meduzz/gml"

type (
	Component[T any]         func(T) gml.Tag
	KeyedComponent[T, K any] func(T, K) gml.Tag
)

func (c Component[T]) Render(data T) func() string {
	return c(data).Render
}

func (c KeyedComponent[T, K]) Render(key T, data K) func() string {
	return c(key, data).Render
}

// Box makes a Compontent[T] behave like a Component[K] with the help of handler: func(T) K.
func Box[T, K any](component Component[K], handler func(T) K) Component[T] {
	return func(data T) gml.Tag {
		return component(handler(data))
	}
}
