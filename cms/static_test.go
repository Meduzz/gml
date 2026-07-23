package cms

import (
	"testing"
)

func TestFromHtml_SimpleTag(t *testing.T) {
	tag, err := FromHtml("<div>Hello</div>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<div>Hello</div>"
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_WithHtml(t *testing.T) {
	tag, err := FromHtml("<html><head><title>Hello!</title></head><body><h1>Hello!</h1></body></html>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<html><head><title>Hello!</title></head><body><h1>Hello!</h1></body></html>"
	if tag.Render() != expected {
		t.Errorf("expected %s, got: %s", expected, tag.Render())
	}
}

func TestFromHtml_NestedTags(t *testing.T) {
	tag, err := FromHtml("<div><span>Nested</span></div>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<div><span>Nested</span></div>"
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_VoidElements(t *testing.T) {
	tests := []struct {
		name   string
		html   string
		expect string
	}{
		{"br", "<br/>", "<br />"},
		{"hr", "<hr/>", "<hr />"},
		{"img", `<img src="test.png"/>`, `<img src="test.png" />`},
		{"input", `<input type="text"/>`, `<input type="text" />`},
		{"meta", `<meta charset="utf-8"/>`, `<meta charset="utf-8" />`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tag, err := FromHtml(tt.html)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if tag.Render() != tt.expect {
				t.Errorf("expected %q, got %q", tt.expect, tag.Render())
			}
		})
	}
}

func TestFromHtml_Atributes(t *testing.T) {
	tag, err := FromHtml(`<a href="https://example.com" class="link">Click</a>`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Token-based parser preserves source order
	expected := `<a href="https://example.com" class="link">Click</a>`
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_SingleAttribute(t *testing.T) {
	tag, err := FromHtml(`<input type="checkbox" checked/>`)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := `<input type="checkbox" checked="" />`
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_MultipleRoots(t *testing.T) {
	tag, err := FromHtml("<div>First</div><span>Second</span>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<div>First</div><span>Second</span>"
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_EmptyText(t *testing.T) {
	tag, err := FromHtml("<div>   </div>")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<div />"
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}

func TestFromHtml_ClosingTagsAutoClosed(t *testing.T) {
	tag, err := FromHtml("<div>Unclosed")
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	expected := "<div>Unclosed</div>"
	if tag.Render() != expected {
		t.Errorf("expected %q, got %q", expected, tag.Render())
	}
}
