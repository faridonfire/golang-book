package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		divisibleBy3 := i%3 == 0
		divisibleBy5 := i%5 == 0

		if divisibleBy3 && divisibleBy5 {
			fmt.Println("FIZZBUZZ")
		} else if divisibleBy3 {
			fmt.Println("FIZZ")
		} else if divisibleBy5 {
			fmt.Println("BUZZ")

		}
	}

}
