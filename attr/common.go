package attr

import "github.com/Meduzz/gml"

func Href(url string) gml.Attribute {
	return gml.StringAttribute("href", url)
}

func Src(url string) gml.Attribute {
	return gml.StringAttribute("src", url)
}

func Type(typ string) gml.Attribute {
	return gml.StringAttribute("type", typ)
}

func Class(clazz string) gml.Attribute {
	return gml.StringAttribute("class", clazz)
}

func For(name string) gml.Attribute {
	return gml.StringAttribute("for", name)
}

func Name(name string) gml.Attribute {
	return gml.StringAttribute("name", name)
}

func Placeholder(text string) gml.Attribute {
	return gml.StringAttribute("placeholder", text)
}

func Title(title string) gml.Attribute {
	return gml.StringAttribute("title", title)
}

func Value(value any) gml.Attribute {
	return gml.AnyAttribute("value", value)
}

func Id(value string) gml.Attribute {
	return gml.StringAttribute("id", value)
}

func Disabled() gml.Attribute {
	return gml.AnyAttribute("disabled", true)
}

func Min(size int) gml.Attribute {
	return gml.AnyAttribute("min", size)
}

func Max(size int) gml.Attribute {
	return gml.AnyAttribute("max", size)
}

func Size(size int) gml.Attribute {
	return gml.AnyAttribute("size", size)
}

func Step(step int) gml.Attribute {
	return gml.AnyAttribute("step", step)
}

func Maxlength(length int) gml.Attribute {
	return gml.AnyAttribute("maxlength", length)
}

func Minlength(length int) gml.Attribute {
	return gml.AnyAttribute("minlength", length)
}

func Pattern(regex string) gml.Attribute {
	return gml.StringAttribute("pattern", regex)
}
