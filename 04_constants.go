package main

import (
	"fmt" 
	"math"
)

const S string = "CONSTANT_STRING"

func main() {
	fmt.Println(S)

	const n = 500000

	const d = 3e20 / n

	fmt.Println(d)

	fmt.Println(int64(d))

	fmt.Println(math.Sin(n))
}