package main

import "fmt"

type person struct {
	first string
	last string
}

type secretAgent struct {
	person
	ltk bool
}

func (s secretAgent) fullName() string  {
	return fmt.Sprintf("%s %s", s.first, s.last)
}

func main()  {
	sa1 := secretAgent{
		person: person{
			"Jorge Luis", "Acosta",
		},
		ltk: true,
	}

	fmt.Println(sa1)
	fmt.Println(sa1.fullName())
}
