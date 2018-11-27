package main

import (
	"fmt"
)

func main() {
	// exercise1()
	// exercise2()
	// exercise3()
	// exercise4()
	// exercise5()
	// exercise6()
	// exercise7()
	// exercise8()
	// exercise9()
	exercise10()
}

func exercise1() {
	a := [5]int{1, 3, 5, 7, 9}
	for _, v := range a {
		fmt.Print(v)
		fmt.Printf(" - %T \n", v)
	}
	fmt.Printf("%T\n", a)
}

func exercise2() {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for _, v := range a {
		fmt.Println(v)
	}
	fmt.Printf("%T \n", a)
}

func exercise3() {
	b := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(b[:5], b[5:10], b[2:7], b[1:6], b)
}

func exercise4() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println(x)
	x = append(x, 53, 54, 55)
	fmt.Println(x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println(x)
}

func exercise5() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := append(x[:3], x[6:]...)
	fmt.Println(y)
}

func exercise6() {
	states := []string{
		`Alabama`, `Alaska`, `Arizona`, `Arkansas`, `California`, `Colorado`, `Connecticut`,
		`Delaware`, `Florida`, `Georgia`, `Hawaii`, `Idaho`, `Illinois`, `Indiana`, `Iowa`,
		`Kansas`, `Kentucky`, `Louisiana`, `Maine`, `Maryland`, `Massachusetts`, `Michigan`,
		`Minnesota`, `Mississippi`, `Missouri`, `Montana`, `Nebraska`, `Nevada`, `New Hampshire`,
		`New Jersey`, `New Mexico`, `New York`, `North Carolina`, `North Dakota`, `Ohio`, `Oklahoma`,
		`Oregon`, `Pennsylvania`, `Rhode Island`, `South Carolina`, `South Dakota`, `Tennessee`,
		`Texas`, `Utah`, `Vermont`, `Virginia`, `Washington`, `West Virginia`, `Wisconsin`, `Wyoming`,
	}
	fmt.Printf("Length: %v\nCapacity: %v \n", len(states), cap(states))
	for i, v := range states {
		fmt.Printf("At index %v is the State of %v\n", i, v)
	}
}

func exercise7() {
	names := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "Moneypenny", "Helloooooo, James."},
	}
	for i, v := range names {
		fmt.Printf("At index %v:\n", i)
		for _, text := range v {
			fmt.Printf("%v\t", text)
		}
		fmt.Println("")
	}
}

func exercise8() {
	people := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	for name, likes := range people {
		fmt.Printf("To %v likes: \n", name)
		for i, like := range likes {
			fmt.Printf("\t%v -> %v\n", i+1, like)
		}
	}
}

func exercise9() {
	people := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	people["Oscar"] = []string{"Cycling", `Climbing`, "learn"}
	for name, likes := range people {
		fmt.Printf("To %v likes: \n", name)
		for i, like := range likes {
			fmt.Printf("\t%v -> %v\n", i+1, like)
		}
	}
}

func exercise10() {
	people := map[string][]string{
		"bond_james":      {`Shaken, not stirred`, `Martinis`, `Women`},
		`moneypenny_miss`: {`James Bond`, `Literature`, `Computer Science`},
		`no_dr`:           {`Being evil`, `Ice cream`, `Sunsets`},
	}
	people["Oscar"] = []string{"Cycling", `Climbing`, "learn"}
	delete(people, "moneypenny_miss")
	for name, likes := range people {
		fmt.Printf("To %v likes: \n", name)
		for i, like := range likes {
			fmt.Printf("\t%v -> %v\n", i+1, like)
		}
	}
}
