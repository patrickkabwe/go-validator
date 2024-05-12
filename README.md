![Coverage](https://github.com/Kazion500/go-validator/actions/workflows/coverage.yml/badge.svg)

# Go Validator

Go Validator is a simple package for validating structs, email addresses, URLs, IP addresses, and empty fields in Go. It also provides a way to validate struct fields using struct tags.

## Installation

To use this package, you need to have Go installed and set up. Then you can run the following command to add the package to your project:

```go
go get github.com/patrickkabwe/go-validator
```

## Features 

- 📧 `Email` validation
- 🌐 `URL` validation
- 🌐 `IP` address validation
- 📝 `Empty` field validation
- 📦 `Struct` validation using struct tags
- 📊 `Map, Slice, Embedded Struct` validation (Coming Soon)

## Usage

Import the validator package into your Go code:

```go
import "github.com/patrickkabwe/go-validator"
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

var input = "" // empty string
var input = "John Doe" // non-empty string

ok, err := v.IsEmpty(input)
if err != nil {
    // handle the error
}
```

### Validating URLs

To validate if a string is a valid URL, use the IsURL() method:

```go

var input = "https://example.com"

ok, err := v.IsURL(input)
if err != nil {
    // handle the error
}
```

### Validating IP Addresses

To validate if a string is a valid IP address, use the IsIP() method:

```go

var input = "127.0.0.1"

ok, err := v.IsIP(input)
if err != nil {
    // handle the error
}
```

### Validating Structs

To perform structural validation on a struct, use the ValidateStruct() method. It checks for fields with validate tags and returns a slice of errors:

```go

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
}

input := User{
    Name:  "John Doe",
    Email: "test@gmail.com",
}

errors := v.ValidateStruct(input)
if len(errors) > 0 {
    // handle the validation errors
}
```


### Contributing

Contributions to this package are welcome. Feel free to submit issues and pull requests on the GitHub repository.

### License

This package is licensed under the MIT License. See the LICENSE file for more information.
