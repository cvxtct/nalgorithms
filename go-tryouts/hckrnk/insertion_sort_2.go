package hckrnk

import "fmt"

func InsertionSort2(n int, arr []int) []int {
	counter := 0
	for i := 0; i <= len(arr)-1; i++ {
		if i == 0 {
			continue
		}
		counter = 0
		for j := i - 1; j >= 0; j-- {
			if arr[i-counter] < arr[j] {
				tmp := arr[j]
				arr[j] = arr[i-counter]
				arr[i-counter] = tmp
			}
			counter++
		}
		fmt.Println(arr)
	}
	return arr
}
