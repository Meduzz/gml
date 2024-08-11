package logic

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/helper/fp/slice"
)

// When - evaluates condition and returns then or otherwise. When otherwise is nil, gml.Empty() is returned.
func When(condition bool, then gml.Tag, otherwise gml.Tag) gml.Tag {
	if otherwise == nil {
		otherwise = gml.Empty()
	}

	if condition {
		return then
	} else {
		return otherwise
	}
}

// Slice - iterats the provided list. If list is empty then otherwise is returned. If otherwise is nil, then gml.Empty() is returned.
func Slice[T any](list []T, each func(T) gml.Tag, otherwise gml.Tag) gml.Tag {
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

// Map - iterate over a mpa. If list if empty then otherwise is returned. If otherwise is empty then gml.Empty() is returned.
func Map[T comparable, K any](data map[T]K, each func(T, K) gml.Tag, otherwise gml.Tag) gml.Tag {
	if len(data) == 0 {
		if otherwise != nil {
			return otherwise
		}

		return gml.Empty()
	}

	result := make([]gml.Tag, 0)

	for k, v := range data {
		result = append(result, each(k, v))
	}

	return gml.Tags(result...)
}
