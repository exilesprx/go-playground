package hworld

import "fmt"

var message = "Hello"

// PrintMessage Prints the message hello world
func PrintMessage() {
	ending := "World"

	message += " " + ending

	fmt.Println(message)
}
