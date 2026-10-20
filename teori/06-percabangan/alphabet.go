package main

import "fmt"

func main () {
	var kar int

	fmt.Scanf("%c", &kar)
	if 'a' <= kar && kar <= 'z' || 'A' <= kar && kar <= 'Z' {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Not Alphabet")
	}
}