package entities

type Duck struct {
	Name string
}

func (duck Duck) Speak() string {
	return "Cap Cap"
}

func (duck Duck) Type() string {
	return duck.Name
}
