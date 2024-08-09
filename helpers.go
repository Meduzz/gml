package gml

import (
	"fmt"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Children []Tag
	TextTag  string
	EmptyTag struct{}
)

// Tags - create a tag out of multiple tags.
func Tags(children ...Tag) Tag {
	return Children(children)
}

// Text - create a tag out of a piece of text.
func Text(value string) Tag {
	return TextTag(value)
}

func Empty() Tag {
	return &EmptyTag{}
}

func (c Children) Render() string {
	return slice.Fold(c, "", func(in Tag, agg string) string {
		if agg == "" {
			return in.Render()
		}

		return fmt.Sprintf("%s\n%s", agg, in.Render())
	})
}

func (t TextTag) Render() string {
	return string(t)
}

func (e *EmptyTag) Render() string {
	return ""
}
