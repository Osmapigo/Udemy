package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

//With Atomic
func main() {
	var wg sync.WaitGroup
	var counter int64
	routines := 100
	wg.Add(routines)
	for i := 0; i < routines; i++ {
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("end value:", counter)
}

// With Mutex
// func main() {
// 	var wg sync.WaitGroup
// 	var mu sync.Mutex
// 	counter := 0
// 	routines := 100
// 	wg.Add(routines)
// 	for i := 0; i < routines; i++ {
// 		go func() {
// 			mu.Lock()
// 			counter++
// 			mu.Unlock()
// 			wg.Done()
// 		}()
// 	}
// 	wg.Wait()
// 	fmt.Println("end value:", counter)
// }

// go run -race Level_9/race_condition_fixed/main.go
