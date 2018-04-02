package main

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

//

// START OMIT
func main() {
	//weird := "🍻\U0001f3fd" // beer
	//weird := "e\u0301" // acute accent e
	// U0001f3fd is EMOJI MODIFIER FITZPATRICK TYPE-4
	//weird := "\U0001f575\U0001f3fd" // SPY with dark skin
	weird := "\U0001F608" // 😈
	fmt.Println(weird)
	fmt.Printf("%+q\n", weird)

	fmt.Printf("length in bytes: %d\n", len(weird))

	for i := 0; i < len(weird); i++ {
		fmt.Printf("%x ", weird[i])
	}

	fmt.Printf("\nlength in runes: %d\n", len([]rune(weird)))
	fmt.Printf("Number of chars: %d\n", countGraphemes(weird))

	not := norm.NFC.String(weird)
	fmt.Printf("%+q\n", not)
	fmt.Printf("Number of chars: %d\n", countGraphemes(not))

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
