package sort

import (
	"go-algorithms/constraints"
	"log"
)

func getNextGap(gap int) int {
	gap = (gap * 10) / 13
	if gap < 1 {
		return 1
	}
	return gap
}

// Comb is impruvement of bouble sort
func Comb[T constraints.Ordered](arr []T) []T {
	n := len(arr)
	gap := n
	swapped := true
	stepps := 0
	for swapped || gap != 1 {
		gap = getNextGap(gap)
		log.Println("Gap: ", gap)
		swapped = false
		for i := 0; i < n-gap; i++ {
			if arr[i] > arr[i+gap] {
				arr[i], arr[i+gap] = arr[i+gap], arr[i]
				log.Println(arr)
				stepps++
				swapped = true
			}
		}
	}
	log.Println("Stepps: ", stepps)
	return arr
}
