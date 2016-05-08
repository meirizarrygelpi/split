// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package split

import (
	"fmt"
	"math"
	"testing"
)

var (
	zero = &Complex{0, 0}
	one  = &Complex{1, 0}
	s    = &Complex{0, 1}
)

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

func ExampleInf() {
	fmt.Println(new(Complex).Inf(-1, -1))
	fmt.Println(new(Complex).Inf(-1, +1))
	fmt.Println(new(Complex).Inf(+1, -1))
	fmt.Println(new(Complex).Inf(+1, +1))
	// Output:
	// (-Inf-Infs)
	// (-Inf+Infs)
	// (+Inf-Infs)
	// (+Inf+Infs)
}

func ExampleNaN() {
	fmt.Println(new(Complex).NaN())
	// Output:
	// (NaN+NaNs)
}

func TestReal(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want float64
	}{
		{zero, 0},
		{one, 1},
		{s, 0},
	}
	for _, test := range tests {
		if got := test.z.Real(); notEquals(got, test.want) {
			t.Errorf("Real(%v) = %v, want %v", test.z, got, test.want)
		}
	}
}

func TestImag(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want float64
	}{
		{zero, 0},
		{one, 0},
		{s, 1},
	}
	for _, test := range tests {
		if got := test.z.Imag(); notEquals(got, test.want) {
			t.Errorf("Imag(%v) = %v, want %v", test.z, got, test.want)
		}
	}
}

func TestSetReal(t *testing.T) {
	var tests = []struct {
		a    float64
		want *Complex
	}{
		{-1, &Complex{-1, 0}},
		{+2, &Complex{+2, 0}},
	}
	for _, test := range tests {
		got := new(Complex)
		got.SetReal(test.a)
		if !got.Equals(test.want) {
			t.Errorf("SetReal(%v) = %v, want %v", test.a, got, test.want)
		}
	}
}

func TestSetImag(t *testing.T) {
	var tests = []struct {
		a    float64
		want *Complex
	}{
		{-1, &Complex{0, -1}},
		{+2, &Complex{0, +2}},
	}
	for _, test := range tests {
		got := new(Complex)
		got.SetImag(test.a)
		if !got.Equals(test.want) {
			t.Errorf("SetImag(%v) = %v, want %v", test.a, got, test.want)
		}
	}
}

func TestCartesian(t *testing.T) {
	var tests = []struct {
		z                  *Complex
		wantReal, wantImag float64
	}{
		{zero, 0, 0},
		{one, 1, 0},
		{s, 0, 1},
	}
	for _, test := range tests {
		if gotReal, gotImag := test.z.Cartesian(); notEquals(gotReal, test.wantReal) ||
			notEquals(gotImag, test.wantImag) {
			t.Errorf("Cartesian(%v) = %v, %v; want %v, %v",
				test.z, gotReal, gotImag, test.wantReal, test.wantImag)
		}
	}
}

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

func TestEquals(t *testing.T) {
	var tests = []struct {
		z, y *Complex
		want bool
	}{
		{zero, zero, true},
		{one, one, true},
		{zero, one, false},
		{&Complex{1, 2}, &Complex{3, 4}, false},
	}
	for _, test := range tests {
		if got := test.z.Equals(test.y); got != test.want {
			t.Errorf("Equals(%v, %v) = %v", test.z, test.y, got)
		}
	}
}

func TestCopy(t *testing.T) {
	var tests = []struct {
		y, want *Complex
	}{
		{zero, zero},
		{&Complex{1, 2}, &Complex{1, 2}},
	}
	for _, test := range tests {
		if got := new(Complex).Copy(test.y); !got.Equals(test.want) {
			t.Errorf("Copy(%v) = %v, want %v", test.y, got, test.want)
		}
	}
}

func TestIsInf(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want bool
	}{
		{zero, false},
		{one, false},
		{s, false},
		{&Complex{3, 4}, false},
		{&Complex{1, math.Inf(0)}, true},
	}
	for _, test := range tests {
		if got := test.z.IsInf(); got != test.want {
			t.Errorf("IsInf(%v) = %v", test.z, got)
		}
	}
}

func TestIsNaN(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want bool
	}{
		{zero, false},
		{one, false},
		{s, false},
		{&Complex{3, 4}, false},
		{&Complex{1, math.Inf(0)}, false},
		{&Complex{1, math.NaN()}, true},
		{&Complex{math.Inf(0), math.NaN()}, false},
	}
	for _, test := range tests {
		if got := test.z.IsNaN(); got != test.want {
			t.Errorf("IsNaN(%v) = %v", test.z, got)
		}
	}
}

func TestScal(t *testing.T) {
	var tests = []struct {
		y    *Complex
		a    float64
		want *Complex
	}{
		{zero, 2, zero},
		{one, 3, &Complex{3, 0}},
		{s, 4, &Complex{0, 4}},
	}
	for _, test := range tests {
		if got := new(Complex).Scal(test.y, test.a); !got.Equals(test.want) {
			t.Errorf("Scal(%v, %v) = %v, want %v",
				test.y, test.a, got, test.want)
		}
	}
}

