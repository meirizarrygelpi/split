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

func TestEquals(t *testing.T) {}

func TestCopy(t *testing.T) {}

func TestString(t *testing.T) {}

func TestNew(t *testing.T) {}

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

func TestIsZeroDiv(t *testing.T) {}

func TestQuo(t *testing.T) {}

func ExampleNew() {
	fmt.Println(New(1, 0))
	fmt.Println(New(0, -1))
	fmt.Println(New(3, -1))
	fmt.Println(New(5, 7))
	// Output:
	// (1+0h)
	// (0-1h)
	// (3-1h)
	// (5+7h)
}
