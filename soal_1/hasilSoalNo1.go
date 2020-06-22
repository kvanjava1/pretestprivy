package main

import (
	"fmt"
	"strconv"
)

func main() {
	var getPos int
	var strAfter string
	var str1 string = "19374048"
	var str2 = "aiueobcd"

	fmt.Printf("string pertama adalah '" + str1 + "'\n")
	fmt.Printf("string kedua adalah '" + str2 + "'\n")

	// cari posisi keberapa char 7 pada string pertama
	for pos, char := range str1 {

		// jika ketemu simpan posisi
		// langsung break selesai untuk perulangan ini
		// karena engga perlu lagi selanjutnya
		if char == '7' {
			getPos = pos
			break
		}
	}

	fmt.Printf("karakter '7' ditemukan pada posisi ke " + strconv.Itoa(getPos) + " (dimulai dari angka 0)\n")
	fmt.Printf("ambil karakter mulai dari posisi tersebut sampai karakter terakhir pada string kedua\n")

	// persiapan perulangan dari posisi yg didapat td sampai maksimal panjang string kedua
	for i := getPos; i <= len(str2)-1; i++ {
		fmt.Printf("karakter " + string(str2[i]) + " pada posisi " + strconv.Itoa(i) + "\n")
		//ambil index karakter ke i pada string kedua simpan di strAfter
		strAfter = strAfter + string(str2[i])
	}

	//tampilkan hasilnya
	fmt.Printf("karakter dari posisi " + strconv.Itoa(getPos) + " sampai posisi terakhir pada string ke 2 '" +
		str2 + "' adalah '" + strAfter + "'\n")
}
