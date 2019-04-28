package main

import (
	"github.com/jorgebenzeno/courses-go/fromGolangbot.com/002-structs/structs/computer"
	"fmt"
)

func main()  {
	var spec computer.Spec
	spec.Maker = "appel"
	spec.Price = 5000

	spec2 := computer.Spec{
		Maker: "appel",
		Price: 5000,
	}

	spec3 := computer.Spec{
		Maker: "Samsung",
		Price: 5000,
	}

	if spec == spec2 {
		fmt.Println("They're equals")
	}else {
		fmt.Println("They aren't equals")
	}

	if spec3 == spec2 {
		fmt.Println("They're equals")
	}else {
		fmt.Println("They aren't equals")
	}
}