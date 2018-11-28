package main

import (
	"fmt"
)

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	exercise4()
}

func exercise1() {
	type person struct {
		firstName   string
		lastName    string
		favIcecream []string
	}

	p1 := person{
		"Oscar",
		"Piragauta",
		[]string{
			"coconut",
		},
	}

	p2 := person{
		firstName: "Fausto",
		lastName:  "Piragauta",
		favIcecream: []string{
			"vanila",
			"bread",
		},
	}
	fmt.Println(p1.firstName, p1.lastName)
	for _, v := range p1.favIcecream {
		fmt.Println(v)
	}
	fmt.Println(p2.firstName, p2.lastName)
	for _, v := range p2.favIcecream {
		fmt.Println(v)
	}
}

func exercise2() {

	type person struct {
		firstName   string
		lastName    string
		favIcecream []string
	}

	p := person{
		firstName: "Fausto",
		lastName:  "Piragauta",
		favIcecream: []string{
			"vanila",
			"bread",
		},
	}
	a := map[string]person{
		p.lastName: p,
	}

	fmt.Printf("%T", a)
	for _, v := range a {
		fmt.Println(v.firstName, v.lastName)
		for _, iceCream := range v.favIcecream {
			fmt.Println(iceCream)
		}
	}
}

func exercise3() {
	type vehicle struct {
		doors int
		color string
	}

	type sedan struct {
		vehicle
		luxury bool
	}

	type truck struct {
		vehicle
		fourWheels bool
	}

	tesla := sedan{
		vehicle: vehicle{
			doors: 3,
			color: "red",
		},
		luxury: true,
	}

	jac := truck{
		vehicle: vehicle{
			doors: 2,
			color: "black",
		},
		fourWheels: false,
	}

	fmt.Println(tesla)
	fmt.Println(jac)
	fmt.Printf("The sedan is %v and the truck is %v\n", tesla.color, jac.color)
}

func exercise4() {
	anonymus := struct {
		size     string
		color    string
		quantity int
	}{
		size:     "L",
		color:    "Blue",
		quantity: 3,
	}
	fmt.Println(anonymus)
}
