package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	tests := []struct {
		name    string
		list    []interface{}
		element interface{}
		expect  bool
	}{
		{
			name:    "Contains element in int slice",
			list:    []interface{}{1, 2, 3, 4, 5},
			element: 3,
			expect:  true,
		},
		{
			name:    "Does not contain element in int slice",
			list:    []interface{}{1, 2, 3, 4, 5},
			element: 6,
			expect:  false,
		},
		{
			name:    "Contains element in string slice",
			list:    []interface{}{"apple", "banana", "cherry"},
			element: "banana",
			expect:  true,
		},
		{
			name:    "Does not contain element in string slice",
			list:    []interface{}{"apple", "banana", "cherry"},
			element: "date",
			expect:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Contains(tt.list, tt.element)
			assert.Equal(t, tt.expect, result)
		})
	}
}
