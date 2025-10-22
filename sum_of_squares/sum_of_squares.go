package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func get_input() string {
	reader := bufio.NewReader(os.Stdin)
	input, error := reader.ReadString('\n')

	// Handle potential errors
	if error != nil {
		fmt.Println("Error reading input:", error)

		return ""
	}

	// Trim any leading/trailing whitespace or newline characters
	input = strings.TrimSpace(input)

	return input
}

func is_digit(input string) bool {
	// Check if each character in the input string is a digit
	for _, char := range input {
		if char < '0' || char > '9' {
			return false
		}
	}

	return true
}

func iterate_function(iterations int, function func() int) {

	counter := 0

	if counter < iterations {
		function()

		counter += 1

		iterate_function(iterations-counter, function)
	}
}

func get_valid_digit() int {
	input := get_input()

	if !is_digit(input) {
		return get_valid_digit()
	}

	number, _ := strconv.Atoi(input)

	return number
}

func getNumberOfIterations() int {

	iterations := get_valid_digit()

	return iterations
}

func get_number_of_values() int {
	amount := get_valid_digit()

	return amount
}

func main() {
	iteration := getNumberOfIterations()

	iterate_function(iteration, get_number_of_values)
}
