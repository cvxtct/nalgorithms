package main

import "fmt"

func (app *Config) FooFunc() {
	fmt.Println("Call with receiver")
}
