package main

import (
	"fmt"
	"runtime"
)

func main() {
	c := make(chan int)
	fmt.Println("ROUTINES", runtime.NumGoroutine())
	for i := 0; i < 10; i++ {
		go func() {
			for i := 0; i < 100; i++ {
				c <- i
			}
		}()
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
	for j := 0; j < 100; j++ {
		fmt.Println(<-c)
	}
	fmt.Println("ROUTINES", runtime.NumGoroutine())
}
