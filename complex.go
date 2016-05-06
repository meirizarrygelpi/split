// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package split

import (
	"fmt"
	"math"
	"strings"
)

// A Complex represents a split-complex number as an ordered array of float64
// values.
type Complex [2]float64

// String returns the string version of a Complex value. If z = a + bs, then
// the string is "(a+bs)", similar to complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%g", z[0])
	switch {
	case math.Signbit(z[1]):
		a[2] = fmt.Sprintf("%g", z[1])
	case math.IsInf(z[1], +1):
		a[2] = "+Inf"
	default:
		a[2] = fmt.Sprintf("+%g", z[1])
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if z and x are equal.
func (z *Complex) Equals(x *Complex) bool {
	for i, v := range x {
		if notEquals(v, z[i]) {
			return false
		}
	}
	return true
}

// Copy copies x onto z, and returns z.
func (z *Complex) Copy(x *Complex) *Complex {
	for i, v := range x {
		z[i] = v
	}
	return z
}

// New returns a pointer to a Complex value made from two given real float64
// values.
func New(a, b float64) *Complex {
	z := new(Complex)
	z[0] = a
	z[1] = b
	return z
}

// IsInf returns true if any of the components of z are infinite.
func (z *Complex) IsInf() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return true
		}
	}
	return false
}

// Inf returns a pointer to a split-complex infinity value.
func Inf(a, b int) *Complex {
	return New(math.Inf(a), math.Inf(b))
}

// IsNaN returns true if any component of z is NaN and neither is an infinity.
func (z *Complex) IsNaN() bool {
	for _, v := range z {
		if math.IsInf(v, 0) {
			return false
		}
	}
	for _, v := range z {
		if math.IsNaN(v) {
			return true
		}
	}
	return false
}

// NaN returns a pointer to a split-complex NaN value.
func NaN() *Complex {
	nan := math.NaN()
	return New(nan, nan)
}

// Scal sets z equal to x scaled by a, and returns z.
func (z *Complex) Scal(x *Complex, a float64) *Complex {
	for i, v := range x {
		z[i] = a * v
	}
	return z
}

// Neg sets z equal to the negative of x, and returns z.
func (z *Complex) Neg(x *Complex) *Complex {
	return z.Scal(x, -1)
}

// Conj sets z equal to the conjugate of x, and returns z.
func (z *Complex) Conj(x *Complex) *Complex {
	z[0] = +x[0]
	z[1] = -x[1]
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	for i, v := range x {
		z[i] = v + y[i]
	}
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	for i, v := range x {
		z[i] = v - y[i]
	}
	return z
}

// Mul sets z to the product of x and y, and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Copy(x)
	q := new(Complex).Copy(y)
	z[0] = (p[0] * q[0]) + (p[1] * q[1])
	z[1] = (p[0] * q[1]) + (p[1] * q[0])
	return z
}

// Quad returns the quadrance of z, which can be either positive, negative, or
// zero.
func (z *Complex) Quad() float64 {
	return (new(Complex).Mul(z, new(Complex).Conj(z)))[0]
}

// IsZeroDiv returns true if z is a zero divisor (i.e. if z has vanishing
// quadrance).
func (z *Complex) IsZeroDiv() bool {
	return !notEquals(z.Quad(), 0)
}

// Inv sets z equal to the inverse of x, and returns z. If x is a zero divisor,
// then Inv panics.
func (z *Complex) Inv(x *Complex) *Complex {
	if x.IsZeroDiv() {
		panic("zero divisor has no unique inverse")
	}
	return z.Scal(new(Complex).Conj(x), 1/x.Quad())
}

// Quo sets z equal to the quotient x/y, and returns z. If y is a zero divisor,
// then Quo panics.
func (z *Complex) Quo(x, y *Complex) *Complex {
	if y.IsZeroDiv() {
		panic("denominator is a zero divisor")
	}
	return z.Scal(new(Complex).Mul(x, new(Complex).Conj(y)), 1/y.Quad())
}

// IsIndempotent returns true if z is an indempotent (i.e. if z = z*z).
func (z *Complex) IsIndempotent() bool {
	return z.Equals(new(Complex).Mul(z, z))
}

// Rect returns a Complex value made from given curvilinear coordinates and
// quadrance sign.
func Rect(r, ξ float64, sign int) *Complex {
	z := new(Complex)
	if sign > 0 {
		z[0] = r * math.Cosh(ξ)
		z[1] = r * math.Sinh(ξ)
		return z
	}
	if sign < 0 {
		z[0] = r * math.Sinh(ξ)
		z[1] = r * math.Cosh(ξ)
		return z
	}
	z[0] = r
	z[1] = r
	return z
}

// Curv returns the curvilinear coordinates of a Complex value, along with the
// sign of the quadrance.
func (z *Complex) Curv() (r, ξ float64, sign int) {
	quad := z.Quad()
	if quad > 0 {
		r = math.Sqrt(quad)
		ξ = math.Atanh(z[1] / z[0])
		sign = +1
		return
	}
	if quad < 0 {
		r = math.Sqrt(-quad)
		ξ = math.Atanh(z[0] / z[1])
		sign = -1
		return
	}
	r = z[0]
	ξ = math.NaN()
	sign = 0
	return
}
