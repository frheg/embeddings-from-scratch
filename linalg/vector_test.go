package linalg

import (
	"math"
	"testing"
)

const epsilon = 1e-6

func almostEqual(a, b float32) bool {
	return math.Abs(float64(a-b)) < epsilon
}

func TestDot(t *testing.T) {
	tests := []struct {
		name string
		a, b []float32
		want float32
	}{
		{"basic", []float32{1, 2, 3}, []float32{4, 5, 6}, 32},
		{"empty", []float32{}, []float32{}, 0},
		{"orthogonal", []float32{1, 0}, []float32{0, 1}, 0},
		{"negatives", []float32{-1, 2}, []float32{3, -4}, -11},
		{"zero vector", []float32{0, 0, 0}, []float32{7, 8, 9}, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Dot(tt.a, tt.b)
			if !almostEqual(got, tt.want) {
				t.Errorf("Dot(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}

func TestDotPanicsOnMismatch(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic, got none")
		}
	}()
	Dot([]float32{1, 2}, []float32{1})
}
