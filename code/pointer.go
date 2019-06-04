package main

import "fmt"

func zero(i int) {
	i = 0
}

func main() {
	i := 1
	fmt.Println("initial:", i)

	zero(i)
	fmt.Println("zero:", i)
}
