package main

import (
	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/attr"
	"github.com/Meduzz/gml/cms"
	"github.com/Meduzz/gml/components"
	. "github.com/Meduzz/gml/tags"
)

func main() {
	// these all print the same thing.
	println(html.Render())
	println(component(&Greeting{"Hello world!", "Welcome to the future?"}).Render())
	println(cms.Render(cmsJson))
}

// plain gml.Tag
var html = Html(gml.Tags(Head(Title(gml.Text("Hello world!"))), Body(gml.Tags(H1(gml.Text("Hello world!"), attr.Class("blink")), Div(gml.Text("Welcome to the future?"))))))

// cms generated json
var cmsJson = &cms.Static{
	Tag: "html",
	Children: []*cms.Static{
		{
			Tag: "head",
			Children: []*cms.Static{
				{
					Tag:  "title",
					Text: "Hello world!",
				},
			},
		}, {
			Tag: "body",
			Children: []*cms.Static{
				{
					Tag:  "h1",
					Text: "Hello world!",
					Attributes: []string{
						"class",
						"blink",
					},
				}, {
					Tag:  "div",
					Text: "Welcome to the future?",
				},
			},
		},
	},
}

type Greeting struct {
	Greeting string
	Message  string
}

// gml component
var component = func(data *Greeting) gml.Tag {
	return Html(gml.Tags(Head(Title(gml.Text(data.Greeting))), Body(gml.Tags(H1(gml.Text(data.Greeting), attr.Class("blink")), Div(gml.Text(data.Message))))))
}

var (
	// "proof" that component is of type Component[*Greeting]
	_ components.Component[*Greeting] = component
)
