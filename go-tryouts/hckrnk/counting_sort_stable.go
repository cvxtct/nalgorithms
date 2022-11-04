package hckrnk

import "fmt"

func CountingSortStable(arr []int) []int {
	arr_len := len(arr)
	var min int
	var max int

	for i, e := range arr {
		if i == 0 || e > max {
			max = e
		}
	}

	for i, e := range arr {
		if i == 0 || e < min {
			min = e
		}
	}

	countArrLen := max + 1 - min

	count := make([]int, countArrLen)

	for _, e := range arr {
		count[e-min]++
	}
	counter := 1
	for countArrLen > counter {
		count[counter] = count[counter] + count[counter-1]
		counter++
	}

	oredered := make([]int, arr_len)
	fmt.Println("count:", len(count))
	fmt.Println("ordered:", len(oredered))
	fmt.Println("arr: ", arr_len)
	for arr_len > 0 {

		arr_len--
		oredered[count[arr[arr_len]-min]-1] = arr[arr_len]
		count[arr[arr_len]-min]--

	}

	// fmt.Println(count)
	fmt.Println(oredered)

	return oredered
}
