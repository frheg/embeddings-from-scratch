package linalg

// Dot returns the dot product of a and b.
// It panics if a and b have different lengths.
func Dot(a, b []float32) float32 {
	if len(a) != len(b) {
		panic("linalg: Dot: length mismatch")
	}

	var total float32

	for i := range a {
		total += a[i] * b[i]
	}

	return total
}
