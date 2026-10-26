package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	if n < 0 {
		n = -n
	}

	fmt.Println(n)
}

//program nilaiMutlak
//kamus
//    n : integer
//algoritma
//    input(n)
//    if n < 0 then
//        n <- -n
//    endif
//    output(n)
//endprogram