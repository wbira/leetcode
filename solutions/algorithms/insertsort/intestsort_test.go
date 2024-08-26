package insertsort

import (
	"reflect"
	"testing"
)

func TestInsertSort(t *testing.T) {
	tests := []struct {
		array  []int
		sorted []int
	}{
		{[]int{2, 7, 11, 15}, []int{2, 7, 11, 15}},
		{[]int{11, 2, 15, 4}, []int{2, 4, 11, 15}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
	}

	for _, tt := range tests {
		got := InstertSort(tt.array)

		if !reflect.DeepEqual(got, tt.sorted) {
			t.Errorf("InsertSort(%v) = %v; want %v", tt.array, got, tt.sorted)
		}
	}
}
