package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "Hello World"
	printMe(printValue)

	var numerator int = 11
	var denominator int = 2
	var result, remainder, err = intDivision(numerator, denominator)

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("Result: %v", result)
	// } else {
	// 	fmt.Printf("Result: %v, Remainder: %v", result, remainder)
	// }

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("Result: %v", result)
	default:
		fmt.Printf("Result: %v, Remainder: %v", result, remainder)
	}

	switch remainder {
	case 0:
		fmt.Println("Exact division")
	case 1, 2:
		fmt.Println("Close division")
	default:
		fmt.Println("Division not close")
	}
}

func printMe(printValue string) {
	fmt.Println(printValue)
}

func intDivision(numerator int, denominator int) (int, int, error) {
	var err error = nil
	if denominator == 0 {
		err = errors.New("cannot divide by zero")
		return 0, 0, err
	}

	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
