package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field string
}

func (e ValidationError) Error() string {
	return fmt.Sprintf("%s is invalid", e.Field)
}

func main() {
	err := registerUser("", 20)

	var errValidation ValidationError
	if errors.As(err, &errValidation) { // will check for type and unpack err to errValidation
		fmt.Println("err is of type ValidationError and has field:", errValidation.Field)
	}
}

func registerUser(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "name"}
	}
	if age <= 0 {
		return ValidationError{Field: "age"}
	}
	return nil
}
