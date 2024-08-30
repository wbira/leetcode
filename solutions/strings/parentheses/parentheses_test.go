package parentheses

import (
	"reflect"
	"testing"
)

func TestValidParentheses(t *testing.T) {
	tests := []struct {
		input string
		want  bool
	}{
		{"()", true},
		{"()[]{}", true},
		{"((", false},
		{"([])", true},
	}

	for _, tt := range tests {
		got := IsValid(tt.input)

		if !reflect.DeepEqual(got, tt.want) {
			t.Errorf("IsValid(%v) = %v; want %v", tt.input, got, tt.want)
		}
	}
}
