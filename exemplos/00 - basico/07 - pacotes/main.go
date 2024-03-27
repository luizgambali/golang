package main

import (
	"fmt"
	"pacotes/message"
	"pacotes/util"
)

func main() {

	fmt.Printf("Area of rectangle: %f\n", util.Area(4, 3))

	message.LogMessage("This is a log message")
	message.WarningMessage("This is a warning message")
	message.ErrorMessage("This is an error message")

	fmt.Printf("Area of rectangle: %f\n", util.Area2(4, 3))

}
