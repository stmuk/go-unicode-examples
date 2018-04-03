package main

import "fmt"

func main() {
	// START OMIT

	fraction := "⅜" // VULGAR FRACTION THREE EIGHTHS

	fmt.Printf("codepoint: %x\n", []rune(fraction))

	fmt.Printf("UTF-8 hex bytes: %x\n", fraction)

	// END OMIT
}
