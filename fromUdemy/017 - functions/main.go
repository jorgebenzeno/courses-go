package main

import "fmt"

func main()  {
	a, b := mouse()
	fmt.Println(a, b);

	// Unfurling slice
	xi := []int{2,3,4,5}
	x := sum(xi...)
	fmt.Println(x)

	// Defer
}

func sum(v...int) int {
	sum := 0
	for v,_ := range v {
		sum += v
	}

	return sum
}

func foo()  {
	fmt.Println("foo")
}

func woo() string {
	return fmt.Sprint("Hello")
}

func mouse() (string, bool)  {
	return "hello", true
}