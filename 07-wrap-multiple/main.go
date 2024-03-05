package main

import (
	"errors"
	"fmt"
)

func main() {
	err := getTwoWrappedErrors()
	if errors.As(err, &FirstError{}) {
		fmt.Println("err is wrapping FirstError")
	}
	if errors.As(err, &SecondError{}) {
		fmt.Println("err is wrapping SecondError")
	}
}

func getTwoWrappedErrors() error {
	var (
		f FirstError
		s SecondError
	)
	return fmt.Errorf("i am wrapping to errors: %w and %w", f, s)
}

type FirstError struct {
}

func (FirstError) Error() string {
	return "i'm custom error"
}

type SecondError struct {
}

func (SecondError) Error() string {
	return "i'm second error"
}
