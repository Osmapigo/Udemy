package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func (p person) speak() {
	fmt.Println("My name is", p.first, p.last, "and my age is", p.age, "years old")
}

func main() {
	p1 := person{
		"Oscar",
		"Piragauta",
		29,
	}
	p1.speak()
}
