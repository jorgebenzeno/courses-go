package main

import "fmt"

func main()  {
	m := map[string]int {
		"Jame": 32,
		"Luis": 28,
	}
	fmt.Println(m)
	fmt.Println(m["Luis"])
	fmt.Println(m["J"]) // It returns "0" value

	v, ok := m["Luis"]
	fmt.Printf("%v %v\n", v, ok)
	v, ok = m["Louis"]
	fmt.Printf("%v %v\n", v, ok)

	if v, ok := m["Luis"]; ok {
		fmt.Printf("it's ok value %v\n", v)
	}

	m["Louis"] = 28
	for k,v := range m {
		fmt.Printf("Key %v Value %v\n", k, v)
	}

	fmt.Println("")
	delete(m, "Louis")
	fmt.Println(m)
}

