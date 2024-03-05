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
	err := registerUser("andrei", -2)

	var errValidation ValidationError
	if errors.As(err, &errValidation) {
		fmt.Println("Field", errValidation.Field)
	}
}

func registerUser(name string, age int) error {
	err := validate(name, age)
	if err != nil {
		return fmt.Errorf("error on validation: %w", err)
	}
	if name == "andrei" {
		return errors.New("")
	}
	return nil
}

func validate(name string, age int) error {
	if name == "" {
		return ValidationError{Field: "name"}
	}
	if age <= 0 {
		return ValidationError{Field: "age"}
	}
	return nil
}
