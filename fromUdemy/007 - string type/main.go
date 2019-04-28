package main

import "fmt"

func main()  {
	a := "Hello Wold"

	fmt.Println(a)

	bs := []byte(a)

	fmt.Println(bs)

	for i := 0; i < len(a); i++ {
		fmt.Printf("%#U ", a[i])
	}

	fmt.Println("")

	for i, s := range (a) {
		fmt.Printf("%d - %x\n", i, s)
	}
}

