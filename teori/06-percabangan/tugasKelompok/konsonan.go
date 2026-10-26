package main

import "fmt"

func main() {
	var c byte

	fmt.Scanf("%c", &c)

	if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' || c == 'A' || c == 'I' || c == 'U' || c == 'E' || c == 'O' {
		fmt.Println("Bukan Konsonan")
	} else {
		fmt.Println("Konsonan")
	}
}

//program konsonan
//kamus
//    c : char
//algoritma
//    input(c)
//    if c == 'a' || c == 'i' || c == 'u' || c == 'e' || c == 'o' || c == 'A' || c == 'I' || c == 'U' || c == 'E' || c == 'O' then
//        output("Bukan Konsonan")
//    else
//        output("Konsonan")
//    endif
//endprogram