package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Bitte geben Sie eine Zeichenkette ein:")

    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(strings.ToLower(input))

    if strings.HasPrefix(input, "i") && strings.HasSuffix(input, "n") && strings.Contains(input, "a") {
        fmt.Println("Gefunden!")
    } else {
        fmt.Println("Nicht gefunden!")
    }
}