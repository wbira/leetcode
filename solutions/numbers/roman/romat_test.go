package roman

import (
	"reflect"
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		roman string
		want  int
	}{
		// {"III", 3},
		// {"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for _, tt := range tests {
		got := RomanToInt(tt.roman)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("RomanToInt(%v) = %v; want %v", tt.roman, got, tt.want)
		}
	}
}
