package main

import "fmt"

func main() {

	// For loop
	fmt.Println("For loop")
	for i := 0; i < 5; i++ {
		fmt.Printf("\t%v\n", i)
	}

	// While loop
	fmt.Println("While loop")
	j := 0
	for j < 5 {
		fmt.Printf("\t%v\n", j)
		j++
	}

	// Infinite loop
	// for {
	// 	println("Infinite loop")
	// }

	var days = []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"}
	var count = 0

	// Controlled loop
	fmt.Println("Controlled loop")
	for count < len(days) {
		fmt.Printf("\t%v\n", days[count])
		count++
	}

	// Range loop
	fmt.Println("'For each' loop")
	for index, day := range days {
		fmt.Printf("\t%v - %v\n", index, day)
	}

	// Continue or Break loop
	fmt.Println("Continue or Break loop")

	for k := 0; k < 9; k++ {
		if k == 3 {
			continue
		}

		if k == 8 {
			break
		}

		fmt.Printf("\t%v\n", k)
	}

}
