package main

import (
	"fmt"
)

func main() {
  var x, y float64
  var output float64

  fmt.Print("Entert your first number: ")
  fmt.Scan(&x)
  fmt.Print("Entert your second number: ")
  fmt.Scan(&y)
	
  output = (x + y)
  
  fmt.Print("Result: ", output)
}
