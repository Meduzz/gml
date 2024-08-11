package logic

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/helper/fp/slice"
)

// If - evaluates condition and returns then or otherwise. If otherwise is nil, gml.Empty() is returned.
func If(condition bool, then gml.Tag, otherwise gml.Tag) gml.Tag {
	if otherwise == nil {
		otherwise = gml.Empty()
	}

	if condition {
		return then
	} else {
		return otherwise
	}
}

// Iterate - iterats the provided list. If list is empty then otherwise is returned. If otherwise is nil, then gml.Empty() is returned.
func Iterate[T any](list []T, each func(T) gml.Tag, otherwise gml.Tag) gml.Tag {
	if len(list) == 0 {
		if otherwise != nil {
			return otherwise
		}

		return gml.Empty()
	}

	result := slice.Map(list, func(t T) gml.Tag {
		return each(t)
	})

	return gml.Tags(result...)
}
