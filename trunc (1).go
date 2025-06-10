package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Bitte geben Sie eine Fliesskommazahl ein:")

    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("Fehler beim Einlesen der input:", err)
        return
    }

    input = strings.TrimSpace(input)
    number, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Unguelitge Eingabe, bitte geben Sie eine gueltige Fliesskommazahl ein.")
        return
    }

    truncated := int(number)
    fmt.Println("Die abgeschnittene ganze Zahl ist:", truncated)
}