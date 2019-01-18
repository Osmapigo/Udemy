package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// fmt.Println("GoArch", runtime.GOARCH)
	// fmt.Println("GoOS", runtime.GOOS)
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	// fmt.Println("GoRoutines", runtime.NumCPU())
	wg.Add(2)
	go foo()
	go bar()
	fmt.Println("GoRoutines", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for index := 0; index < 10; index++ {
		fmt.Println("This is foo")
	}
	wg.Done()
}

func bar() {
	for index := 0; index < 10; index++ {
		fmt.Println("This is bar")
	}
	wg.Done()
}
