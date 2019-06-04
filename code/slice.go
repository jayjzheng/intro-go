package main

import "fmt"

func main() {
	s := []string{}
	fmt.Printf("type: %T, values: %v\n", s, s)

	fmt.Println("len:", len(s))

	s = append(s, "a", "b", "c", "d", "e")
	fmt.Println(s)
	fmt.Println("len", len(s))

	fmt.Println(s[2:4])
	fmt.Println(s[2:])
	fmt.Println(s[:3])
}
