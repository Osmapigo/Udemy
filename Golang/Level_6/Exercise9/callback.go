package main

import "fmt"

func main() {
	foo(func() {
		fmt.Println("this is the callback")
	})
}

func foo(f func()) {
	fmt.Println("This is foo")
	f()
}
