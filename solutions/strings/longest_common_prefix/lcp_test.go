package longestcommonprefix

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		input []string
		want  string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
		{[]string{""}, ""},
		{[]string{"flower", "flo", "floight"}, "flo"},
	}

	for _, tt := range tests {
		got := LongestCommonPrefix(tt.input)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RomanToInt(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
