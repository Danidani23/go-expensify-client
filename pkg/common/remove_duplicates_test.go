package common

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		input    []string
		expected []string
	}{
		{input: []string{"apple", "banana", "apple", "cherry", "banana"}, expected: []string{"apple", "banana", "cherry"}},
		{input: []string{"dog", "cat", "dog", "bird"}, expected: []string{"dog", "cat", "bird"}},
		{input: []string{"one", "one", "one", "one"}, expected: []string{"one"}},
		{input: []string{"a", "b", "c"}, expected: []string{"a", "b", "c"}},
		{input: []string{}, expected: []string{}},
	}

	for _, test := range tests {
		output := RemoveDuplicates(test.input)
		if !reflect.DeepEqual(output, test.expected) {
			t.Errorf("RemoveDuplicates(%v) = %v; expected %v", test.input, output, test.expected)
		}
	}
}
