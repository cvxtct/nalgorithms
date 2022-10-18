package main

import (
	"fmt"
	"go-tryouts/area"
	"go-tryouts/basics"
	"go-tryouts/dpr"
	"go-tryouts/hckrnk"
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

	fmt.Println("================== candles ========================")
	b := []int32{3, 2, 1, 3, 5, 3, 6, 2, 6, 6}
	fmt.Println(app.birthdayCakeCandles(b))

	fmt.Println("================= insertion sort 1 ===============")
	n := 5
	arr := []int{1, 2, 4, 5, 3}
	fmt.Println(arr[2])
	fmt.Println(arr[n-1])
	arr[2] = arr[3]
	fmt.Println(arr[2])
	hckrnk.InsertionSort1(n, arr)

	fmt.Println("================= insertion sort 2 ===============")
	n = 8
	arr = []int{8, 3, 6, 7, 4, 3, 2, 1}
	hckrnk.InsertionSort2(n, arr)

	fmt.Println("================= Quick sort ===============")
	n = 5
	arr = []int{4, 5, 3, 7, 2}
	result := hckrnk.QuickSort(n, arr)
	fmt.Println(result)

	fmt.Println("=============== Counting sort ==============")
	n = 100
	arr = []int{63, 25, 73, 1, 98, 73, 56, 84, 86, 57, 16, 83, 8, 25, 81, 56, 9, 53, 98, 67, 99, 12, 83, 89, 80, 91, 39, 86, 76, 85, 74, 39, 25, 90, 59, 10, 94, 32, 44, 3, 89, 30, 27, 79, 46, 96, 27, 32, 18, 21, 92, 69, 81, 40, 40, 34, 68, 78, 24, 87, 42, 69, 23, 41, 78, 22, 6, 90, 99, 89, 50, 30, 20, 1, 43, 3, 70, 95, 33, 46, 44, 9, 69, 48, 33, 60, 65, 16, 82, 67, 61, 32, 21, 79, 75, 75, 13, 87, 70, 33}
	fmt.Println(hckrnk.CountingSort(n, arr))

	fmt.Println("=========== Basics ================")
	basics.Basics()
	basics.PointerBasics()
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
