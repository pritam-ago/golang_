package main

import "fmt"
import "strings"

func main() {
  var stro string
  fmt.Scanf("%s", &stro)
  str := strings.ToLower(stro)
  if ( strings.HasPrefix(str, "i") && strings.HasSuffix(str, "n") && strings.Contains(str, "a") ) {
   fmt.Println("Found!")
  } else {
   fmt.Println("Not Found!")
  }
}
