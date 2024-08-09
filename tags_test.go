package gml_test

import (
	"testing"

	"github.com/Meduzz/gml"
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
	})
}
