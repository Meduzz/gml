package gml

import (
	"fmt"
	"iter"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
)

type (
	Tag interface {
		// Render render the tag
		Render() string
		// All return iterator of all tags matching provided tags (or all tags if nil)
		All(Tag) iter.Seq[Tag]
		// Match return TorF depending if tags match
		Match(Tag) bool
	}

	TagImpl struct {
		Name       string
		Child      Tag
		Attributes []Attribute
	}
)

var (
	_ Tag = (*TagImpl)(nil)
)

func New(name string, child Tag, attributes ...Attribute) Tag {
	return &TagImpl{
		Name:       name,
		Child:      child,
		Attributes: attributes,
	}
}

func (t *TagImpl) Render() string {
	// If Name is empty, render as empty string (even if attributes exist, they are ignored)
	if t.Name == "" {
		return ""
	}

	attributes := merge(t.Attributes)

	if t.Child == nil {
		if len(attributes) > 0 {
			return fmt.Sprintf("<%s %s />", t.Name, attributes)
		}
		return fmt.Sprintf("<%s />", t.Name)
	}

	child := t.Child.Render()
	if len(attributes) > 0 {
		return fmt.Sprintf("<%s %s>%s</%s>", t.Name, attributes, child, t.Name)
	}
	return fmt.Sprintf("<%s>%s</%s>", t.Name, child, t.Name)
}

func (t *TagImpl) Attribute(key, value string) {
	t.Attributes = append(t.Attributes, StringAttribute(key, value))
}

func (t *TagImpl) All(needle Tag) iter.Seq[Tag] {
	return func(yield func(Tag) bool) {
		if t.Name != "" {
			if needle == nil || t.Match(needle) {
				if !yield(t) {
					return
				}
			}

			if t.Child != nil {
				for it := range t.Child.All(needle) {
					if !yield(it) {
						break
					}
				}
			}
		}
	}
}

func (t *TagImpl) Match(other Tag) bool {
	it, ok := other.(*TagImpl)

	if !ok {
		return false
	}

	if t.Name != it.Name {
		return false
	}

	// Use the tag with fewer attributes as the "needle"
	needle, haystack := it, t
	if len(t.Attributes) < len(it.Attributes) {
		needle, haystack = t, it
	}

	matches := slice.Filter(needle.Attributes, func(pair Attribute) bool {
		return slice.Contains(haystack.Attributes, pair)
	})

	return len(matches) == len(needle.Attributes)
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
