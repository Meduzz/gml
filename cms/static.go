package cms

import (
	"fmt"
	"io"
	"strings"

	"golang.org/x/net/html"

	"github.com/Meduzz/gml"
	"github.com/Meduzz/helper/fp/slice"
)

type Static struct {
	Tag        string    `json:"tag"`
	Text       string    `json:"value,omitempty"`
	Children   []*Static `json:"children,omitempty"`
	Attributes []string  `json:"attributes,omitempty"`
}

func Render(data *Static) string {
	return Transform(data).Render()
}

func Transform(in *Static) gml.Tag {
	if len(in.Children) > 0 {
		children := slice.Map(in.Children, Transform)
		if in.Tag == "" {
			// For multiple roots (document node), return children directly without wrapper
			return gml.Tags(children...)
		}
		return gml.New(in.Tag, gml.Tags(children...), gml.Attributes(in.Attributes...))
	} else if in.Text != "" {
		return gml.New(in.Tag, gml.Text(in.Text), gml.Attributes(in.Attributes...))
	} else {
		return gml.New(in.Tag, nil, gml.Attributes(in.Attributes...))
	}
}

// FromHtml parses HTML string and returns a gml.Tag representing the parsed HTML.
// It uses the html.Token API (stream-based parsing) instead of html.Parse.
// It returns an error if the HTML cannot be parsed.
func FromHtml(htmlStr string) (gml.Tag, error) {
	z := html.NewTokenizer(strings.NewReader(htmlStr))
	static, err := tokensToStatic(z)
	if err != nil {
		return nil, err
	}
	return Transform(static), nil
}

// tokensToStatic converts a stream of HTML tokens into a *Static structure
// using a stack-based approach to track nesting.
func tokensToStatic(z *html.Tokenizer) (*Static, error) {
	var stack []*Static
	var roots []*Static

	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			return nil, fmt.Errorf("failed to parse HTML: %w", z.Err())
		}

		token := z.Token()

		switch token.Type {
		case html.TextToken:
			text := strings.TrimSpace(token.Data)
			if text != "" && len(stack) > 0 {
				stack[len(stack)-1].Text = text
			}

		case html.StartTagToken:
			s := &Static{
				Tag:        strings.ToLower(token.Data),
				Attributes: tokenAttrToFlat(token.Attr),
			}
			if len(stack) > 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, s)
			} else {
				roots = append(roots, s)
			}
			stack = append(stack, s)

		case html.SelfClosingTagToken:
			s := &Static{
				Tag:        strings.ToLower(token.Data),
				Attributes: tokenAttrToFlat(token.Attr),
			}
			if len(stack) > 0 {
				stack[len(stack)-1].Children = append(stack[len(stack)-1].Children, s)
			} else {
				roots = append(roots, s)
			}
			// Do not push self-closing tags onto the stack

		case html.EndTagToken:
			// Pop from stack to match the opening tag
			if len(stack) > 0 {
				stack = stack[:len(stack)-1]
			}
		}
	}

	// If we have multiple root elements, wrap them as children without a tag name
	// (matching the original html.Parse behavior for DocumentNode)
	if len(roots) == 1 {
		return roots[0], nil
	} else if len(roots) > 1 {
		return &Static{
			Tag:      "",
			Children: roots,
		}, nil
	}

	return nil, nil
}

// tokenAttrToFlat converts html.Attribute slice to a flat []string (key1, value1, key2, value2...).
// This flat format is stored in Static for JSON serialization.
func tokenAttrToFlat(attrs []html.Attribute) []string {
	if len(attrs) == 0 {
		return nil
	}
	result := make([]string, 0, len(attrs)*2)
	for _, attr := range attrs {
		result = append(result, attr.Key, attr.Val)
	}
	return result
}
