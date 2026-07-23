package gml

import (
	"fmt"
	"iter"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Children []Tag
	TextTag  string
)

var (
	_ Tag = (Children{})
	_ Tag = (TextTag)("")
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

func (c Children) All() iter.Seq[Tag] {
	return func(yield func(Tag) bool) {
		for _, it := range c {
			for cit := range it.All() {
				if !yield(cit) {
					break
				}
			}
		}
	}
}

func (c Children) Match(other Tag) bool {
	return false
}

func (t TextTag) Render() string {
	return string(t)
}

func (t TextTag) All() iter.Seq[Tag] {
	return func(yield func(Tag) bool) {
		yield(t)
	}
}

func (t TextTag) Match(other Tag) bool {
	return false
}
