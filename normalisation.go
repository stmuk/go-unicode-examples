// https://blog.golang.org/normalization
package main

import "fmt"

import "golang.org/x/text/unicode/norm"

func main() {

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
