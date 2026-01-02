//go:build subtests

package main

import "testing"

import "github.com/stretchr/testify/assert"

func TestRectangleArea_Subtests(t *testing.T) {
	t.Run("Rectangle(3, 5)", func(t *testing.T) {
		r := Rectangle{Width: 3, Height: 5}
		area := r.Area()

		assert.Equal(t, area, 15.0)
	})

	t.Run("Rectangle(10, 20)", func(t *testing.T) {
		r := Rectangle{Width: 10, Height: 20}
		area := r.Area()

		assert.Equal(t, area, 200.0)
	})

	t.Run("Rectangle(0, 5)", func(t *testing.T) {
		r := Rectangle{Width: 0, Height: 5}
		area := r.Area()

		assert.Equal(t, area, 0.0)
	})

	t.Run("Rectangle(-2, 5)", func(t *testing.T) {
		r := Rectangle{Width: -2, Height: 5}
		area := r.Area()

		assert.Equal(t, area, 0.0)
	})
}
