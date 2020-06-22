package main

import "fmt"

func main() {
	// membuat array string untuk input parameter fungsiA()
	b := [2]string{"Penn", "Teller"}

	//panggil fungsiA
	fmt.Printf("hasil fungsi ", fungsiA(b[:]))
}

func fungsiA(slice []string) []string {
	// create sets bernama fungsiMap
	fungsiMap := make(map[string]struct{})

	// setiap value slice digunakan untuk membuat key pada
	// sets fungsiMap yang value pada setiap key pada sets fungsiMap
	// adalah structur kosong
	for _, v := range slice {
		fungsiMap[v] = struct{}{}
	}

	// create array sepanjang sets fungsiMap dengan data kosong bukan ""
	fungsiSlice := make([]string, 0, len(fungsiMap))

	// perulangan sepanjang sets fungsiMap
	// kemudian memindah nama key pada fungsiMap menjadi value pada fungsiSlice
	for v := range fungsiMap {
		fungsiSlice = append(fungsiSlice, v)
	}

	//kembalikan nilai
	return fungsiSlice
}
