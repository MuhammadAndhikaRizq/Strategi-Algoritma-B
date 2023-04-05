package main

import (
	"fmt"
	"time"
)

// func sequential_search(array [1000000]int, size int, nilai int) bool {
// 	/*
// 		Kita inisialisasikan variable ketemu bertipe boolean dengan nilai false
// 	*/
// 	ketemu := false
// 	/*
// 		lalu kita looping semua data yang ada pada array data, sesuai jumlah arraynya.
// 		Dengan kondisi jika i kurang dari size dan variable ketemu bernilai false
// 	*/
// 	for i := 0; i < size && !ketemu; {
// 		/*
// 			Kita cek nilai yang ada pada array tersebut apakah sama dengan nilai yang dicari
// 		*/
// 		if array[i] == nilai {
// 			/*
// 				Jika nilainya yang ada pada array tersebut sama dengan nilai yang dicari,
// 				maka kita akan menset variable ketemu menjadi true
// 			*/
// 			ketemu = true
// 		} else {
// 			/*
// 				Jika nilainya yang ada pada array tersebut tidak sama dengan nilai yang dicari,
// 				maka akan menambahkan nilai +1
// 			*/
// 			i++
// 		}
// 	}
// 	return ketemu

// }

// func print_array(array [6]int, size int) {
// 	fmt.Printf("[")

// 	for i := 0; i < size; i++ {
// 		fmt.Printf("%d", array[i])

// 		if i != size-1 {
// 			fmt.Printf(", ")
// 		}
// 	}
// 	fmt.Printf("]\n")

// }

func main() {
	now := time.Now()
	//nilai := [4]int{10, 100, 1000, 10000}
	// var data [1000000]int
	k := 1
	n := 1000000
	ketemu := false
	// defer func() {
	// 	fmt.Printf("Dari jumlah %d data\n\n", len(data))
	// }()

	fmt.Println("================= ")
	fmt.Println("Sequential Search")
	fmt.Println("================= ")

	for k <= n && !ketemu {
		if n-1 == k {
			ketemu = true
		} else {
			k++
		}
	}

	if ketemu {
		fmt.Println("Ketemu", k)
	} else {
		fmt.Println("Tidak Ketemu")
	}
	fmt.Println("\nEksekusi waktu ", time.Now().Sub(now))

	//fmt.Println("Masukkan angka yang dicari: ")
	// kita beri variable cari sebagai penampung inputan
	//var cari int
	// lalu kita akan memproses inputan
	//fmt.Scanln(&cari)
	// panggil fungsinya
	// fmt.Printf("Isi array: \n")
	// print_array(data, len(data)) // tampilkan data jika mau

	// for i := 0; i < len(nilai); i++ {
	// 	status := sequential_search(data, len(data), nilai[i])
	// 	fmt.Printf("\nNilai %d %t\n", nilai[i], status)
	// 	fmt.Println("\nEksekusi waktu ", time.Now().Sub(now))
	// }

	// //test
	// n := 10
	// var count int // inisialisasi variabel untuk menghitung jumlah operasi dasar
	// for i := 0; i < n; i++ {
	// 	for j := 0; j < n; j++ {
	// 		count++
	// 	}
	// 	count *= 6
	// }
	// count += 1
	// fmt.Println("Jumlah :", count)
}
