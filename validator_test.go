package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[T any] struct {
	name     string
	input    T
	expected bool
}

type StructTest struct {
	Name  string `validate:"required" json:"name"`
	Age   int    `validate:"optional" json:"age"`
	Phone int    `validate:"required" json:"phone"`
}

func TestValidate__isEmail(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid email address (Missing @)",
			input:    "test.com",
			expected: false,
		},
		{
			name:     "not valid email address (Empty string)",
			input:    "",
			expected: false,
		},
		{
			name:     "not valid email address (Missing .)",
			input:    "test@testcom",
			expected: false,
		},
		{
			name:     "not valid email address (Missing domain)",
			input:    "test@.com",
			expected: false,
		},
		{
			name:     "not valid email address (Missing local part)",
			input:    "@test.com",
			expected: false,
		},
		{
			name:     "not valid email address (Missing local part and domain)",
			input:    "@.com",
			expected: false,
		},
		{
			name:     "valid email address",
			input:    "test@gmail.com",
			expected: true,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsEmail(tc.input)
			if tc.expected {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, ErrEmailNotValid)
			}
			assert.Equal(t, tc.expected, isValid)
		})
	}
}

func TestValidate__isEmpty(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "not empty string",
			input:    "test",
			expected: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsEmpty(tc.input)
			if tc.expected {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, ErrNotEmptyField)
			}
			assert.Equal(t, tc.expected, isValid)
		})
	}
}

func TestValidate__isURL(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid url",
			input:    "",
			expected: false,
		},
		{
			name:     "not valid url",
			input:    "test.com",
			expected: false,
		},
		{
			name:     "valid url",
			input:    "https://www.google.com",
			expected: true,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsURL(tc.input)
			if tc.expected {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, ErrUrlNotValid)
			}
			assert.Equal(t, tc.expected, isValid)
		})
	}
}

func TestValidate__isIP(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid ip",
			input:    "192.168.0",
			expected: false,
		},
		{
			name:     "valid ip",
			input:    "192.168.0.1",
			expected: true,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsIP(tc.input)
			if tc.expected {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, ErrIpAddressNotValid)
			}
			assert.Equal(t, tc.expected, isValid)
		})
	}

}

func TestValidate__validateStruct(t *testing.T) {
	testCases := []TestCase[StructTest]{
		{
			name: "struct with optional field",
			input: StructTest{
				Name: "test",
				Phone: 1234567890,
			},
			expected: true,
		},
		{
			name: "struct with missing required name field",
			input: StructTest{
				Age: 10,
				Phone: 1234567890,
			},
			expected: false,
		},
		{
			name: "struct with missing required phone field",
			input: StructTest{
				Name: "test",
				Age: 10,
			},
			expected: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errs := validator.ValidateStruct(tc.input)

			if tc.expected {
				assert.Equal(t, 0, len(errs))
			} else {
				assert.Equal(t, 1, len(errs))
			}

		})
	}
}
