package main

import "fmt"

func main() {
	fmt.Println(foo())
	fmt.Println(bar())
}

func foo() int {
	return 25
}

func bar() (int, string) {
	return 55, "Oscar"
}
