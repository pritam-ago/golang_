package main

import (
    "fmt"
)

func main() {
    var num1, num2 float64

    fmt.Print("Enter first floating point number: ")
    fmt.Scan(&num1)
    fmt.Println("Truncated integer:", int(num1))

    fmt.Print("Enter second floating point number: ")
    fmt.Scan(&num2)
    fmt.Println("Truncated integer:", int(num2)) 
}
