
package main

import "fmt"

func main() {
    var X, total int
    fmt.Scan(&X)

    for X > 0 {
        digit := X % 10
        kuadrat := digit * digit

        fmt.Print(kuadrat, " ")

        total += kuadrat
        X /= 10
    }

    fmt.Println()
    fmt.Println(total)
}


// program kuadratDigital
// kamus
//     x, total, digit, kuadrat : integer
// algoritma
//     input(x)

//     while x > 0 do
//         digit <- x mod 10
//         kuadrat <- digit * digit
//         output(kuadrat, " ")
//         total <- total + kuadrat
//         x <- x div 10
//     endwhile

//     output(total)
// endprogram