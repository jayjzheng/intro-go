package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	const n = 500000000 // no type
	const d = 3e20 / n  // arbitrary precision

	fmt.Println(int64(d))
	fmt.Println(math.Sin(n))
}
