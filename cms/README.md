## CMS

Perhaps the worst naming of the bunch.

But imagine a slate inspired model that allows you to define the page in one place and render it via `gml` in another.

### Definition

```go
	Static struct {
		Tag        string    `json:"tag"`
		Text       string    `json:"value,omitempty"`
		Children   []*Static `json:"children,omitempty"`
		Attributes []string  `json:"attributes,omitempty"`
	}
```

Named static, so there's no confusion as to how to template these things. Since you cant. But that also keeps the door open for a future variant called `Dynamic`... or something else.

### Bonus content

I figure I'll need to mix gml tags with good old html, so I made a function `FromHtml` that slurps html as a string and turns it into gml tags.