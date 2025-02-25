package gml

import (
	"fmt"
)

type (
	Attribute []string
)

func Attributes(pairs ...string) Attribute {
	return pairs
}

func StringAttribute(key, value string) Attribute {
	return []string{key, value}
}

func AnyAttribute(key string, value any) Attribute {
	return []string{key, fmt.Sprintf("%v", value)}
}
