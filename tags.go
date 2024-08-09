package gml

func Div(child Tag, attributes ...string) Tag {
	return New("div", child, attributes...)
}

func P(child Tag, attributes ...string) Tag {
	return New("p", child, attributes...)
}

func Span(child Tag, attributes ...string) Tag {
	return New("span", child, attributes...)
}
