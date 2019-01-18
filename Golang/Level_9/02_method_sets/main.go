package main

import "fmt"

type person struct {
	Name string
}

type parrot struct {
	Name string
}

type speaker interface {
	Speak()
}

func (p person) Speak() {
	fmt.Println("Hello!, I am", p.Name)
}

func main() {
	oscar := person{
		Name: "Oscar",
	}
	saySomething(oscar)

	//The Parrot type doesn't implement Speak() method
	// lola := parrot{
	// 	Name: "Lola",
	// }
	// saySomething(lola)
}

func saySomething(s speaker) {
	s.Speak()
}
