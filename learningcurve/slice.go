package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	slice := make([]int, 0, 3)

	fmt.Println("Enter integers to add to the slice. Enter 'X' to quit.")

	for {
		fmt.Print("Enter an integer (or 'X' to quit): ")
		var input string
		fmt.Scanln(&input)

		// Check if user wants to quit
		if strings.ToUpper(input) == "X" {
			fmt.Println("Exiting program.")
			break
		}

		// Try to convert input to integer
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer or 'X' to quit.")
			continue
		}

		// Add the integer to the slice
		slice = append(slice, num)

		// Sort the slice
		sort.Ints(slice)

		// Print the contents of the slice in sorted order
		fmt.Printf("Sorted slice: %v\n", slice)
		fmt.Printf("Current slice length: %d, capacity: %d\n\n", len(slice), cap(slice))
	}
}
