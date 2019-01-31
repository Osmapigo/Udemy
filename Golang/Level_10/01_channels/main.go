package main

import (
	"fmt"
)

func main() {
	withBuffer()
	withFuncLiteral()
}

func withBuffer() {
	c := make(chan int, 1)
	c <- 42
	fmt.Println(<-c)
}

func withFuncLiteral() {
	c := make(chan int)
	go func() {
		c <- 43
	}()
	fmt.Println(<-c)
}
