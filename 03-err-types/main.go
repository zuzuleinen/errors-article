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

	if errors.As(err, &ValidationError{}) {
		fmt.Printf("Validation failed: %s\n", err)
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
