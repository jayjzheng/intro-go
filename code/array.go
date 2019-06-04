package main

import "fmt"

func main() {
	var a [5]int
	fmt.Printf("type: %T, values: %v\n", a, a)

	a[4] = 100
	fmt.Println("a[4]=", a[4])
	fmt.Println(a)

	fmt.Println("len:", len(a))
}
