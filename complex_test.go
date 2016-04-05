package split

import (
	"fmt"
	"testing"
)

var (
	zero = New(0, 0)
	e0   = New(1, 0)
	e1   = New(0, 1)
)

func TestQuad(t *testing.T) {
	var tests = []struct {
		x    *Complex
		want float64
	}{
		{zero, 0},
		{e0, 1},
		{e1, -1},
		{New(1, 1), 0},
	}

	for _, test := range tests {
		if got := test.x.Quad(); notEquals(got, test.want) {
			t.Errorf("Quad(%v) = %v, want %v", test.x, got, test.want)
		}
	}
}

func ExampleNew() {
	fmt.Println(New(1, 0))
	fmt.Println(New(0, 1))
	fmt.Println(New(3, -1))
	fmt.Println(New(5, 7))
	// Output:
	// 1
	// 1ι
	// 3 - 1ι
	// 5 + 7ι
}

func ExampleAdd() {
	fmt.Println(new(Complex).Add(New(1, 2), New(3, 4)))
	// Output:
	// 4 + 6ι
}

func ExampleQuad() {
	fmt.Println(New(1, 1).Quad())
	// Output:
	// 0
}
