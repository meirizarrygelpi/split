package split

import "fmt"

const epsilon = 0.00000001

// Complex type represents a split-complex number over the real numbers.
type Complex [2]float64

// notEquals function.
func notEquals(a, b float64) bool {
	return ((a - b) > epsilon) || ((b - a) > epsilon)
}

// Equals method returns true if z and x are equal.
func (z *Complex) Equals(x *Complex) bool {
	for i, v := range x {
		if notEquals(v, z[i]) {
			return false
		}
	}

	return true
}

// Set method sets z equal to x.
func (z *Complex) Set(x *Complex) *Complex {
	for i, v := range x {
		z[i] = v
	}

	return z
}

// String method returns the string version of a Complex value.
func (z Complex) String() string {
	if z[1] == 0 {
		return fmt.Sprintf("%g", z[0])
	}

	if z[0] == 0 {
		return fmt.Sprintf("%gι", z[1])
	}

	if z[1] < 0 {
		return fmt.Sprintf("%g - %gι", z[0], -z[1])
	}

	return fmt.Sprintf("%g + %gι", z[0], z[1])
}

// New function returns a pointer to a Complex value made from two given real
// numbers (i.e. float64s): a + bι.
func New(a, b float64) *Complex {
	z := new(Complex)
	z[0] = a
	z[1] = b

	return z
}

// Scalar method sets z equal to a*x, and returns z.
func (z *Complex) Scalar(x *Complex, a float64) *Complex {
	for i, v := range x {
		z[i] = a * v
	}

	return z
}

// Neg method sets z equal to the negative of x, and returns z.
func (z *Complex) Neg(x *Complex) *Complex {
	return z.Scalar(x, -1)
}

// Conj method sets z equal to the conjugate of x, and returns z.
func (z *Complex) Conj(x *Complex) *Complex {
	z[0] = +x[0]
	z[1] = -x[1]

	return z
}

// Add method sets z to the sum of x and y, and returns z.
func (z *Complex) Add(x, y *Complex) *Complex {
	for i, v := range x {
		z[i] = v + y[i]
	}

	return z
}

// Sub method sets z to the difference of x and y, and returns z.
func (z *Complex) Sub(x, y *Complex) *Complex {
	for i, v := range x {
		z[i] = v - y[i]
	}

	return z
}

// Mul method sets z to the product of x and y, and returns z.
func (z *Complex) Mul(x, y *Complex) *Complex {
	p := new(Complex).Set(x)
	q := new(Complex).Set(y)
	z[0] = (p[0] * q[0]) + (p[1] * q[1])
	z[1] = (p[0] * q[1]) + (p[1] * q[0])

	return z
}

// Quad method returns the quadrance of z, which can be either positive,
// negative, or zero.
func (z *Complex) Quad() float64 {
	return (new(Complex).Mul(z, new(Complex).Conj(z)))[0]
}
