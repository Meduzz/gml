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
		slice.ForEach(classes, func(s string) {
			if len(s) == 0 {
				return
			}

			classesToKeep = append(classesToKeep, s)
		})
	}

	return classesToKeep
}

func fetchAttrs(hstring string) []Attribute {
	attrStart := strings.Index(hstring, "(")

	if attrStart > -1 {
		// drop everything before (
		hstring = hstring[attrStart+1:]
	} else {
		return []Attribute{}
	}

	attrEnd := strings.Index(hstring, ")")

	if attrEnd > 0 {
		hstring = hstring[:attrEnd]
	}

	// the easy way out
	// return []Attribute{hstring}

	// Define the regular expression to match attribute-value pairs
	re := regexp.MustCompile(`(\w+)(?:="([^"]*)")?`)

	// Find all matches in the input string
	matches := re.FindAllStringSubmatch(hstring, -1)

	return slice.Map(matches, func(attr []string) Attribute {
		if len(attr) == 3 {
			return StringAttribute(attr[1], attr[2])
		} else if len(attr) == 2 {
			return Attribute{attr[1]}
		} else {
			return Attribute{}
		}
	})
}
