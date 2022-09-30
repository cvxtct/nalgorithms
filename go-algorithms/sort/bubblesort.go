package sort

import (
	"go-algorithms/constraints"
	"log"
)

func Bubble[T constraints.Ordered](arr []T) []T {
	stepps := 0
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			// https://stackoverflow.com/questions/71523913/what-is-the-difference-between-comparable-and-any
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				log.Println(arr)
				stepps++
				swapped = true
			}
		}
	}
	log.Println("Stepps: ", stepps)
	return arr
}
