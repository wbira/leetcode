package searchstr

import (
	"reflect"
	"testing"
)

func TestSearchStr(t *testing.T) {
	tests := []struct {
		haystack string
		needle   string
		want     int
	}{
		// {"sadbutsad", "sad", 0},
		// {"leetcode", "leeto", -1},
		//{"mississippi", "issip", 4},
		{"mississippi", "issi", 1},
	}

	for _, tt := range tests {
		got := searchStr(tt.haystack, tt.needle)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("searchStr(%v, %v) = %v; want %v", tt.haystack, tt.needle, got, tt.want)
		}
	}
}
