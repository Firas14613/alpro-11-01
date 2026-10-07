# <h1 align="center">Tugas Pendahuluan Modul 04 Algoritma Pemrograman</h1>
### <p align="center">"RUNTUTAN/SEKUENSI"<p>


### Soal 1 - Evaluasi Ekspresi Kontrol dalam Go

Diberikan segmen kode berikut, di mana setiap ekspresi kontrol yang terdaftar di bawah dapat digunakan sebagai kondisi pada pernyataan if:

```go
intNum := 5
intOther := 10
var sngNum float64 = -3

if _____________ {
    fmt.Println("Beep")
}
```

Untuk setiap latihan di bawah ini, tentukan apakah ekspresi kontrol tersebut menghasilkan true atau false. Tuliskan hasilnya dengan kata "true" atau "false".

1. intNum > 5  
   - False

2. intNum >= 5 && intOther < 11  
   - True

3. sngNum != -1 || intOther < 0  
   - True

4. !(intNum > 3) || intNum <= 5  
   - True

5. !(intOther >= intNum)  
   - False

6. 0 - sngNum > 0  
   - True

7. 4 / 2 == intOther / intNum  
   - True

8. intOther % 2 == 0  
   - True

9. intOther + 2 * intNum != 30 || !(sngNum > 0)  
   - True

10. intOther > 0 && intNum > 0 || sngNum > 0  
    - True

11. sngNum > 0 || (intNum >= 0 && -1 * intOther == -10)  
    - True

12. intNum == 5  
    - True

13. intNum > 0 || (sngNum <= 0 && intOther == 13)  
    - True

14. !(!(!(!(intNum > 0))))  
    - True


#
### Soal 2 Tracing : Evaluasi Pernyataan Kondisi

Berikut adalah kode dalam bahasa Go yang berisi beberapa pernyataan kondisi. Tugas Anda adalah melacak alur eksekusi program dan menentukan nilai akhir dari setiap variabel setelah program selesai dijalankan. Selain itu, tuliskan juga output yang dihasilkan oleh program.

```go
package main

import "fmt"

func main() {
	x := 10
	y := 5
	z := 15
	result := 0

	if x > 5 {
		if y < 10 {
			result = x + y
		} else {
			result = x - y
		}
	}

	if z > 10 && x == 10 {
		result += z
	} else {
		result = z - x
	}

	if x == 10 || y > 10 {
		result += 5
	} else if y == 5 && z > 10 {
		result -= 5
	} else {
		result *= 2
	}

	if !(x < 15 && y < 10) {
		result += 10
	} else {
		result -= 10
	}

	fmt.Println("Nilai akhir result:", result)
}
```

#### Pertanyaan:
1. Berapa nilai akhir dari variabel result setelah semua pernyataan kondisi dieksekusi?
2. Apa output yang dihasilkan oleh program?
3. Tuliskan langkah-langkah alur eksekusi program berdasarkan kondisi yang diberikan:
    - Kondisi 1: Apakah kondisi x > 5 benar? Jika ya, apa yang terjadi selanjutnya?
    - Kondisi 2: Apakah kondisi z > 10 && x == 10 benar? Bagaimana hal ini memengaruhi nilai result?
    - Kondisi 3: Apakah salah satu dari kondisi x == 10 || y > 10 benar? Apa yang terjadi?
    - Kondisi 4: Bagaimana kondisi !(x < 15 && y < 10) dievaluasi? Apa dampaknya pada result?

#### Jawaban:
1. 25
2. Output: Nilai akhir result: 25
3. Langkah-langkah alur eksekusi program:
    - Kondisi 1: True, result = 10 + 5 = 15
    - Kondisi 2: True, result = result(15, dari kondisi sebelumnya) + 15 = 30
    - Kondisi 3: True, result = result(30, dari kondisi sebelumnya) + 5 = 35
    - Kondisi 4: False, result = result(35, dari kondisi sebelumnya) - 10 = 25


### Soal 3  Menentukan Jumlah Hari dalam Sebulan Berdasarkan Tahun dan Bulan

Buatlah program dalam bahasa Go yang meminta input berupa tahun dan tiga huruf pertama dari nama bulan (dengan huruf pertama kapital) dari pengguna. Program kemudian menampilkan jumlah hari dalam bulan tersebut. Jika input nama bulan tidak valid, tampilkan pesan kesalahan seperti pada contoh.

