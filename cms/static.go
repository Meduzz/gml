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

	rawHtml struct {
		html string
	}
)

func Render(data *Static) string {
	return Transform(data).Render()
}

func Transform(in *Static) gml.Tag {
	if len(in.Children) > 0 {
		children := slice.Map(in.Children, Transform)

		return gml.New(in.Tag, gml.Tags(children...), gml.Attributes(in.Attributes...))
	} else if in.Text != "" {
		return gml.New(in.Tag, gml.Text(in.Text), gml.Attributes(in.Attributes...))
	} else {
		return gml.New(in.Tag, gml.Empty(), gml.Attributes(in.Attributes...))
	}
}

// RawHtml - obviously bad, but also sometimes needed. Use for every day adventures.
func RawHtml(html string) gml.Tag {
	return &rawHtml{
		html: html,
	}
}

func (r *rawHtml) Render() string {
	return r.html
}
