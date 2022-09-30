package area

import (
	"fmt"
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

// https://stackoverflow.com/questions/47969385/go-float-comparison
func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
