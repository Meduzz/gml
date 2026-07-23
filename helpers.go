package gml

import (
	"fmt"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Children []Tag
	TextTag  string
)

// Tags - create a tag out of multiple tags.
func Tags(children ...Tag) Tag {
	if children == nil {
		return nil
	}

	return Children(children)
}

// Text - create a tag out of a piece of text.
func Text(value string) Tag {
	return TextTag(value)
}

func Empty() Tag {
	return New("", nil)
}

func (c Children) Render() string {
	return slice.Fold(c, "", func(in Tag, agg string) string {
		if agg == "" {
			return in.Render()
		}

		return fmt.Sprintf("%s%s", agg, in.Render())
	})
}

func (t TextTag) Render() string {
	return string(t)
}
