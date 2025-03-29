package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strings"
)

// Calculate function with operator validity check
func calculate(a, b float64, op string) (float64, error) {
	var result float64

	// Check for valid operator
	switch op {
	case "+", "add":
		result = a + b
	case "-", "subtract":
		result = a - b
	case "*", "multiply":
		result = a * b
	case "/", "divide":
		if b == 0 {
			// Return -1 signals an error. Return 0 means completed the program, not success or failed
			return -1, fmt.Errorf("cannot divide by zero")
		}
		result = a / b
	case "^", "power":
		result = math.Pow(a, b)
	default:
		return -1, fmt.Errorf("invalid operator: %s", op)
	}
	return result, nil
}

// Outputs the inputted string and takes a pointer to num
func inputNum(prompt string, num *float64) {
	for {
		fmt.Print(prompt)
		_, err := fmt.Scanln(num)
		if err != nil {
			log.Fatal(err)
			continue
		}
		break
	}
}

func main() {
	var num1, num2 float64
	var operation string
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("Calculator!")
		inputNum("Input first number: ", &num1)
		inputNum("Input second number: ", &num2)
		fmt.Print("Input operation: ")
		_, err := fmt.Scanln(&operation)
		if err != nil {
			log.Fatal(err)
		}

		// Validate the operator and calculate the result
		result, err := calculate(num1, num2, operation)
		if err != nil {
			// If there's an error (invalid operator), display an error message and skip the calculation
			fmt.Println("Error:", err)
			continue
		}

		// If calculation was successful, print the result
		fmt.Printf("The result is %v\n", result)

		// Ask if the user wants to calculate again
		var answer string
		for {
			fmt.Print("Calculate again? (y/n, press Enter for y): ")

			// Read user input (this will handle the Enter input correctly)
			answer, err = reader.ReadString('\n')
			if err != nil {
				log.Fatal(err)
			}

			// Trim spaces and newlines from the input
			answer = strings.TrimSpace(answer)

			// If input is empty (Enter pressed), treat it as 'y'
			if answer == "" {
				answer = "y"
			}

			// Handle cases: empty input is treated as 'y'
			switch answer {
			case "y", "yes":
				// Continue the outer loop, perform another calculation
				break
			case "n", "no":
				// Exit the outer loop, ending the program
				fmt.Println("Goodbye!")
				return
			default:
				// Invalid input, prompt again
				fmt.Println("Invalid input. Please enter 'y', 'yes', 'n', or 'no'.")
				continue
			}
			// If the answer was valid (y/yes/n/no), break the loop
			break
		}

		// Exit the program if user chooses 'n' or 'no'
		if answer == "n" || answer == "no" {
			break
		}
	}
}
