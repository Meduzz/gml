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
		Attributes []Attribute
	}
)

func New(name string, child Tag, attributes ...Attribute) Tag {
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
	t.Attributes = append(t.Attributes, StringAttribute(key, value))
}

func merge(data []Attribute) string {
	attribs := slice.Map(data, func(pair Attribute) string {
		if len(pair) == 2 {
			key := pair[0]
			value := pair[1]

			return fmt.Sprintf("%s=\"%s\"", key, value)
		} else if len(pair) == 0 {
			return ""
		} else {
			return pair[0]
		}
	})

	return strings.Join(attribs, " ")
}
