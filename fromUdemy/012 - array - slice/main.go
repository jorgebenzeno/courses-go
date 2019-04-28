package main

import "fmt"

func main()  {
	var x [5]int // simples type
	fmt.Println(x)
	x[3] = 360
	fmt.Println(x)

	// y := []type{values} composite values
	y := []int{1,3,5,7,9}
	//fmt.Println(y...)
	fmt.Println()
	// range
	for i, value := range y {
		fmt.Printf("%v - %v\n", i, value)
	}
	// slicing a slice
	fmt.Println(y[0:3])

	z := []int{1,2,3}
	fmt.Println(z)
	z = append(z, y...)
	fmt.Println(z)
	z = append(z[2:], y[2:]...)
	fmt.Println(z)
}
