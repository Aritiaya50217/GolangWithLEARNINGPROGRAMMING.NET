package entities

type Dog struct {
	Name string
}

func (dog Dog) Speak() string {
	return "Gau Gau"
}

func (dog Dog) Type() string {
	return dog.Name
}
