// Package split implements the split-complex number arithmetic.
package split

const delta = 0.00000001

// notEquals function.
func notEquals(a, b float64) bool {
	return ((a - b) > delta) || ((b - a) > delta)
}
