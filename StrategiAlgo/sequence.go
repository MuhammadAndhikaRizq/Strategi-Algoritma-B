package main

import (
	"fmt"
	"time"
)

func sequential_search(array [1000000]int, size int, nilai int) bool {
	ketemu := false
	
	for i := 0; i < size && !ketemu; {
		
		if array[i] == nilai {
			
			ketemu = true
		} else {
			
			i++
		}
	}
	return ketemu

}

func main() {

	var cari, jumlah int
	var n [1000000]int

	//Memasukan Array
	for i := 0; i < len(n)-1; i++ {
		n[i] = i
	}

	fmt.Println("================= ")
	fmt.Println("Sequential Search")
	fmt.Println("================= ")

	fmt.Println("Masukan berapa angka yang ingin diacari")
	fmt.Scanln(&jumlah)

	for i := 0; i < jumlah; i++ {
		fmt.Println("Masukan angka yang ingin dicari")
		fmt.Scanln(&cari)

		now := time.Now()
		status := sequential_search(n, len(n)-1, n[cari])
		fmt.Printf("\nNilai %d %t\n", n[cari], status)
		fmt.Println("\nEksekusi waktu ", time.Now().Sub(now))
		fmt.Println("================= ")
	}
}
