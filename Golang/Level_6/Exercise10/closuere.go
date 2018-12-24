package main

import "fmt"

func main() {
	x := 33 //Outer scope
	{
		//inner scope
		x := 22
		fmt.Println(x)
	}
	fmt.Println(x)
}
