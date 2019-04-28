package main

import "fmt"

func main()  {
	x := make([]int, 10, 12)
	fmt.Println(x)
	fmt.Println(cap(x))

	x = append(x, 12)
	x = append(x, 13)
	fmt.Println(x)

	x = append(x, 14) // bigger than 12, so it duplicates capacity
	fmt.Println(x)
	fmt.Println(cap(x))
	fmt.Println(len(x))
}
