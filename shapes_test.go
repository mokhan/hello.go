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
		name  string
		shape Shape
		want  float64
	}{
		{"rectangles", Rectangle{12, 6}, 72.0},
		{"circles", Circle{10}, 314.1592653589793},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("got %g want %g", got, test.want)
			}
		})
	}
}
