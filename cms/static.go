package cms

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/helper/fp/slice"
)

type (
	Static struct {
		Tag        string    `json:"tag"`
		Text       string    `json:"value,omitempty"`
		Children   []*Static `json:"children,omitempty"`
		Attributes []string  `json:"attributes,omitempty"`
	}
)

func Render(data *Static) string {
	return transform(data).Render()
}

func transform(in *Static) gml.Tag {
	if len(in.Children) > 0 {
		children := slice.Map(in.Children, func(it *Static) gml.Tag {
			return transform(it)
		})

		return gml.New(in.Tag, gml.Tags(children...), in.Attributes...)
	} else if in.Text != "" {
		return gml.New(in.Tag, gml.Text(in.Text), in.Attributes...)
	} else {
		return gml.New(in.Tag, gml.Empty(), in.Attributes...)
	}
}
