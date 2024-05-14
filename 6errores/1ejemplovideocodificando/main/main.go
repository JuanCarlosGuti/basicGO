package main

import (
	"errors"
	"fmt"
)

// representa a custom error type
type AppError struct {
	State int
}

// Error implements the error interface
func (c *AppError) Error() string {
	return fmt.Sprintf("AppError, State %d", c.State)
}

func Cause(err error) error {
	root := err
	for {
		if err = errors.Unwrap(root); err == nil {
			return root
		}
		root = err
	}
}

func main() {
	if err := firstcall(10); err != nil {
		var ap *AppError
		if errors.As(err, &ap) {
			fmt.Println("As says it is an AppError")
		}
		//Use type as context to determine cause
		switch v := Cause(err).(type) {
		case *AppError:
			fmt.Printf("Custom App Error ", v.State)
		default:
			fmt.Println("Default Error")
		}
		fmt.Println("\n**********************")
		fmt.Println("%v\n", err)
	}
}

func firstcall(i int) error {

	if err := secondCall(i); err != nil {
		return fmt.Errorf("firstCall->secondCall(%d):%w", i, err)
	}
	return nil
}
func secondCall(i int) error {

	if err := thirdCall(); err != nil {
		return fmt.Errorf("secondCall->thirdCall(): %w", err)

	}
	return nil

}

func thirdCall() error {
	return &AppError{99}

}
