package main

import (
	"fmt"
)

type ValidationError struct {
	Message string
}

func (v *ValidationError) Error() string {
	return v.Message
}

type NotFoundError struct {
	Message string
}

func (n *NotFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &ValidationError{Message: "Validation Error"}
	}
	if id != "Dimas" {
		return &NotFoundError{Message: "Data Not Found"}
	}

	return nil
}

func main() {
	err := SaveData("", nil)
	if err != nil {
		// if ValidationError, ok := err.(*ValidationError); ok {
		// 	fmt.Println(ValidationError.Message)
		// } else if NotFoundError, ok := err.(*NotFoundError); ok {
		// 	fmt.Println(NotFoundError.Message)
		// } else {
		// 	fmt.Println(err.Error())
		// }

		switch finalError := err.(type) {
		case *ValidationError:
			fmt.Println(finalError.Error())
		case *NotFoundError:
			fmt.Println(finalError.Error())
		default:
		}

	} else {
		fmt.Println("Sukses")
	}
}