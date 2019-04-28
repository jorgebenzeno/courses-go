package main

import (
	"fmt"
)

func main()  {
	a()
	b()
	fmt.Println(c())

	fmt.Println("\n")

	f()
	fmt.Println("Returned normally from f.")
}

// A deferred function's arguments are evaluated when the defer statement is evaluated.
// Returns 1 0
func a() {
	i := 0
	defer fmt.Println(i)

	i++
	fmt.Println(i);

	return
}

// . Deferred function calls are executed in Last In First Out order after the surrounding function returns.
// Returns 3 2 1 0
func b() {

	defer fmt.Println("\n")
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

//Deferred functions may read and assign to the returning function's named return values.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

// Panic is a built-in function that stops the ordinary flow of control and begins panicking. When the function F calls panic, execution of F stops, any deferred functions in F are executed normally, and then F returns to its caller. To the caller, F then behaves like a call to panic. The process continues up the stack until all functions in the current goroutine have returned, at which point the program crashes. Panics can be initiated by invoking panic directly. They can also be caused by runtime errors, such as out-of-bounds array accesses.
//
//Recover is a built-in function that regains control of a panicking goroutine. Recover is only useful inside deferred functions. During normal execution, a call to recover will return nil and have no other effect. If the current goroutine is panicking, a call to recover will capture the value given to panic and resume normal execution.
func f() (i int) {
	defer func() {
		fmt.Println("Recovered in f ...***")
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()

	defer func() {
		fmt.Println("Recovered in f ...")
	}()

	fmt.Println("Calling g.")
	g(0)

	// It's not executed, g() will panic so, deffers was already executed when g() was in panic
	defer func() {
		fmt.Println("last Recovered in f ....*", i)
	}()

	fmt.Println("Returned normally from g.")

	return 1
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
