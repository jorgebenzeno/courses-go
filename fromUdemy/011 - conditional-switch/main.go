package main

import "fmt"

func main()  {
	switch  {
	case false:
		fmt.Println("this should not print")
	case (2 == 4):
		fmt.Println("this should not print")
	case (4 == 4):
		fmt.Println("print")
		fallthrough
	case (5 == 5):
		fmt.Println("this should print cause fallthrough")
	case (6 == 6):
		fmt.Println("this should not print cause no fallthrough")
	}

	switch "Bond" {
	case "M", "bond", "Bond":
		fmt.Println("should print Bond")
	case "qsd":
		fmt.Println("should not print")
	}
}