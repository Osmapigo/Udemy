package main

import "fmt"

func main() {
	fmt.Println(fibonnaci(5))

}

func fibonnaci(n int) int {
	if n <= 2 {
		return 1
	}
	return fibonnaci(n-1) + fibonnaci(n-2)
}
