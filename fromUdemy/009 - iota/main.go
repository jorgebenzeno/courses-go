package main

import "fmt"

const (
	_ = iota // iota = 0
	kb = 1 << (iota * 10) // iota = 1
	mb = 1 << (iota * 10) // iota = 2
	gb = 1 << (iota * 10) // iota = 3
)

const (
	base = iota + 1 // iota = 0
	kkb = base << 10
	mmb = kkb << 10
	ggb = mmb << 10
	other = iota // iota = 1
)

func main()  {
	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

	fmt.Printf("%d\t\t\t%b\n", kkb, kkb)
	fmt.Printf("%d\t\t\t%b\n", mmb, mmb)
	fmt.Printf("%d\t\t%b\n", ggb, ggb)
}
