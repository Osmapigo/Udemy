package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		c <- 44
	}()
	//reading from populated channel
	v, ok := <-c
	fmt.Println(v, ok)

	close(c)
	//Pulling from empty channel
	v, ok = <-c
	fmt.Println(v, ok)
}
