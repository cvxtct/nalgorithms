package area

import "math"

type Circle struct {
	X float64
	Y float64
	R float64
}

type Rectangle struct {
	X1, Y1, X2, Y2 float64
}

type Shape interface {
	Area() float64
}

// internal use, just a helper
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

// without interface or struct receiver
func RectangleArea(x1, y1, x2, y2 float64) float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
}

// witouth interface or struc receiver
func CircleArea(x, y, r float64) float64 {
	return math.Pi * r * r
}

// struct Circle receiver
func (c *Circle) Area() float64 {
	return math.Pi * c.R * c.R
}

// struct Rectangle receiver
func (r *Rectangle) Area() float64 {
	l := distance(r.X1, r.Y1, r.X1, r.Y2)
	w := distance(r.X1, r.Y1, r.X2, r.Y1)
	return l * w
}

// Notice this one using interface, one function Area() takes care
// Since the above ones implement the Area() method
// Benefit of interface -> one call, based on the struct, it will
// execut the corresponding method
func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}
