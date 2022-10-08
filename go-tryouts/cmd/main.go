package main

import (
	"fmt"
	"go-tryouts/area"
	"go-tryouts/dpr"
)

type Config struct {
}

func main() {
	fmt.Println("Try out Defer, Panic, Recover:")
	dpr.Ffunc()
	fmt.Println("Returned normally from f.")

	fmt.Println("=============================================================")
	app := Config{}
	app.FooFunc()

	fmt.Println("=============================================================")
	fmt.Println("Area with struct - interfaces")
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println(area.RectangleArea(rx1, ry1, rx2, ry2))
	fmt.Println(area.CircleArea(cx, cy, cr))

	// var c area.Circle
	// c := new(area.Circle)
	c := area.Circle{
		X: 0,
		Y: 0,
		R: 5}
	fmt.Println(c.Area())

	// Now let's use the interface
	r := area.Rectangle{
		X1: 0,
		Y1: 0,
		X2: 10,
		Y2: 10,
	}

	fmt.Println("Total area with interface")
	res := area.TotalArea(&c, &r)
	fmt.Println(res)

	fmt.Println("================== return ========================")
	a := []int32{1, 2, 3, 4}
	fmt.Println(app.reverse(a))

}

// func (app *Config) Ffunc() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Recovered in f", r)
// 		}
// 	}()
// 	fmt.Println("Calling g")
// 	g(0)
// 	fmt.Println("Returned normally from g")

// }

// func g(i int) {
// 	if i > 3 {
// 		fmt.Println("Panicking!")
// 		panic(fmt.Sprintf("%v", i))
// 	}
// 	defer fmt.Println("Defer in g", i)
// 	fmt.Println("Printing in g", i)
// 	g(i + 1)
// }
