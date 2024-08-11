package gml_test

import (
	"fmt"
	"testing"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/gml/logic"
	"github.com/Meduzz/gml/tags"
)

func TestTags(t *testing.T) {
	t.Run("h-func does black magic", func(t *testing.T) {
		subject := gml.H(`span#private.blue.underline(title="gml works" so="well")`, gml.Text("Tada"))
		println(subject.Render())

		t.Run("only a tag", func(t *testing.T) {
			subject := gml.H("p", gml.Empty())
			println(subject.Render())
		})

		t.Run("only an id", func(t *testing.T) {
			subject := gml.H("#identifier", gml.Empty())
			println(subject.Render())
		})

		t.Run("only a class", func(t *testing.T) {
			subject := gml.H(".blue", gml.Empty())
			println(subject.Render())
		})

		t.Run("only attributes", func(t *testing.T) {
			subject := gml.H("(attr=\"value\" value=\"attr\")", gml.Empty())
			println(subject.Render())
		})

		t.Run("an empty text tag will make script not self close", func(t *testing.T) {
			subject := gml.New("script", gml.Text(""))

			result := subject.Render()
			println(result)

			if result != "<script></script>" {
				t.Error("script tag was not left open")
			}
		})

		t.Run("empty br will self close", func(t *testing.T) {
			subject := gml.New("br", gml.Empty())

			result := subject.Render()
			println(result)

			if result != "<br />" {
				t.Error("br was not closed when empty")
			}
		})

		t.Run("if conditions are ... fun", func(t *testing.T) {
			subject := logic.If(true, gml.Text("Yes"), gml.Text("No"))

			result := subject.Render()
			println(result)

			if result != "Yes" {
				t.Error("result was not yes...")
			}

			t.Run("and the else?", func(t *testing.T) {
				subject := logic.If(false, gml.Text("Yes"), gml.Text("No"))

				result := subject.Render()
				println(result)

				if result != "No" {
					t.Error("result was not No...")
				}
			})
		})

		t.Run("Iterate stuff and other goodies", func(t *testing.T) {
			longList := []int{1, 2, 3}

			subject := logic.Iterate(longList, func(t int) gml.Tag { return tags.Li(gml.Text(fmt.Sprintf("%d", t))) }, nil)
			result := subject.Render()

			println(result)

			expected := "<li>1</li><li>2</li><li>3</li>"

			if result != expected {
				t.Error("the iteration went wrong!")
			}

			t.Run("and the else?", func(t *testing.T) {
				subject := logic.Iterate(nil, func(t int) gml.Tag { return gml.Empty() }, gml.Text("Im empty!"))
				result := subject.Render()

				println(result)

				expected := "Im empty!"

				if result != expected {
					t.Error("the (else) iteration went wrong!")
				}
			})
		})
	})
}
