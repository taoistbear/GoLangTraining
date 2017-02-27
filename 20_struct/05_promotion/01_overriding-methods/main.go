package main

import "fmt"

type Person struct {
	First string
	Last  string
	Age   int
}

type DoubleZero struct {
	Person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := DoubleZero{
		Person: Person{
			First: "Miss",
			Last:  "MoneyPenny",
			Age:   19,
		},
		First:         "If looks could kill",
		LicenseToKill: false,
	}

	fmt.Println(p1.First, p1.Person.First, p1.Last, p1.Age, p1.LicenseToKill)
	fmt.Println(p2.First, p2.Person.First, p2.Last, p2.Age, p2.LicenseToKill)
}
