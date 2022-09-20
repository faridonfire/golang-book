package main

import "fmt"

func main() {
  var x string = "Hello World"
  fmt.Println(x)

	var y string
  y = "Hello World"
  fmt.Println(y)

	var a string
  a = "first"
  fmt.Println(a)
  a = "second"
  fmt.Println(a)

	var b string
	b = "first "
	fmt.Println(b)
	b = b + "second"
	fmt.Println(b)

	var c string = "hello"
	var d string = "hello"
	fmt.Println(c == d)


	const e string = "Hello World"
  fmt.Println(e)

	fmt.Print("Enter a number: ")
  var input float64
  fmt.Scanf("%f", &input)

  output := input * 2

  fmt.Println(output)
}