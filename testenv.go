package main

import "fmt"

func main() {

	// This will throw an error
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(result, err)
	}

	// This will not throw an error
	result, err = divide(10, 2)
	if err == nil {
		fmt.Println(result)
	}
}

func divide(x, y float32) (float32, error) {
	var result float32

	// returning error
	if y == 0 {
		return result, fmt.Errorf("Cannot divide by zero")
	}

	result = x / y
	return result, nil
}
