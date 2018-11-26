package main

import (
	"fmt"
)

func main() {
	// tenThousandNumbers()
	// tripleUpperCase()
	// yearsAlive()
	// yearsAlive2()
	// modulus()
	// conditionals(4)
	// emptySwitch()
	favSport()
}

func tenThousandNumbers() {
	for i := 1; i <= 10000; i++ {
		fmt.Println(i)
	}
}

func tripleUpperCase() {
	for i := 65; i <= 90; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Printf("%#U \n", i)
		}
	}
}

func yearsAlive() {
	birthDate := 1989

	for birthDate <= 2018 {
		fmt.Println(birthDate)
		birthDate++
	}
}

func yearsAlive2() {
	birthDate := 1989

	for {
		if birthDate > 2018 {
			break
		}
		fmt.Println(birthDate)
		birthDate++
	}
}

func modulus() {
	for i := 10; i <= 100; i++ {
		fmt.Println(i % 4)
	}
}

func conditionals(number int) {
	if number < 4 {
		fmt.Println("Less than 4")
	} else if number > 4 {
		fmt.Println("Greater than 4")
	} else {
		fmt.Println("Equal to 4")
	}
}

func emptySwitch() {
	switch {
	case false:
		fmt.Println("this should not print")
	case true:
		fmt.Println("this should print")
	}

}

func favSport() {
	favSport := "cycling"
	switch favSport {
	case "football":
		fmt.Println("not so nice")
	case "cycling":
		fmt.Println("that's really nice")
	}
}
