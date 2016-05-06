// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

package split

import (
	"fmt"
	"math"
	"strings"
)

// A Complex represents a split-complex number.
type Complex struct {
	a, b float64
}

// Real returns the real part of z, a float64 value.
func (z *Complex) Real() float64 {
	return z.a
}

// Imag returns the imaginary part of z, a float64 value.
func (z *Complex) Imag() float64 {
	return z.b
}

// SetReal sets the real part of z equal to a.
func (z *Complex) SetReal(a float64) {
	z.a = a
}

// SetImag sets the imaginary part of z equal to b.
func (z *Complex) SetImag(b float64) {
	z.b = b
}

// String returns the string version of a Complex value. If z = a + bs, then
// the string is "(a+bs)", similar to complex128 values.
func (z *Complex) String() string {
	a := make([]string, 5)
	a[0] = "("
	a[1] = fmt.Sprintf("%g", z.Real())
	switch {
	case math.Signbit(z.Imag()):
		a[2] = fmt.Sprintf("%g", z.Imag())
	case math.IsInf(z.Imag(), +1):
		a[2] = "+Inf"
	default:
		a[2] = fmt.Sprintf("+%g", z.Imag())
	}
	a[3] = "s"
	a[4] = ")"
	return strings.Join(a, "")
}

// Equals returns true if y and z are equal.
func (z *Complex) Equals(y *Complex) bool {
	if notEquals(z.Real(), y.Real()) || notEquals(z.Imag(), y.Imag()) {
		return false
	}
	return true
}

// Copy copies y onto z, and returns z.
func (z *Complex) Copy(y *Complex) *Complex {
	z.SetReal(y.Real())
	z.SetImag(y.Imag())
	return z
}

// New returns a pointer to a Complex value made from two given real float64
// values.
func New(a, b float64) *Complex {
	z := new(Complex)
	z.SetReal(a)
	z.SetImag(b)
	return z
}

// IsInf returns true if any of the components of z are infinite.
func (z *Complex) IsInf() bool {
	if math.IsInf(z.Real(), 0) || math.IsInf(z.Imag(), 0) {
		return true
	}
	return false
}

// Inf returns a pointer to a split-complex infinity value.
func Inf(a, b int) *Complex {
	z := new(Complex)
	z.SetReal(math.Inf(a))
	z.SetImag(math.Inf(b))
	return z
}

// IsNaN returns true if any component of z is NaN and neither is an infinity.
func (z *Complex) IsNaN() bool {
	if math.IsInf(z.Real(), 0) || math.IsInf(z.Imag(), 0) {
		return false
	}
	if math.IsNaN(z.Real()) || math.IsNaN(z.Imag()) {
		return true
	}
	return false
}

// NaN returns a pointer to a split-complex NaN value.
func NaN() *Complex {
	nan := math.NaN()
	z := new(Complex)
	z.SetReal(nan)
	z.SetImag(nan)
	return z
}

// Scal sets z equal to y scaled by a, and returns z.
func (z *Complex) Scal(y *Complex, a float64) *Complex {
	z.SetReal(y.Real() * a)
	z.SetImag(y.Imag() * a)
	return z
}

// Neg sets z equal to the negative of y, and returns z.
func (z *Complex) Neg(y *Complex) *Complex {
	return z.Scal(y, -1)
}

// Conj sets z equal to the conjugate of y, and returns z.
func (z *Complex) Conj(y *Complex) *Complex {
	z.SetReal(y.Real())
	z.SetImag(y.Imag() * -1)
	return z
}

// Add sets z to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	z.SetReal(x.Real() + y.Real())
	z.SetImag(x.Imag() + y.Imag())
	return z
}

// Sub sets z to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	z.SetReal(x.Real() - y.Real())
	z.SetImag(x.Imag() - y.Imag())
	return z
}

// Mul sets z to the product of x and y, and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Copy(x)
	q := new(Complex).Copy(y)
	z.SetReal((p.Real() * q.Real()) + (p.Imag() * q.Imag()))
	z.SetImag((p.Real() * q.Imag()) + (p.Imag() * q.Real()))
	return z
}

// Quad returns the quadrance of z, which can be either positive, negative, or
// zero.
func (z *Complex) Quad() float64 {
	return (new(Complex).Mul(z, new(Complex).Conj(z))).Real()
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

// Idempotent sets z equal to one of two possible idempotents (i.e. z = z*z).
func (z *Complex) Idempotent(sign int) *Complex {
	z.SetReal(0.5)
	if sign < 0 {
		z.SetImag(-0.5)
		return z
	}
	z.SetImag(0.5)
	return z
}

// Rect sets z equal to the Complex value made from given curvilinear
// coordinates and quadrance sign, and returns z.
func (z *Complex) Rect(r, ξ float64, sign int) *Complex {
	if sign > 0 {
		z.SetReal(r * math.Cosh(ξ))
		z.SetImag(r * math.Sinh(ξ))
		return z
	}
	if sign < 0 {
		z.SetReal(r * math.Sinh(ξ))
		z.SetImag(r * math.Cosh(ξ))
		return z
	}
	// sign = 0
	z.SetReal(r)
	z.SetImag(r)
	return z
}

// Curv returns the curvilinear coordinates of a Complex value, along with the
// sign of the quadrance.
func (z *Complex) Curv() (r, ξ float64, sign int) {
	quad := z.Quad()
	if quad > 0 {
		r = math.Sqrt(quad)
		ξ = math.Atanh(z.Imag() / z.Real())
		sign = +1
		return
	}
	if quad < 0 {
		r = math.Sqrt(-quad)
		ξ = math.Atanh(z.Real() / z.Imag())
		sign = -1
		return
	}
	r = z.Real()
	ξ = math.NaN()
	sign = 0
	return
}
