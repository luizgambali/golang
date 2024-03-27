package main

import "fmt"

func main() {

	// Variables
	var name string = "John"
	var age int = 25

	// Shorthand
	// name := "John"
	// age := 25

	// Multiple variables
	// var name, age = "John", 25
	// name, age := "John", 25

	fmt.Printf("Name: %s, Age: %d\n", name, age)

	// Constants
	const pi = 3.14
	const (
		monday  = "Monday"
		tuesday = "Tuesday"
		fator   = 10.6
	)

	fmt.Printf("Pi: %f, Monday: %s, Tuesday: %s fator %f\n", pi, monday, tuesday, fator)

	// Arrays
	var days [7]string
	days[0] = "Sunday"
	days[1] = "Monday"
	days[2] = "Tuesday"
	days[3] = "Wednesday"
	days[4] = "Thursday"
	days[5] = "Friday"

	fmt.Printf("Days: %v\n", days)
	fmt.Printf("Day[4]]: %v\n", days[4])

	// Slices
	var months = []string{"January", "February", "March", "April", "May", "June", "July"}

	fmt.Printf("Months: %v\n", months)

	months = append(months, "August", "September", "October", "November", "December")

	fmt.Printf("Months: %v\n", months)

	// Maps
	var person = map[string]string{
		"name": "John",
		"age":  "25",
	}

	fmt.Printf("Person: %v\n", person)

}
