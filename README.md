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