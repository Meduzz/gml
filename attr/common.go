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

func Action(action string) gml.Attribute {
	return gml.StringAttribute("action", action)
}

func Alpha() gml.Attribute {
	return gml.AnyAttribute("alpha", true)
}

func MinLength(size int) gml.Attribute {
	return gml.AnyAttribute("minlength", size)
}

func MaxLength(size int) gml.Attribute {
	return gml.AnyAttribute("maxlength", size)
}

func Multiple() gml.Attribute {
	return gml.AnyAttribute("multiple", true)
}

func Rel(value string) gml.Attribute {
	return gml.StringAttribute("rel", value)
}

func Async() gml.Attribute {
	return gml.AnyAttribute("async", true)
}

func Module() gml.Attribute {
	return gml.AnyAttribute("module", true)
}

func Defer() gml.Attribute {
	return gml.AnyAttribute("defer", true)
}

func Autocomplete(value string) gml.Attribute {
	return gml.StringAttribute("autocomplete", value)
}

func Checked(value bool) gml.Attribute {
	return gml.AnyAttribute("checked", value)
}

func Contenteditable(value bool) gml.Attribute {
	return gml.AnyAttribute("contenteditable", value)
}

func Rows(value int) gml.Attribute {
	return gml.AnyAttribute("rows", value)
}

func Method(value string) gml.Attribute {
	return gml.StringAttribute("method", value)
}

func Content(value string) gml.Attribute {
	return gml.StringAttribute("content", value)
}

func Media(value string) gml.Attribute {
	return gml.StringAttribute("media", value)
}

func As(value string) gml.Attribute {
	return gml.StringAttribute("as", value)
}

func Sizes(value string) gml.Attribute {
	return gml.StringAttribute("sizes", value)
}

func Charset(value string) gml.Attribute {
	return gml.StringAttribute("charset", value)
}
