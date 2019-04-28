package main

import "fmt"

const a = 42
const b = 42.78 // const untype in declaration
const c string = "James Bond" // const of type in declaration

const (
	d = 42
	e = 42.5
)

const (
	x = iota
	y = iota
	z
)

const (
	f = iota
	g
	h
)

type int2 int

func main()  {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	fmt.Println(d)
	fmt.Println(e)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)

	const ff = 5
	const fff int = 10
	var intt int2 = 5
	intt = ff // Does work because 'ff' is a const of untyped type
	// intt = fff // Does not work because 'fff' is a const of int type
	fmt.Println(intt)
}