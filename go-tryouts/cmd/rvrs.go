package main

import "fmt"

func (c *Config) reverse(a []int32) []int32 {
	ret := make([]int32, 0)
	// fmt.Println(len(a))
	// fmt.Println(a[0])
	for i := len(a) - 1; i >= 0; i-- {
		fmt.Println(a[i])
		ret = append(ret, a[i])
	}
	return ret
}
