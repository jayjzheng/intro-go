package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func div(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func main() {
	fmt.Println(add(1, 2))

	fmt.Println(div(2, 2))
	fmt.Println(div(2, 0))
}
