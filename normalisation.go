package main

import "fmt"

import "golang.org/x/text/unicode/norm"

func main() {

	// START OMIT
	one := "e\u0301"
	two := "\u00e9"

	fmt.Printf("bytes %x\n", one)

	fmt.Printf("one: %q two: %q\n", one, two)

	if one == two {
		fmt.Println("equals")
	} else {
		fmt.Printf("%x not equals %x\n", one, two)
	}

	onen := norm.NFC.String(one)
	twon := norm.NFC.String(two)

	if onen == twon {
		fmt.Printf("%x equals %x\n", onen, twon)
	} else {
		fmt.Println("not equals")
	}

	// END OMIT
}
