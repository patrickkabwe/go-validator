# Go Validator

The go-validator package provides a set of functions for validating common data types in a Go Application. It includes validation for email addresses, URLs, IP addresses, and structural validation of structs based on tag annotations.

![Coverage](https://github.com/Kazion500/go-validator/actions/workflows/coverage.yml/badge.svg)

## Installation

To use this package, you need to have Go installed and set up. Then you can run the following command to add the package to your project:

```go
go get github.com/kazion500/go-validator
```

## Usage

Import the validator package into your Go code:

```go
import "github.com/kazion500/go-validator"
```

## Creating a Validator

To create a new instance of the Validator interface, use the New() function:

```go
v := validator.New()
```

### Validating Email Addresses

To check if a string is a valid email address, use the IsEmail() method:

```go
ok, err := v.IsEmail(email)
if err != nil {
    // handle the error
}
```

### Checking for Empty Fields

To validate if a string is empty, use the IsEmpty() method:

```go
ok, err := v.IsEmpty(input)
if err != nil {
    // handle the error
}
```

### Validating URLs

To validate if a string is a valid URL, use the IsURL() method:

```go
ok, err := v.IsURL(input)
if err != nil {
    // handle the error
}
```

### Validating IP Addresses

To validate if a string is a valid IP address, use the IsIP() method:

```go
ok, err := v.IsIP(input)
if err != nil {
    // handle the error
}
```

### Validating Structs

To perform structural validation on a struct, use the ValidateStruct() method. It checks for fields with validate tags and returns a slice of errors:

```go
errors := v.ValidateStruct(input)
if len(errors) > 0 {
    // handle the validation errors
}
```

### Error Types

The package defines the following error types:

`ErrNotImplement`: Returned when a method is not implemented.

`ErrEmailNotValid`: Returned when an email address is not valid.

`ErrEmptyField`: Returned when a field is empty.

`ErrNotEmptyField`: Returned when a field is not empty.

`ErrUrlNotValid`: Returned when a URL is not valid.

`ErrIpAddressNotValid`: Returned when an IP address is not valid.

### Contributing

Contributions to this package are welcome. Feel free to submit issues and pull requests on the GitHub repository.

### License

This package is licensed under the MIT License. See the LICENSE file for more information.
