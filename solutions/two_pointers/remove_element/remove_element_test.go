package removeelement

import (
	"reflect"
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		nums  []int
		value int
		want  int
	}{
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}

	for _, tt := range tests {
		got := removeElement(tt.nums, tt.value)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("removeDuplicates(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
