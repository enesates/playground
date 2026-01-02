//go:build unittests

package main

import "testing"

func TestRectangleArea(t *testing.T) {
	tests := []struct {
		width, height float64
		result        float64
	}{
		{10, 20, 200},
		{3, 5, 15},
		{0, 5, 0},
		{-2, 5, 0},
	}

	for _, test := range tests {
		r := Rectangle{Width: test.width, Height: test.height}

		if got := r.Area(); got != test.result {
			t.Errorf("Area(%.2f, %2.f) = %.2f, expected: %.2f", test.width, test.height, got, test.result)
		}
	}
}
