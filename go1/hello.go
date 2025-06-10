package main


import (
	"fmt"
	"strings"
)

func main() {
	var input string

	fmt.Print("Enter a string: ")
	fmt.Scanln(&input)

	lowerInput := strings.ToLower(input)

	if len(lowerInput) == 0 {
		fmt.Println("Not Found!")
		return
	}

	startsWithI := lowerInput[0] == 'i'
	endsWithN := lowerInput[len(lowerInput)-1] == 'n'
	containsA := strings.Contains(lowerInput, "a")

	if startsWithI && endsWithN && containsA {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}