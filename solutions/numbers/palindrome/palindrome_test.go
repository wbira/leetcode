package palindrome

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		nums int
		want bool
	}{
		{121, true},
		{123, false},
		{-121, false},
		{10, false},
		{323, true},
	}

	for _, tt := range tests {
		got := isPalindrome(tt.nums)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IsPalindrome(%v) = %v; want %v", tt.nums, got, tt.want)
		}
	}
}
