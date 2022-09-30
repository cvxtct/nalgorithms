package main

import "fmt"

func (app *Config) Ffunc() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g")
	app.g(0)
	fmt.Println("Returned normally from g")

}

func (app *Config) g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	app.g(i + 1)
}
