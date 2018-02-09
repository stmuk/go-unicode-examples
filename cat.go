package main

import "fmt"

func main() {
	// cat crying with joy
	fmt.Println()
	s := "\U0001f639"
	fmt.Println(s)
	fmt.Printf("length in bytes %d\n", len(s))
	fmt.Printf("length in runes %d\n", len([]rune(s)))
}
