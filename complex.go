package split

import "fmt"

const (
	epsilon = 0.00000001
)

// Complex type represents a split-complex number over the real numbers.
type Complex struct {
	__ [2]float64
}

// E0 method returns the real part of z.
func (z *Complex) E0() float64 { return z.__[0] }

// E1 method returns the dual part of z.
func (z *Complex) E1() float64 { return z.__[1] }

// equals function.
func equals(x, y float64) bool {
	return ((x - y) < epsilon) && ((y - x) < epsilon)
}

// Equals method returns true if z and x are equal.
func (z *Complex) Equals(x *Complex) bool {
	return (equals(z.__[0], x.__[0]) &&
		equals(z.__[1], x.__[1]))
}

// Set method sets z equal to x.
func (z *Complex) Set(x *Complex) *Complex {
	for i, v := range x.__ {
		z.__[i] = v
	}

	return z
}

// String method returns the string version of a Complex value.
func (z Complex) String() string {
	if z.__[1] == 0 {
		return fmt.Sprintf("%g", z.__[0])
	}

	if z.__[0] == 0 {
		return fmt.Sprintf("%gι", z.__[1])
	}

	if z.__[1] < 0 {
		return fmt.Sprintf("%g - %gι", z.__[0], -z.__[1])
	}

	return fmt.Sprintf("%g + %gι", z.__[0], z.__[1])
}

// New function returns a pointer to a Complex value made from two given real
// numbers (i.e. float64s).
func New(a, b float64) *Complex {
	z := new(Complex)
	z.__[0] = a
	z.__[1] = b

	return z
}

// Scalar method sets z equal to s*x, and returns z.
func (z *Complex) Scalar(x *Complex, s float64) *Complex {
	for i, v := range x.__ {
		z.__[i] = s * v
	}

	return z
}

// Neg method sets z equal to the negative of x, and returns z.
func (z *Complex) Neg(x *Complex) *Complex {
	return z.Scalar(x, -1)
}

// Conj method sets z equal to the conjugate of x, and returns z.
func (z *Complex) Conj(x *Complex) *Complex {
	z.__[0] = x.__[0]
	z.__[1] = -x.__[1]

	return z
}

// Add method sets z to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	for i, v := range x.__ {
		z.__[i] = v + y.__[i]
	}

	return z
}

// Sub method sets z to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	for i, v := range x.__ {
		z.__[i] = v - y.__[i]
	}

	return z
}

// Mul method sets z to the product of x and y, and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Set(x)
	q := new(Complex).Set(y)
	z.__[0] = (p.__[0] * q.__[0]) + (p.__[1] * q.__[1])
	z.__[1] = (p.__[0] * q.__[1]) + (p.__[1] * q.__[0])

	return z
}

// Quad method returns the quadrance of z, which can be either positive,
// negative, or zero.
func (z *Complex) Quad() float64 {
	return (new(Complex).Mul(z, new(Complex).Conj(z))).__[0]
}
