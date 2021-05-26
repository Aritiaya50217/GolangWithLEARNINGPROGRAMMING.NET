package entities

type Cat struct {
	Name string
}

func (cat Cat) Speak() string {
	return "Meo Meo"
}

func (cat Cat) Type() string {
	return cat.Name
}
