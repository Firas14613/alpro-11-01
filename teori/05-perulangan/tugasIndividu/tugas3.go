package main

import "fmt"

func main() {
	var n, m, x, y int
	var donat int

	fmt.Scan(&n, &m, &x, &y)
	
	for n >= x && m >=y{
		n -= x
		m -= y
		donat++
	}

	fmt.Println(donat)
}


// program hitungDonat
// kamus
// 	n, m, x, y : integer
// algoritma
// 	input(n, m, x, y)

// 	while n >= x and m >= y do
// 		n <- n - x
// 		m <- m - y
// 		donat <- donat + 1
// 	endwhile

// 	output(donat)
// endprogram