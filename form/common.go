package form

import (
	"fmt"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/tags"
)

func FormControl(label string, child gml.Tag) gml.Tag {
	return tags.Div(gml.Tags(
		tags.Label(gml.Text(fmt.Sprintf("%s:", label)), attr.Class("label")),
		child,
	), attr.Class("control"))
}

func Form(method, action string, form gml.Tag) gml.Tag {
	return tags.Form(form, gml.StringAttribute("method", method), attr.Action(action))
}
