package main

import (
	"fmt"
	"math/big"
)

func main() {
	// math
	fmt.Println("MATH")
  fmt.Println("1 + 1 =", 1 + 1)
	fmt.Println("1 + 1 =", 1.0 + 1.0)
	// strings
	fmt.Println("STRINGS")
	fmt.Println(len("Hello World"))
  fmt.Println("Hello World"[1])
  fmt.Println("Hello " + "World")
	// booleans
	fmt.Println("BOOLEANS")
	fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
	a := new(big.Int)
	_, err := fmt.Sscan("1982379418293479182734981723948619827346918276349817623498172634987162349817264987236984712693487123", a)
	if err != nil {

	}
	b := new(big.Int)
	_, err2 := fmt.Sscan("1289283171212983479182374918723498127349817239487128349738271178712", b)
	if err2 != nil {

	}
	fmt.Println(a.Mul(a,b))
}