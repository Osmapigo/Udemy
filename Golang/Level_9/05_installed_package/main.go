package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Processors:", runtime.NumCPU())
}
