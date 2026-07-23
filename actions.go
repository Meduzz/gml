package gml

type (
	Action interface {
		Execute(Tag)
	}

	visitorAction struct {
		needle Tag
		action func(Tag)
	}
)

var (
	_ Action = (*visitorAction)(nil)
)

func CreateAction(needle Tag, action func(Tag)) Action {
	return &visitorAction{
		needle: needle,
		action: action,
	}
}

func (v *visitorAction) Execute(root Tag) {
	for it := range root.All(v.needle) {
		v.action(it)
	}
}

func UpdateAction(target Tag) func(Tag) {
	it, ok := target.(*TagImpl)

	if !ok {
		return func(t Tag) {}
	}

	return func(t Tag) {
		other, ok := t.(*TagImpl)

		if ok {
			other.Name = it.Name
			other.Attributes = it.Attributes
			other.Child = it.Child
		}
	}
}

func RemoveAction(match Tag) {
	it, ok := match.(*TagImpl)

	if ok {
		it.Name = ""
	}
}
