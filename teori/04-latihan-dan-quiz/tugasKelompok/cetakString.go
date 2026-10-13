package main

import "fmt"

func main () {
	var n, i int
	var str string

	fmt.Scan(&n)
	fmt.Scan(&str)

	for i = 0; i < n; i++ {
		fmt.Println(str)
	}
}


//program cetakString
//kamus
//    n, i : integer
//    str : string

//algoritma
//    input(n)
//    input(str)

//    i <- 0
//    while i < n do
//        output(str)
//        i <- i + 1
//    endwhile
//endprogram