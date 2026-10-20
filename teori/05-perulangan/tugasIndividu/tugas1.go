package main

import "fmt"

func main() {
	var bulat, total int

	fmt.Scanln(&bulat)
	for bulat != 0 {
		total += bulat
		fmt.Scanln(&bulat)
	}
	fmt.Println(total)
}

// program poinGame
// kamus
// 	bulat, total : integer
// algoritma
// 	input(bulat)

// 	while bulat != 0 do
// 		total <- total + bulat
// 		input(bulat)
// 	endwhile

// 	output(total)
// endprogram