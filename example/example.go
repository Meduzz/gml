package main

import (
	"github.com/Meduzz/gml"
	. "github.com/Meduzz/gml/tags"
)

func main() {
	println(html.Render())
}

var html = Html(gml.Tags(Head(Title(gml.Text("Hello world!"))), Body(gml.Tags(H1(gml.Text("Hello world!")), Div(gml.Text("Welcome to the future?"))))))
