package main

import "fmt"

// Create a type person
type person struct {
	first string
	last string
	age int
}

type secretAgent struct {
	person
	ltk bool
	first string
}

func main()  {
	person1 := person{"Jorge Luis", "Acosta", 28} // create a value of type person
	fmt.Println(person1)

	sa1 := secretAgent{
		person: person{
			"Antuanett", "Frometa", 27,
		},
		ltk: true,
		first: "Collision with person.first. It's not present, it use person.first automatically",
	}

	fmt.Println(sa1)
	fmt.Println(sa1.first, sa1.last) //sa1.last "it's a promoted field" we dont need do sa1.person.last
	fmt.Println(sa1.person.first)
}
