package area

import (
	"fmt"
	"go-tryouts/area"
	"math"
	"testing"
)

func TestDistance(t *testing.T) {
	var distanceTest = []struct {
		X1, Y1, X2, Y2 float64
		expected       float64
	}{
		{X1: 0, Y1: 0, X2: 10, Y2: 10, expected: 14.142136},
		{X1: 0, Y1: 0, X2: 5, Y2: 5, expected: 7.071068},
	}

	for _, tt := range distanceTest {
		testname := fmt.Sprintf("%f, %f, %f, %f", tt.X1, tt.Y1, tt.X2, tt.Y2)
		t.Run(testname, func(t *testing.T) {
			res := distance(tt.X1, tt.Y1, tt.X2, tt.Y2)
			round := roundFloat(res, 6)
			if round != tt.expected {
				t.Errorf("got %f expected %f", round, tt.expected)
			}
		})
	}
}

func BenchmarkDistance(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distance(1, 2, 3, 4)
	}

}

func TestRectangleArea(t *testing.T) {
	var rectangleAreaTest = []struct {
		X1, Y1, X2, Y2 float64
		expected       float64
	}{
		{X1: 0, Y1: 0, X2: 0, Y2: 0, expected: 0.000000},
		{X1: 0, Y1: 0, X2: 10, Y2: 10, expected: 100.000000},
		{X1: 5, Y1: 5, X2: 10, Y2: 10, expected: 25.000000},
	}

	for _, tt := range rectangleAreaTest {
		testname := fmt.Sprintf("%f, %f, %f, %f", tt.X1, tt.Y1, tt.X2, tt.Y2)
		t.Run(testname, func(t *testing.T) {
			res := RectangleArea(tt.X1, tt.Y1, tt.X2, tt.Y2)
			round := roundFloat(res, 6)
			if round != tt.expected {
				t.Errorf("got %f expected %f", round, tt.expected)
			}
		})
	}
}

func BenchmarkRectangleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distance(1, 2, 3, 4)
	}

}

func TestCircleArea(t *testing.T) {
	var circleAreaTest = []struct {
		X, Y, R  float64
		expected float64
	}{
		{X: 0, Y: 0, R: 5, expected: 78.539816},
	}

	for _, tt := range circleAreaTest {
		testname := fmt.Sprintf("%f, %f, %f", tt.X, tt.Y, tt.R)
		t.Run(testname, func(t *testing.T) {
			res := CircleArea(tt.X, tt.Y, tt.R)
			round := roundFloat(res, 6)
			if round != tt.expected {
				t.Errorf("got %f expected %f", round, tt.expected)
			}
		})
	}
}

func TestCiArea(t *testing.T) {
	var ciAreaTest = []struct {
		X, Y, R  float64
		expected float64
	}{
		{X: 0, Y: 0, R: 5, expected: 78.539816},
	}

	for _, tt := range ciAreaTest {
		testname := fmt.Sprintf("%f, %f, %f", tt.X, tt.Y, tt.R)
		t.Run(testname, func(t *testing.T) {
			c := area.Circle{
				X: tt.X,
				Y: tt.Y,
				R: tt.R,
			}
			res := c.Area()
			round := roundFloat(res, 6)
			if round != tt.expected {
				t.Errorf("got %f expected %f", round, tt.expected)
			}
		})
	}
}

func BenchmarkCircleArea(b *testing.B) {
	for i := 0; i < b.N; i++ {
		distance(1, 2, 3, 4)
	}

}

// https://stackoverflow.com/questions/47969385/go-float-comparison
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
