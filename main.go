package main

import (
	"fmt"
	"time"
)

func main() {

	//Start
	fmt.Println("Start ", time.Now().Second(), "Detik")

	bilangan := []float64{4, 6, 8, 5, 7}

	// Menghitung jumlah bilangan
	var total float64 = 0
	for _, nilai := range bilangan {
		total += nilai
	}

	// Menghitung rata-rata
	rata_rata := total / float64(len(bilangan))

	// Menampilkan hasil
	fmt.Printf("Rata-rata dari %v adalah %.2f", bilangan, rata_rata)
	fmt.Println(" ")

	//Finish
	fmt.Println("Finish", time.Now().Second(), "Detik")
}
