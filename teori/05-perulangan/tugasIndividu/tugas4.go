package main

import "fmt"

func main() {
	var x int
	var cantik bool

	fmt.Scan(&x)


	if x > 0 {
		prevDigit := x % 10
		x /= 10
		
		for x > 0 {
			currDigit := x % 10

			if currDigit > prevDigit {
				cantik = false
				break
			}
			prevDigit = currDigit
			x /= 10
		}
	}

	fmt.Println(cantik)
}

// program angkaCantik
// kamus
// 	x : integer
// 	cantik : boolean
// algoritma
// 	input(x)

// 	if x > 0 then
// 		prevDigit := x mod 10
// 		x := x div 10
		
// 		while x > 0 do
// 			currDigit := x mod 10

// 			if currDigit > prevDigit then
// 				cantik := false
// 				break
// 			end if
// 			prevDigit := currDigit
// 			x := x div 10
// 		end while
// 	end if

// 	output(cantik)
// endprogram