package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

// START OMIT
func main() {

	fmt.Println("accent on its own \u0301")

	accent := "e\u0301" // acute accent e ́
	fmt.Println(accent)
	fmt.Printf("%+q\n", accent)

	fmt.Printf("length in bytes: %d\n", len(accent))

	for i := 0; i < len(accent); i++ {
		fmt.Printf("%x ", accent[i])
	}

	fmt.Printf("\nlength in runes: %d\n", len([]rune(accent)))
	fmt.Printf("Number of chars: %d\n", countGraphemes(accent))

	nor := norm.NFC.String(accent)
	fmt.Printf("%+q\n", nor)
	fmt.Printf("Number of chars: %d\n", countGraphemes(nor))

}

// END OMIT

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
