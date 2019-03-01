package main

import (
	"fmt"
)

type customErr struct {
	err string
}

func (ce customErr) Error() string {
	return fmt.Sprintf("This is the custom error message, as the error was %v", ce.err)
}

func main() {
	c1 := customErr{
		err: "string",
	}
	foo(c1)
}

func foo(e error) {
	fmt.Println(e)
}
