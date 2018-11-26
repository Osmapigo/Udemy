package main

import (
	"fmt"
)

var x2 int
var y2 string
var z2 bool

var x3 = 44
var y3 = "third String"
var z3 = false

type fausto int

var myDog fausto

var y5 int

func main() {
	excercise1()
	excercise2()
	excercise3()
	excercise4()
	excercise5()
}

func excercise1() {
	x := 42
	y := "James Bond"
	z := true
	fmt.Println(x, y, z)
	fmt.Print(x, "\n")
	fmt.Print(y, "\n")
	fmt.Print(z, "\n")
}

func excercise2() {
	x2 = 1
	y2 = "second String"
	z2 = false
	fmt.Println(x2, y2, z2)
}

func excercise3() {
	s := fmt.Sprintf("%v %v %v", x3, y3, z3)
	fmt.Println(s)
}

func excercise4() {
	fmt.Printf("%v \n", myDog)
	fmt.Printf("%T \n", myDog)
	myDog = 42
	fmt.Printf("%v \n", myDog)
}

func excercise5() {
	y5 = int(myDog)
	fmt.Println(y5)
	fmt.Printf("%T", y5)
}

// Quiz
// https://docs.google.com/forms/d/e/1FAIpQLSfyN4xMJZPoz_2bVy-BXctXfb1a64n4deYF9jj6JLnhpwA3dw/viewscore?viewscore=AE0zAgAT_RYqcVXoZPkurrZopnshJmWGanW0hbfN4P2ns1GQWmeXSLYC7DeYmm8
