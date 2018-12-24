package main

import "fmt"

func main() {
	slice := []int{1, 24, 5, 6, 7, 88}
	a := foo(slice...)
	b := bar(slice)
	fmt.Println(a, b)
}

func foo(a ...int) int {
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}

func bar(b []int) int {
	total := 0
	for _, v := range b {
		total += v
	}
	return total
}
