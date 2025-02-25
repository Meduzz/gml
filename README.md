# gml
Pure go approach to html templates.

Ie:

```
import (
    "github.com/Meduzz/gml/tags"
    "github.com/Meduzz/gml"
)

func main() {
    var tmpl = tags.Div(gml.Text("Hello world!"))
    println(tmpl.Render())
}
```

Results in:

```
<div>Hello world!</div>
```

Did you see the [example](example/) package already?

## Common tags

Looking for commont tags? Look no further than the [tags](tags/) package.

## Common attributes

Looking for common attributes? Look no further than the [attr](attr/) package.

## Components

This lib just keep checking checkboxes, doesnt it? Have a look at the [components](components/) package for more details.

## HTMX attributes?

Sure thing, have a look at [htmx](htmx/) package.

## Ifs and butts

In the ["logic"](logic/) package you find helpers to iterate slices and maps. It also contains a helper for `if then else` situations.

## "I like to inflict pain on myself"

We go you covered, I think you will like the [cms](cms/) package. There we keep an experimental way to generate html from a ... lets call it Slate inspired model/api.

## Empty tags (aka known issues)
If tags are fed a `gml.Empty()`, then they will close directly, ie become `<script />` instead of `<script></script>`. To keep them "open" feed anything else in there, like an empty text: `gml.Text("")`.

Arguments could be made that feeding a tag `nil` should make it close instantly while `gml.Empty()` would not. We'll see in a future release.
