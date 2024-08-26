package expensify

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSplitFilenames(t *testing.T) {
	tests := []struct {
		name      string
		input     string
		expected  []string
		expectErr bool
		errMsg    string
	}{
		{
			name:      "Single Filename",
			input:     "report1.pdf",
			expected:  []string{"report1.pdf"},
			expectErr: false,
		},
		{
			name:      "Multiple Filenames",
			input:     "report1.pdf,report2.pdf,report3.pdf",
			expected:  []string{"report1.pdf", "report2.pdf", "report3.pdf"},
			expectErr: false,
		},
		{
			name:      "Filenames with Spaces",
			input:     "  report1.pdf  , report2.pdf ,   report3.pdf   ",
			expected:  []string{"report1.pdf", "report2.pdf", "report3.pdf"},
			expectErr: false,
		},
		{
			name:      "Empty Input",
			input:     "",
			expectErr: true,
			errMsg:    "invalid filename: filenames cannot be empty",
		},
		{
			name:      "Trailing Comma",
			input:     "report1.pdf,report2.pdf,",
			expectErr: true,
			errMsg:    "invalid filename: filenames cannot be empty",
		},
		{
			name:      "Comma Only",
			input:     ",",
			expectErr: true,
			errMsg:    "invalid filename: filenames cannot be empty",
		},
		{
			name:      "Leading and Trailing Spaces with Newlines",
			input:     "  report1.pdf \n, \nreport2.pdf ,\n  report3.pdf  ",
			expected:  []string{"report1.pdf", "report2.pdf", "report3.pdf"},
			expectErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := SplitFilenames(tt.input)
			if tt.expectErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expected, result)
			}
		})
	}
}
