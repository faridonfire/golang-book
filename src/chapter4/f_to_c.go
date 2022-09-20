package main

import "fmt"

func main() {

	fmt.Print("Enter temperature in fahrenheit: ")
  var fahrenheit float64
  fmt.Scanf("%f", &fahrenheit)

  celsius := (fahrenheit - 32.0) * 5.0 / 9.0

  fmt.Println("celsius: " + fmt.Sprintf("%f", celsius))
}