package main

import "fmt"

func (c *Config) birthdayCakeCandles(candles []int32) int32 {
	// Write your code here
	var max int32 = 0

	for e := range candles {
		fmt.Println(candles[e])
		if candles[e] > max {
			max = candles[e]
		}
	}
	fmt.Println("max", max)
	count := 0
	for i := range candles {
		if max == candles[i] {
			//fmt.Println(i)
			count++
		}
	}
	return int32(count)
}
