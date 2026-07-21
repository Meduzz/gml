package form

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/tags"
)

func Textarea(label string, value gml.Tag, attributes ...gml.Attribute) gml.Tag {
	return FormControl(label, tags.Textarea(value, attributes...))
}
