package main

import "fmt"

func main() {
	isRed := false
	isGreen := true

	// && || !

	if isRed || isGreen {
		fmt.Println("Is red!")
	}
}
