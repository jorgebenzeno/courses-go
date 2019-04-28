package main

import "fmt"

type person struct {
	firstName string
	lastName string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) speak() {
	fmt.Println("I am ", s.firstName, s.lastName)
}

func (p person) speak()  {
	fmt.Println("I am ", p.firstName, p.lastName)
}

// Everybody implementing all humain's methods is a "humain" too
type humain interface {
	speak()
}

func foobar(h humain)  {
	h.speak()
}

func main()  {
	s1 := secretAgent{
		person: person{
			"Jorge Luis",
			"Acosta",
		},
		ltk: true,
	}
	p1 := person{
		"Antuanett",
		"Frometa",
	}

	foobar(s1)
	foobar(p1)
}
