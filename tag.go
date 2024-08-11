package gml

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Tag interface {
		Render() string
	}

	TagImpl struct {
		Name       string
		Child      Tag
		Attributes []string
	}
)

func New(name string, child Tag, attributes ...string) Tag {
	return &TagImpl{
		Name:       name,
		Child:      child,
		Attributes: attributes,
	}
}

func (t *TagImpl) Render() string {
	attributes := merge(t.Attributes)

	if len(attributes) > 0 {
		if reflect.TypeOf(t.Child) == reflect.TypeOf(&EmptyTag{}) || t.Child == nil {
			return fmt.Sprintf("<%s %s />", t.Name, attributes)
		} else {
			child := t.Child.Render()
			return fmt.Sprintf("<%s %s>%s</%s>", t.Name, attributes, child, t.Name)
		}
	} else {
		if reflect.TypeOf(t.Child) == reflect.TypeOf(&EmptyTag{}) || t.Child == nil {
			return fmt.Sprintf("<%s />", t.Name)
		} else {
			child := t.Child.Render()
			return fmt.Sprintf("<%s>%s</%s>", t.Name, child, t.Name)
		}
	}
}

func (t *TagImpl) Attribute(key, value string) {
	t.Attributes = append(t.Attributes, key, value)
}

func merge(data []string) string {
	stage := make([]string, 0)

	length := len(data)

	for i := 0; i < length; i += 2 {
		pair := slice.Take(slice.Skip(data, i), 2)

		if len(pair) < 2 {
			continue
		}

		stage = append(stage, fmt.Sprintf(`%s="%s"`, pair[0], pair[1]))
	}

	return strings.Join(stage, " ")
}
