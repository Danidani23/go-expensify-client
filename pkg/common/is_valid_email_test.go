package common

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidEmail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		valid bool
	}{
		{
			name:  "Valid email",
			email: "test@example.com",
			valid: true,
		},
		{
			name:  "Invalid email - missing @",
			email: "invalid-email",
			valid: false,
		},
		{
			name:  "Valid email with subdomain",
			email: "another.test@sub.domain.co",
			valid: true,
		},
		{
			name:  "Invalid email - missing local part",
			email: "@missinglocalpart.com",
			valid: false,
		},
		{
			name:  "Invalid email - missing domain part",
			email: "missingdomainpart@.com",
			valid: false,
		},
		{
			name:  "Invalid email - missing @ and domain",
			email: "missingatsign.com",
			valid: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.valid, IsValidEmail(tt.email))
		})
	}
}
