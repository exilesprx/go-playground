package main

import (
	"fmt"

	"github.com/exilesprx/go-playground/hworld"

	"github.com/exilesprx/go-playground/pointers"
)

func main() {

	hworld.PrintMessage()

	value := 0

	pointers.Increment(&value)

	fmt.Println(value)

	x := 1

	y := 2

	fmt.Printf("X: %d Y: %d", x, y)

	pointers.Swap(&x, &y)

	fmt.Printf("X: %d Y: %d", x, y)
}
