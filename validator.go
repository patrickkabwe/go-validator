package validator

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"regexp"
)

var (
	// ErrNotImplement is returned when a method is not implemented.
	ErrNotImplement      = errors.New("not implement")
	ErrEmailNotValid     = errors.New("email not valid")
	ErrEmptyField        = errors.New("field is empty")
	ErrNotEmptyField     = errors.New("field not empty")
	ErrUrlNotValid       = errors.New("url not valid")
	ErrIpAddressNotValid = errors.New("ip address not valid")
)

type ValidatorError interface {
	Error() string
}

type errorMap map[string]ValidatorError

// Validator is the interface that wraps the Validate method.
type Validator interface {
	IsEmail(input string) (ok bool, err error)
	IsEmpty(input string) (ok bool, err error)
	IsURL(input string) (ok bool, err error)
	IsIP(input string) (ok bool, err error)
	ValidateStruct(input any) []error
}

// validator is an implementation of Validator.
type validator struct {
}

// New returns a new Validator.
func New() Validator {
	return &validator{}
}

// IsEmail returns an error if the string is not a valid email address.
func (v *validator) IsEmail(email string) (ok bool, err error) {

	if email == "" {
		return false, ErrEmailNotValid
	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)

	if !emailRegex.MatchString(email) {
		return false, ErrEmailNotValid
	}

	return true, nil
}

// IsEmpty returns an error if the string is not empty.
func (v *validator) IsEmpty(input string) (ok bool, err error) {
	if input != "" {
		return false, ErrNotEmptyField
	}
	return true, nil
}

// IsURL returns an error if the string is not a valid URL.
func (v *validator) IsURL(input string) (ok bool, err error) {
	parsedURL, err := url.Parse(input)
	if err != nil {
		return false, ErrUrlNotValid
	}

	if !parsedURL.IsAbs() {
		return false, ErrUrlNotValid
	}

	return true, nil
}

// IsIP returns an error if the string is not a valid IP address.
func (v *validator) IsIP(input string) (ok bool, err error) {

	parsedIP := net.ParseIP(input)
	if parsedIP == nil {
		return false, ErrIpAddressNotValid
	}

	return true, nil
}

// ValidateStruct returns an a slice of errors if the struct input values have 'validate' tags.
func (v *validator) ValidateStruct(input any) []error {
	errors := []error{}
	st := reflect.TypeOf(input)

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		name := field.Name
		tag := field.Tag.Get("validate")
		fieldType := fmt.Sprintf("%v", field.Type)
		val := reflect.ValueOf(input).Field(i)

		switch fieldType {
		case "string":
			if tag == "required" && val.String() == "" {
				errors = append(errors, fmt.Errorf("%s is required", name))
			}
		case "int":
			if tag == "required" && val.Int() == 0 {
				errors = append(errors, fmt.Errorf("%s is required", name))
			}
		}
		// TODO: handle embedded structs, slices, maps, etc.
	}

	return errors
}
