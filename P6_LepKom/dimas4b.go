package main

import "fmt"

var nilai1, nilai2 int

func main() {
	defer fmt.Println("-----SELESAI-----")
	fmt.Print("Masukkan Bilangan 1: ")
	fmt.Scan(&nilai1)
	fmt.Print("Masukkan Bilangan 2: ")
	fmt.Scan(&nilai2)

	hasil := nilai1 % nilai2
	fmt.Printf("Sisa Bagi dari Nilai1 / Nilai2 = %d\n", hasil)
}
