// Copyright (c) 2016 Melvin Eloy Irizarry-Gelpí
// Licenced under the MIT License.

// Package split implements the split-complex number arithmetic.
package split

const delta = 0.00000001

// notEquals function.
func notEquals(a, b float64) bool {
	// Need to implement a better way to compare float64 values.
	return a != b
}
