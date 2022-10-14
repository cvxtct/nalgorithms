package hckrnk

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInsertionSort1(t *testing.T) {
	var insertionSort1Test = []struct {
		arr      []int
		n        int
		expected []int
	}{
		{arr: []int{1, 2, 4, 5, 3}, n: 5, expected: []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range insertionSort1Test {
		testname := fmt.Sprintf("%v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			actual := InsertionSort1(tt.n, tt.arr)
			sorted := reflect.DeepEqual(actual, tt.expected)
			if !sorted {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}

// TODO benchmark test

func TestInsertionSort2(t *testing.T) {
	var insertionSort1Test = []struct {
		arr      []int
		n        int
		expected []int
	}{
		{arr: []int{1, 2, 4, 5, 3}, n: 5, expected: []int{1, 2, 3, 4, 5}},
		{arr: []int{8, 7, 6, 5, 4, 3, 2, 1}, n: 5, expected: []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}

	for _, tt := range insertionSort1Test {
		testname := fmt.Sprintf("%v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			actual := InsertionSort2(tt.n, tt.arr)
			sorted := reflect.DeepEqual(actual, tt.expected)
			if !sorted {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}

// TODO benchmark test
