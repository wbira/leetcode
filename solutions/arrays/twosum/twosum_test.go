package twosum

import (
	"reflect"
	"sort"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{[]int{2, 7, 11, 15}, 9, []int{0, 1}},
		{[]int{3, 2, 4}, 6, []int{1, 2}},
		{[]int{3, 3}, 6, []int{0, 1}},
		{[]int{3, 2, 3}, 6, []int{0, 2}}, // Additional test case
	}

	for _, tt := range tests {
		got := TwoSum(tt.nums, tt.target)

		// Sort both slices to ensure order doesn't matter
		sort.Ints(got)
		sort.Ints(tt.want)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("TwoSum(%v, %v) = %v; want %v", tt.nums, tt.target, got, tt.want)
		}
	}
}
