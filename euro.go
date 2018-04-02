package main

import "fmt"

func main() {
	fmt.Println()
	// START OMIT
	s := "\U000020A0"
	fmt.Println(s)
	fmt.Printf("length in bytes %d\n", len(s))
	fmt.Printf("length in runes %d\n", len([]rune(s)))
	// END OMIT
}
