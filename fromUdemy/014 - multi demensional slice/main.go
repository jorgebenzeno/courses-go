package main

import "fmt"

func main()  {
	x := []int{1,2}
	y := []int{3,4}
	z := [][]int{x, y}

	fmt.Println(z)
}
