package validator

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"reflect"
	"regexp"
	"strings"
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
	ValidateStruct(input any) errorMap
}

// validator is an implementation of Validator.
type validator struct {
	errors errorMap
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
	if input == "" {
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

func (v *validator) isInt(input any) (ok bool, err error) {
	switch input := input.(type) {
	case int:
		return true, nil
	default:
		return false, fmt.Errorf("expected int, got %T", input)
	}
}

// ValidateStruct returns an a slice of errors if the struct input values have 'validate' tags.
func (v *validator) ValidateStruct(input any) errorMap {
	st := reflect.TypeOf(input)
	v.errors = make(errorMap)
	inputKind := st.Kind()

	if inputKind == reflect.Ptr {
		st = st.Elem()
		input = reflect.ValueOf(input).Elem().Interface()
		fmt.Println(input)
	}	

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldName := field.Name
		validateTag := field.Tag.Get("validate")
		fieldValue := reflect.ValueOf(input).Field(i)
		validateTags := strings.Split(validateTag, ",")

		var actualValue any
		switch fieldValue.Kind() {
			case reflect.Interface:
				actualValue = fieldValue.Interface()
			default:
				actualValue = fieldValue.String()
		}
		
		v.handleStructValidation( validateTags, strings.ToLower(fieldName), actualValue)

		// TODO: handle embedded structs, slices, maps, etc.
	}

	return v.errors
}


func (v *validator) handleStructValidation(input []string, fieldName string, fieldValue any) {
	for _, field := range input {
		switch field {
		case "required":
			if fieldValue == "" {
				v.errors = errorMap{fieldName: ErrEmptyField}
			}
		case "int":
			if ok, err := v.isInt(fieldValue); !ok {
				v.errors = errorMap{fieldName: err}
			}
		case "email":
			if ok, err := v.IsEmail(fieldValue.(string)); !ok {
				v.errors = errorMap{fieldName: err}
			}

		case "url":
			ok, err := v.IsURL(fieldValue.(string))
			if !ok {
				v.errors = errorMap{fieldName: err}
			}
		}
	}
}

