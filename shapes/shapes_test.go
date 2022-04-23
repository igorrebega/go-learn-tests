package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	tests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 10, Height: 10}, 40.0},
		{Circle{Radius: 10}, 62.83185307179586},
		{Triangle{Base: 2, Height: 3}, 8.0},
	}

	for _, tt := range tests {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.want)
		}
	}
}