func TestNeg(t *testing.T) {
	var tests = []struct {
		y, want *Complex
	}{
		{zero, zero},
		{one, &Complex{-1, 0}},
		{s, &Complex{0, -1}},
	}
	for _, test := range tests {
		if got := new(Complex).Neg(test.y); !got.Equals(test.want) {
			t.Errorf("Neg(%v) = %v, want %v", test.y, got, test.want)
		}
	}
}

func TestConj(t *testing.T) {
	var tests = []struct {
		y, want *Complex
	}{
		{zero, zero},
		{one, &Complex{1, 0}},
		{s, &Complex{0, -1}},
	}
	for _, test := range tests {
		if got := new(Complex).Conj(test.y); !got.Equals(test.want) {
			t.Errorf("Conj(%v) = %v, want %v", test.y, got, test.want)
		}
	}
}

func TestInvolutions(t *testing.T) {
	var tests = []struct {
		y, want *Complex
	}{
		{zero, zero},
		{one, one},
		{s, s},
		{&Complex{3, 4}, &Complex{3, 4}},
	}
	for _, test := range tests {
		w := new(Complex)
		w.Neg(test.y)
		w.Neg(w)
		if !w.Equals(test.want) {
			t.Error("Neg is not involutive")
		}
		x := new(Complex)
		x.Conj(test.y)
		x.Conj(x)
		if !x.Equals(test.want) {
			t.Error("Conj is not involutive")
		}
	}
}

func TestAdd(t *testing.T) {
	var tests = []struct {
		x, y, want *Complex
	}{
		{zero, zero, zero},
		{one, zero, one},
		{zero, one, one},
		{one, one, &Complex{2, 0}},
		{one, s, &Complex{1, 1}},
	}
	for _, test := range tests {
		if got := new(Complex).Add(test.x, test.y); !got.Equals(test.want) {
			t.Errorf("Add(%v, %v) = %v, want %v",
				test.x, test.y, got, test.want)
		}
	}
}

func TestSub(t *testing.T) {
	var tests = []struct {
		x, y, want *Complex
	}{
		{zero, zero, zero},
		{one, zero, one},
		{zero, one, &Complex{-1, 0}},
		{one, one, zero},
		{one, s, &Complex{1, -1}},
	}
	for _, test := range tests {
		if got := new(Complex).Sub(test.x, test.y); !got.Equals(test.want) {
			t.Errorf("Sub(%v, %v) = %v, want %v",
				test.x, test.y, got, test.want)
		}
	}
}

func TestMul(t *testing.T) {
	var tests = []struct {
		x, y, want *Complex
	}{
		{zero, zero, zero},
		{one, zero, zero},
		{zero, one, zero},
		{one, one, one},
		{one, s, s},
		{s, s, one},
	}
	for _, test := range tests {
		if got := new(Complex).Mul(test.x, test.y); !got.Equals(test.want) {
			t.Errorf("Mul(%v, %v) = %v, want %v",
				test.x, test.y, got, test.want)
		}
	}
}

func TestSymmetry(t *testing.T) {
	var tests = []struct {
		x, y *Complex
	}{
		{&Complex{1, 2}, &Complex{3, 4}},
		{&Complex{5, 6}, &Complex{7, 8}},
	}
	for _, test := range tests {
		if !new(Complex).Add(test.x, test.y).Equals(new(Complex).Add(test.y, test.x)) {
			t.Error("Add is not symmetric")
		}
		if !new(Complex).Mul(test.x, test.y).Equals(new(Complex).Mul(test.y, test.x)) {
			t.Error("Mul is not symmetric")
		}
	}
}

func TestQuad(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want float64
	}{
		{zero, 0},
		{one, 1},
		{s, -1},
		{&Complex{1, 1}, 0},
	}
	for _, test := range tests {
		if got := test.z.Quad(); notEquals(got, test.want) {
			t.Errorf("Quad(%v) = %v, want %v", test.z, got, test.want)
		}
	}
}

func TestIsZeroDiv(t *testing.T) {
	var tests = []struct {
		z    *Complex
		want bool
	}{
		{zero, true},
		{one, false},
		{s, false},
		{&Complex{2, 3}, false},
		{&Complex{5, 5}, true},
	}
	for _, test := range tests {
		if got := test.z.IsZeroDiv(); got != test.want {
			t.Errorf("IsZeroDiv(%v) = %v", test.z, got)
		}
	}
}

func TestInv(t *testing.T) {}

func TestQuo(t *testing.T) {}

func TestIdempotence(t *testing.T) {
	var tests = []struct {
		sign int
		want bool
	}{
		{-1, true},
		{+1, true},
	}
	for _, test := range tests {
		z := new(Complex).Idempotent(test.sign)
		if got := z.Equals(new(Complex).Mul(z, z)); got != test.want {
			t.Errorf("Idempotent(%v) is not idempotent", test.sign)
		}
	}
}

func TestRect(t *testing.T) {}

func TestCurv(t *testing.T) {}
