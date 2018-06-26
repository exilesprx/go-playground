package pointers

// Swap Swap the two values
func Swap(x *int, y *int) {
	tmp := *x

	*x = *y

	*y = tmp
}
