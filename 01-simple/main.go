package main

import (
	"errors"
	"log"
)

func main() {
	if err := registerUser("Andrei", -1); err != nil {
		log.Fatalf("error registering user: %s", err)
	}
}

func registerUser(name string, age int) error {
	if name == "" {
		return errors.New("name is empty")
	}
	if age <= 0 {
		return errors.New("age is <= 0")
	}
	return nil
}
