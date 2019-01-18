package main

import (
	"fmt"
	"sort"
)

type User struct {
	First   string
	Last    string
	Age     int
	Sayings []string
}

// ByAge implements sort.Interface for []Person based on the Age field.
type ByAge []User

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }

// ByLast implements sort.Interface for []Person based on the Last name field.
type ByLast []User

func (bl ByLast) Len() int           { return len(bl) }
func (bl ByLast) Swap(i, j int)      { bl[i], bl[j] = bl[j], bl[i] }
func (bl ByLast) Less(i, j int) bool { return bl[i].Last < bl[j].Last }

// BySaying implements sort.Interface for []Person based on the Sayin field.
type BySaying []string

func (bs BySaying) Len() int           { return len(bs) }
func (bs BySaying) Swap(i, j int)      { bs[i], bs[j] = bs[j], bs[i] }
func (bs BySaying) Less(i, j int) bool { return bs[i] < bs[j] }

func main() {
	u1 := User{
		First: "James",
		Last:  "Bond",
		Age:   32,
		Sayings: []string{
			"Shaken, not stirred",
			"Youth is no guarantee of innovation",
			"In his majesty's royal service",
		},
	}
	u2 := User{
		First: "Miss",
		Last:  "Moneypenny",
		Age:   27,
		Sayings: []string{
			"James, it is soo good to see you",
			"Would you like me to take care of that for you, James?",
			"I would really prefer to be a secret agent myself.",
		},
	}

	u3 := User{
		First: "M",
		Last:  "Hmmmm",
		Age:   54,
		Sayings: []string{
			"Oh, James. You didn't.",
			"Dear God, what has James done now?",
			"Can someone please tell me where James Bond is?",
		},
	}

	users := []User{u1, u2, u3}

	fmt.Println("---Original----")
	fmt.Println(users)
	//sort sayings
	for _, user := range users {
		sort.Sort(BySaying(user.Sayings))
	}
	fmt.Println("---------------")
	//Sort by Age
	sort.Sort(ByAge(users))
	printUsers(users)
	//Sort by Last
	sort.Sort(ByLast(users))
	printUsers(users)
}

func printUsers(users []User) {
	fmt.Println("-------")
	for _, user := range users {
		fmt.Println("Name:", user.First, user.Last)
		fmt.Println("Age:", user.Age)
		fmt.Println("Sayings:")
		for _, saying := range user.Sayings {
			fmt.Println("\t", saying)
		}
	}

}
