package main
import "strc"

import "fmt"

func main()  {
	p1 := struct {
		first string
	}{
		"Jorge Luis",
	}

	fmt.Println(p1)
}