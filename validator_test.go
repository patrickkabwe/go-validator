package validator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase[T any] struct {
	name     string
	input    T
	expectedErr bool
}

type StructTest struct {
	Name  string `validate:"required" json:"name"`
	Age   any    `validate:"int,required" json:"age"`
	Phone int    `validate:"required" json:"phone"`
	Email string `validate:"email" json:"email"`
}

func TestValidate__isEmail(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid email address (Missing @)",
			input:    "test.com",
			expectedErr: true,
		},
		{
			name:     "not valid email address (Empty string)",
			input:    "",
			expectedErr: true,
		},
		{
			name:     "not valid email address (Missing .)",
			input:    "test@testcom",
			expectedErr: true,
		},
		{
			name:     "not valid email address (Missing domain)",
			input:    "test@.com",
			expectedErr: true,
		},
		{
			name:     "not valid email address (Missing local part)",
			input:    "@test.com",
			expectedErr: true,
		},
		{
			name:     "not valid email address (Missing local part and domain)",
			input:    "@.com",
			expectedErr: true,
		},
		{
			name:     "valid email address",
			input:    "test@gmail.com",
			expectedErr: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsEmail(tc.input)
			if tc.expectedErr {
				assert.ErrorIs(t, err, ErrEmailNotValid)
			} else {
				assert.NoError(t, err)
				assert.True(t, isValid)
			}
			
		})
	}
}

func TestValidate__isEmpty(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "empty string",
			input:    "",
			expectedErr: true,
		},
		{
			name:     "not empty string",
			input:    "test",
			expectedErr: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsEmpty(tc.input)
			if tc.expectedErr {
				assert.ErrorIs(t, err, ErrNotEmptyField)
			} else {
				assert.NoError(t, err)
				assert.True(t, isValid)
			
			}
		})
	}
}

func TestValidate__isURL(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid url",
			input:    "",
			expectedErr: true,
		},
		{
			name:     "not valid url",
			input:    "test.com",
			expectedErr: true,
		},
		{
			name:     "valid url using http protocol",
			input:    "https://www.google.com",
			expectedErr: false,
		},
		{
			name:     "valid url using postgres protocol",
			input:    "postgres://localhost:5432/testdb",
			expectedErr: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsURL(tc.input)
			if tc.expectedErr {
				assert.ErrorIs(t, err, ErrUrlNotValid)
			} else {
				assert.NoError(t, err)
				assert.True(t, isValid)
			}
		})
	}
}

func TestValidate__isIP(t *testing.T) {
	testCases := []TestCase[string]{
		{
			name:     "not valid ip",
			input:    "192.168.0",
			expectedErr: true,
		},
		{
			name:     "valid ip",
			input:    "192.168.0.1",
			expectedErr: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			isValid, err := validator.IsIP(tc.input)
			if tc.expectedErr {
				assert.ErrorIs(t, err, ErrIpAddressNotValid)
			} else {
				assert.NoError(t, err)
				assert.True(t, isValid)
			}
		})
	}
}

func TestValidate__validateStruct(t *testing.T) {
	testCases := []TestCase[StructTest]{
		{
			name: "invalid struct - age is missing",
			input: StructTest{
				Name: "test",
				Phone: 1234567890,
			},
			expectedErr: true,
		},
		{
			name: "invalid struct - age is not an integer and name is missing",
			input: StructTest{
				Age: "10",
				Phone: 1234567890,
			},
			expectedErr: true,
		},
		{
			name: "invalid struct - name is missing",
			input: StructTest{
				Age: 10,
				Phone: 1234567890,
			},
			expectedErr: true,
		},
		{
			name: "invalid struct - email is not valid",
			input: StructTest{
				Name: "test",
				Age: 10,
				Phone: 1234567890,
				Email: "test",
			},
			expectedErr: true,
		},
		{
			name: "valid struct",
			input: StructTest{
				Name: "test",
				Age: 10,
				Phone: 1234567890,
				Email: "test@gmail.com",
			},
			expectedErr: false,
		},
	}

	validator := New()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			errs := validator.ValidateStruct(tc.input)

			if tc.expectedErr {
				assert.Greater(t, len(errs), 0)
			} else {
				assert.Equal(t, 0, len(errs))
			}

		})
	}
}
