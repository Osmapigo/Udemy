package main

import "fmt"

func main() {
	defer defered()
	fmt.Println("I'll go first")
}

func defered() {
	fmt.Println("I should appear last")
}
