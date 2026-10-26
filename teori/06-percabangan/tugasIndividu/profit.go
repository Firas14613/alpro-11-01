package main

import "fmt"

func main () {
	var profit1, profit2, x float64
	fmt.Scanln(&profit1)
	fmt.Scanln(&profit2)
	
	if profit1 > profit2 {
		x = profit1 - profit2
		fmt.Printf("Penururan sebesar %.0f", x)
	}else if profit2 > profit1 {
		x = profit2 - profit1
		fmt.Printf("Peningkatan sebesar %.0f", x)
	} else if profit1 == profit2 {
		fmt.Println("Tetap")
	}
}

// program profit
// kamus
//     profit1, profit2, x : real
// algoritma
//     input(profit1)
//     input(profit2)

//     if profit1 > profit2 then
//         x <- profit1 - profit2
//         print("Penurunan sebesar ", x)
//     else if profit2 > profit1 then
//         x <- profit2 - profit1
//         print("Peningkatan sebesar ", x)
//     else
//         print("Tetap")
//     endif
// endprogram
