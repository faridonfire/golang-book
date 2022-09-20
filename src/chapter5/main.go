package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}

		switch i {
			case 0: fmt.Println("Zero")
			case 1: fmt.Println("One")
			case 2: fmt.Println("Two")
			case 3: fmt.Println("Three")
			case 4: fmt.Println("Four")
			case 5: fmt.Println("Five")
			default: fmt.Println("Unknown Number")
			}
	}
}