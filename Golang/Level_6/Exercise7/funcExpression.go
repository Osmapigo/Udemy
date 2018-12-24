package main

import "fmt"

func main() {
	af := func(x int) int {
		return x * 2
	}

	fmt.Println(af(4))
}
