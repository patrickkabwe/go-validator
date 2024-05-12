package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorMap(t *testing.T) {
	testCases := []struct {
		name     string
		input    errorMap
		expected map[string]string
	}{
		{
			name:     "empty error map",
			input:    errorMap{},
			expected: map[string]string{},
		},
		{
			name:     "error map with one error",
			input:    errorMap{"email": ErrEmailNotValid},
			expected: map[string]string{"email": ErrEmailNotValid.Error()},
		},
		{
			name:     "error map with multiple errors",
			input:    errorMap{"email": ErrEmailNotValid, "phone": ErrNotEmptyField},
			expected: map[string]string{"email": ErrEmailNotValid.Error(), "phone": ErrNotEmptyField.Error()},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.input.Error(), tc.expected)
		})
	}
}
