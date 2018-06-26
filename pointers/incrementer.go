package pointers

// Increment Increments the value by 1
func Increment(value *int) {
	*value = *value + 1
}
