package message

func LogMessage(message string) {
	println(message)
}

func WarningMessage(message string) {
	println("Warning: " + message)
}

func ErrorMessage(message string) {
	println("Error: " + message)
}
