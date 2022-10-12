package hckrnk

import "fmt"

func InsertionSort1(n int, arr []int) []int {
	target := arr[n-1]
	idx := n - 2
	for target < arr[idx] && idx >= 0 {
		arr[idx+1] = arr[idx]
		fmt.Println(arr)
		idx--
	}
	arr[idx+1] = target
	fmt.Println(arr)
	return arr
}
