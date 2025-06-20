package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Define the name struct with fixed-size string fields
type name struct {
	fname [20]byte
	lname [20]byte
}

func main() {
	// Prompt user for the filename
	fmt.Print("Enter the name of the text file: ")
	var filename string
	fmt.Scanln(&filename)
	
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return
	}
	defer file.Close()
	
	// Create a slice to store the name structs
	var names []name
	
	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	
	// Read each line from the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		
		// Skip empty lines
		if line == "" {
			continue
		}
		
		// Split the line by space to get first and last name
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			// Create a new name struct
			var n name
			
			// Copy first name to fname field (up to 20 characters)
			firstNameBytes := []byte(parts[0])
			copy(n.fname[:], firstNameBytes)
			
			// Copy last name to lname field (up to 20 characters)
			lastNameBytes := []byte(parts[1])
			copy(n.lname[:], lastNameBytes)
			
			// Add the struct to the slice
			names = append(names, n)
		}
	}
	
	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}
	
	// Print all the names from the slice of structs
	fmt.Printf("\nNames read from file:\n")
	for i, n := range names {
		// Convert byte arrays back to strings, removing null bytes
		firstName := string(n.fname[:])
		firstName = strings.TrimRight(firstName, "\x00")
		
		lastName := string(n.lname[:])
		lastName = strings.TrimRight(lastName, "\x00")
		
		fmt.Printf("%d: %s %s\n", i+1, firstName, lastName)
	}
	
	fmt.Printf("\nTotal names read: %d\n", len(names))
}