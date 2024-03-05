package main

import (
	"errors"
	"fmt"
)

var (
	ErrEmptyName  = errors.New("name is empty")
	ErrInvalidAge = errors.New("age is <= 0")
)

func main() {
	err := registerUser("", 20)
	if errors.Is(err, ErrEmptyName) {
		fmt.Println("Your name is empty")
	}
	if errors.Is(err, ErrInvalidAge) {
		fmt.Println("Your age is invalid")
	}
}

func registerUser(name string, age int) error {
	if name == "" {
		return ErrEmptyName
	}
	if age <= 0 {
		return ErrInvalidAge
	}
	return nil
}
