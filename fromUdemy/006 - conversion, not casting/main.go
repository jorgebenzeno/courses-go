package main

import "fmt"

var a int
type hotdog int
var b hotdog

func main()  {
	a = 42
	b = 5

	a = int(b) // conversion T(x) => T = type; x = value

	fmt.Println(a)
}
