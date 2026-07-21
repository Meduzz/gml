package form

import (
	"fmt"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/tags"
)

func Text(typ, label string, attributes ...gml.Attribute) gml.Tag {
	attributes = append(attributes, attr.Class("input"), attr.Type(typ))

	return FormControl(label, tags.Input(gml.Empty(), attributes...))
}

func Checkbox(label string, attributes ...gml.Attribute) gml.Tag {
	attributes = append(attributes, attr.Class("checkbox"), attr.Type("checkbox"))

	return embedded(label, tags.Input(gml.Empty(), attributes...))
}

func Radio(label string, attributes ...gml.Attribute) gml.Tag {
	attributes = append(attributes, attr.Class("radio"), attr.Type("radio"))

	return embedded(label, tags.Input(gml.Empty(), attributes...))
}

func embedded(label string, child gml.Tag) gml.Tag {
	return tags.Div(tags.Label(gml.Tags(
		child,
		gml.Text(fmt.Sprintf("%s:", label)),
	), attr.Class("label")), attr.Class("control"))
}
