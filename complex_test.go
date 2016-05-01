package split

import (
	"fmt"
	"testing"
)

var (
	zero = &Complex{0, 0}
	one  = &Complex{1, 0}
	s    = &Complex{0, 1}
)

func TestString(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want string
	}{
		{zero, "(0+0s)"},
		{one, "(1+0s)"},
		{s, "(0+1s)"},
	}
	for _, test := range tests {
		if got := test.z.String(); got != test.want {
			t.Errorf("String(%v) = %v, want %v",
				test.z, got, test.want)
		}
	}
}

func TestEquals(t *testing.T) {}

func TestCopy(t *testing.T) {}

func ExampleNew() {
	fmt.Println(New(1, 0))
	fmt.Println(New(0, -1))
	fmt.Println(New(3, -1))
	fmt.Println(New(5, 7))
	// Output:
	// (1+0s)
	// (0-1s)
	// (3-1s)
	// (5+7s)
}

func TestIsInf(t *testing.T) {}

func TestInf(t *testing.T) {}

func TestIsNaN(t *testing.T) {}

func TestNaN(t *testing.T) {}

func TestScal(t *testing.T) {}

func TestNeg(t *testing.T) {}

func TestConj(t *testing.T) {}

func TestAdd(t *testing.T) {}

func TestSub(t *testing.T) {}

func TestMul(t *testing.T) {}

func TestQuad(t *testing.T) {
	var tests = []struct {
		x    *Complex
		want float64
	}{
		{zero, 0},
		{one, 1},
		{s, -1},
		{&Complex{1, 1}, 0},
	}

	for _, test := range tests {
		if got := test.x.Quad(); notEquals(got, test.want) {
			t.Errorf("Quad(%v) = %v, want %v", test.x, got, test.want)
		}
	}
}

func TestIsZeroDiv(t *testing.T) {}

func TestInv(t *testing.T) {}

func TestQuo(t *testing.T) {}

func TestIsIndempotent(t *testing.T) {}

func ExampleRect() {}

func TestCurv(t *testing.T) {}
