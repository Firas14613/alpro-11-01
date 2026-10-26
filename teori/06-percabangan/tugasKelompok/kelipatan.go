package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n%3 == 0 {
		fmt.Println("Kelipatan 3")
	}
	if n%5 == 0 {
		fmt.Println("Kelipatan 5")
	}
}

//program kelipatan
//kamus
//    n : integer
//algoritma
//    input(n)
//    if n % 3 == 0 then
//        output("Kelipatan 3")
//    endif
//    if n % 5 == 0 then
//        output("Kelipatan 5")
//    endif
//endprogram