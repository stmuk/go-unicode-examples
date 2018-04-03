package main

import "fmt"

import "golang.org/x/text/unicode/norm"

func main() {

	// START OMIT
	one := "e\u0301"
	two := "\u00e9"

	fmt.Printf("one: %q two: %q\n", one, two)

	if one == two {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}

	if norm.NFC.String(one) == norm.NFC.String(two) {
		fmt.Println("equals")
	} else {
		fmt.Println("not equals")
	}

	// END OMIT
}
