package util

import "pacotes/message"

func Area(width, height float32) float32 {
	return width * height
}

func Area2(width, height float32) float32 {

	message.LogMessage("Area2 function called")

	return width * height
}
