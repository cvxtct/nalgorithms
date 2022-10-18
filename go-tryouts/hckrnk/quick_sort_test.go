package hckrnk

import (
	"fmt"
	"reflect"
	"testing"
)

func TestQuickSort(t *testing.T) {
	var quickSortTest = []struct {
		arr      []int
		n        int
		expected []int
	}{
		{arr: []int{4, 5, 3, 7, 2}, n: 5, expected: []int{3, 2, 4, 5, 7}},
	}

	for _, tt := range quickSortTest {
		testname := fmt.Sprintf("%v", tt.arr)
		t.Run(testname, func(t *testing.T) {
			actual := QuickSort(tt.n, tt.arr)
			sorted := reflect.DeepEqual(actual, tt.expected)
			if !sorted {
				t.Errorf("got %v expected %v", actual, tt.expected)
			}
		})
	}
}
