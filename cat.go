package main

import "fmt"

func main() {
	fmt.Println()
	// START OMIT
	s := "\U0001f639"
	fmt.Println(s)
	fmt.Printf("length in bytes %d\n", len(s))
	fmt.Printf("length in runes %d\n", len([]rune(s)))
	// END OMIT
}
