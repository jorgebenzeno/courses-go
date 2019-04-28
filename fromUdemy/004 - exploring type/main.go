package main

import "fmt"

var y = 42

// STATIC programming language
// a VARIABLE is DECLARED to hold a value of certain TYPE
var z = "Shared, not stirred"

func main()  {
	fmt.Println(y)

	fmt.Printf("%T\n", y)

	fmt.Println(z);
	fmt.Printf("%T\n", z)

	// z = 42
}
