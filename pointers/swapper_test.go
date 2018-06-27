package pointers_test

import (
	"testing"

	"github.com/exilesprx/go-playground/pointers"
)

func TestSwap(t *testing.T) {
	x := 1
	y := 2
	pointers.Swap(&x, &y)

	if x != 2 || y != 1 {
		t.Fail()
	}
}
