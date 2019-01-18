package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	counter := 0
	routines := 100
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			new := counter
			runtime.Gosched()
			new++
			counter = new
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)
}

// go run -race Level_9/race_condition/main.go
