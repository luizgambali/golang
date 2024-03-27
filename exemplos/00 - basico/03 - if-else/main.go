package main

func main() {

	// If-Else
	if 7%2 == 0 {
		println("7 is even")
	} else {
		println("7 is odd")
	}

	// If without else
	if 8%4 == 0 {
		println("8 is divisible by 4")
	}

	// A statement can precede conditionals
	if num := 9; num < 0 {
		println(num, "is negative")
	} else if num < 10 {
		println(num, "has 1 digit")
	} else {
		println(num, "has multiple digits")
	}

}
