package form

import (
	"fmt"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/logic"
	"github.com/Meduzz/gml/tags"
)

func Select(label string, options map[any]any, attributes ...gml.Attribute) gml.Tag {
	attributes = append(attributes, attr.Class("select"))

	return FormControl(label, tags.Select(
		logic.Map(options, func(key any, value any) gml.Tag {
			return tags.Option(gml.Text(fmt.Sprintf("%v", value)), attr.Value(fmt.Sprintf("%v", key)))
		}, nil), attributes...,
	))
}
