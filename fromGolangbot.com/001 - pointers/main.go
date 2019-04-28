package main

import "fmt"

// bad way
func change_1(arr *[5]int)  {
	arr[0] = 28
}

// good way, using slice
func change_2(arr []int)  {
	arr[0] = 28
}

func main()  {
	var arr = []int{1,2,3,4,5}
	fmt.Println(arr...)

	change_1(&arr)
	change_2(arr[:])
	change_2()
}
