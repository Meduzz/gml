package gml

import (
	"regexp"
	"strings"

	"github.com/Meduzz/helper/fp/slice"
)

func H(hstring string, child Tag) Tag {
	// hstring ~= #id.class.class(attr)

	tag := &TagImpl{}
	tag.Child = child

	tag.Name = fetchTag(hstring)

	id := fetchId(hstring)

	if id != "" {
		tag.Attribute("id", id)
	}

	classes := fetchClasses(hstring)

	if len(classes) > 0 {
		tag.Attribute("class", strings.Join(classes, " "))
	}

	attrs := fetchAttrs(hstring)

	if len(attrs) > 0 {
		tag.Attributes = append(tag.Attributes, attrs...)
	}

	return tag
}

func fetchTag(hstring string) string {
	tagEnd := strings.IndexAny(hstring, "#.(")

	if tagEnd > -1 {
		if tagEnd == 0 {
			return "div"
		}

		return hstring[0:tagEnd]
	}

	if len(hstring) > 0 {
		return hstring
	}

	return "div"
}

func fetchId(hstring string) string {
	idStart := strings.Index(hstring, "#")

	if idStart > -1 {
		idEnd := strings.IndexAny(hstring, ".(")
		if idEnd > 0 {
			return hstring[idStart+1 : idEnd] // remove the #
		} else {
			return hstring[idStart+1:] // remove the #
		}
	}

	return ""
}

func fetchClasses(hstring string) []string {
	classesToKeep := make([]string, 0)
	classStart := strings.Index(hstring, ".")

	if classStart == -1 {
		return classesToKeep
	} else {
		hstring = hstring[classStart:]
	}

	classEnd := strings.Index(hstring, "(")

	if classEnd > 0 {
		hstring = hstring[:classEnd]
	}

	classes := strings.Split(hstring, ".")

	if len(classes) > 0 {
		slice.ForEach[string](classes, func(s string) {
			if len(s) == 0 {
				return
			}

			classesToKeep = append(classesToKeep, s)
		})
	}

	return classesToKeep
}

func fetchAttrs(hstring string) []string {
	attrStart := strings.Index(hstring, "(")

	if attrStart > -1 {
		hstring = hstring[attrStart+1:]
	} else {
		return make([]string, 0)
	}

	attrEnd := strings.Index(hstring, ")")

	if attrEnd > 0 {
		hstring = hstring[:attrEnd]
	}

	var rawAttrs []string

	r := regexp.MustCompile("\" ")

	if !r.MatchString(hstring) {
		rawAttrs = []string{hstring}
	} else {
		rawAttrs = r.Split(hstring, -1)
	}

	return slice.FlatMap(rawAttrs, func(attr string) []string {
		parts := strings.Split(attr, "=")

		return []string{parts[0], strings.Trim(parts[1], `"`)}
	})
}
