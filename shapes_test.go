package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct {
		name     string
		shape    Shape
		expected float64
	}{
		{name: "rectangles", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{name: "circles", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{name: "triangles", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.expected {
				t.Errorf("got %g want %g", got, test.expected)
			}
		})
	}
}
