package main

import (
	"fmt"
)

var message = "Hello"

func main() {

	ending := "World"

	message += " " + ending

	fmt.Println(message)

	value := 0

	increment(&value)

	fmt.Println(value)

	x := 1

	y := 2

	fmt.Printf("X: %d Y: %d", x, y)

	swap(&x, &y)

	fmt.Printf("X: %d Y: %d", x, y)
}

func increment(value *int) {
	*value = *value + 1
}

func swap(x *int, y *int) {
	tmp := *x

	*x = *y

	*y = tmp
}
