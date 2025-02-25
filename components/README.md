## Components

This is an attempt at making re-usable gml tags. We'll see how usefull they become.

### Definition

```go
type (
	Component[T any]         func(T) gml.Tag
	KeyedComponent[T, K any] func(T, K) gml.Tag
)
```

If you understand the basics of components from pesky js libs, then this prolly makes sense. The focus is on re-use and typing... and Go.