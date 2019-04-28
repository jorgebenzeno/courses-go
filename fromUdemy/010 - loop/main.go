package main

import "fmt"

/*
There's not While statement
 */

func main()  {
	for i := 33; i < 122; i++ {
		fmt.Printf("%d\t\t%#x\t\t%#U\n", i,i,i)
	}

	for i := 0;; {
		fmt.Println(i)

		if i == 10 {
			break
		}

		i++
	}

	i:=0
	for {
		fmt.Println(i)

		if i == 10 {
			break
		}

		i++
	}
}
