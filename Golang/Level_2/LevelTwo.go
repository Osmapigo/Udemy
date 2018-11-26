package main

import (
	"fmt"
)

func main() {
	// numericalSystems(44)
	// comparisons()
	// printConstants()
	// printShiftedInt(1019)
	// rawStringLiteral()
	yearsIota()
}

func numericalSystems(number int) {
	fmt.Printf("%v \t %b \t %#x \n", number, number, number)
}

func comparisons() {
	a := 1 == 2
	b := 1 <= 2
	c := 1 >= 2
	d := 1 != 2
	e := 1 > 2
	f := 1 < 2
	fmt.Printf("%v, %v, %v, %v, %v, %v \n", a, b, c, d, e, f)
}

func printConstants() {
	const A = 1
	const B int = 1
	fmt.Println(A)
	fmt.Println(B)
}

func printShiftedInt(number int) {
	fmt.Printf("%v - %b - %x \n", number, number, number)
	number = number << 1
	fmt.Printf("%v - %b - %x \n", number, number, number)
}

func rawStringLiteral() {
	s := `This is a literal
	with multiple "lines" inside
	:)`
	fmt.Println(s)
}

func yearsIota() {
	const (
		y1 = 2018 + iota
		y2 = 2018 + iota
		y3 = 2018 + iota
		y4 = 2018 + iota
	)
	fmt.Println(y1, y2, y3, y4)
}

//Quiz: https://docs.google.com/forms/d/e/1FAIpQLSfjhxXjo0r_OsVys58B1lVs35CLPpneVcjiEKTPsLuQs4mftA/formResponse
