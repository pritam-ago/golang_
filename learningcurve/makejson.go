package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	
	// Prompt for name
	fmt.Print("Enter a name: ")
	scanner.Scan()
	name := strings.TrimSpace(scanner.Text())
	
	// Prompt for address
	fmt.Print("Enter an address: ")
	scanner.Scan()
	address := strings.TrimSpace(scanner.Text())
	
	// Create a map and add the name and address
	data := make(map[string]string)
	data["name"] = name
	data["address"] = address
	
	// Use Marshal() to create a JSON object from the map
	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Printf("Error creating JSON: %v\n", err)
		return
	}
	
	// Print the JSON object
	fmt.Println("JSON object:")
	fmt.Println(string(jsonData))
}