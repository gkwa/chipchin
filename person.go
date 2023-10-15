package chipchin

type Person struct {
	FirstName string
	LastName  string
}

func (p Person) FullName() string {
	return p.FirstName + " " + p.LastName
}

func (p *Person) ChangeLastName(newLastName string) {
	p.LastName = newLastName
}
