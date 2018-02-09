// https://blog.golang.org/normalization
package main

import "fmt"

import "golang.org/x/text/unicode/norm"

func main() {
	// cat crying with joy
	fmt.Println()
	s := "\U0001f639"
	fmt.Println(s)
	fmt.Printf("length in bytes %d\n", len(s))
	fmt.Printf("length in runes %d\n", len([]rune(s)))

	// weird d character
	fmt.Println()
	weird := "D\u0323\u0307"
	fmt.Println(weird)
	fmt.Printf("length in bytes %d\n", len(weird))
	fmt.Printf("length in runes %d\n", len([]rune(weird)))
	fmt.Printf("Number of chars: %d\n", countGraphemes(weird))

	// normalisation
	one := "\u00e9"
	two := "e\u0301"

	fmt.Println(one)
	fmt.Println(two)

	if one == two {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}

	if norm.NFC.String(one) == norm.NFC.String(two) {
		fmt.Println("OK")
	} else {
		fmt.Println("NOT OK")
	}

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
