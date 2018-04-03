package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

// START OMIT
func main() {
	weird := "D\u0323\u0307"
	fmt.Println(weird)
	fmt.Printf("length in bytes: %d\n", len(weird))
	fmt.Printf("length in runes: %d\n", len([]rune(weird)))
	fmt.Printf("Number of graphemes: %d\n", countGraphemes(weird))
}

func countGraphemes(s string) int {
	var ia norm.Iter
	ia.InitString(norm.NFD, s)
	nc := 0
	for !ia.Done() {
		nc = nc + 1
		ia.Next()
	}
	return nc
}

// END OMIT
