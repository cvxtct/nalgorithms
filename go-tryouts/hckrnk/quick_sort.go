package hckrnk

import "log"

func QuickSort(n int, arr []int) []int {
	pivot := arr[0]
	e := make([]int, 0)
	l := make([]int, 0)
	r := make([]int, 0)
	helpers := make([][]int, 0)
	res := make([]int, 0)

	for i, _ := range arr {
		if arr[i] < pivot {
			l = append(l, arr[i])
		}
		if arr[i] > pivot {
			r = append(r, arr[i])
		}
		if arr[i] == pivot {
			e = append(e, arr[i])
		}
	}
	// concatenate slices
	helpers = append(helpers, l, e, r)
	log.Println(helpers)
	for _, v := range helpers {
		// fmt.Println(v)
		res = append(res, v...)

	}
	// fmt.Println(res)

	return res
}
