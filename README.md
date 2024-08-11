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

### Empty tags
If tags are fed a `gml.Empty()`, then they will close directly, ie become `<script />` instead of `<script></script>`. To keep them "open" feed anything else in there, like an empty text: `gml.Text("")`.