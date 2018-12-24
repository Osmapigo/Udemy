package main

import "fmt"

func main() {
	f := bar(3)
	fmt.Println(f())
}

func bar(x int) func() int {
	return func() int {
		return x * x
	}
}
