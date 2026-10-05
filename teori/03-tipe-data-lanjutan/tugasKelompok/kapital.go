package main

import "fmt"

func main() {
	var x byte
	var check bool

	fmt.Scanf("%c", &x)
	check = x >= 'A' && x <= 'Z'
	fmt.Println(check)
}


/*
program kapital
kamus
	x : char
	check : boolean
algoritma
	input (x)
	check := x >= 'A' and x <= 'Z'
	output (check)
endprogram
*/