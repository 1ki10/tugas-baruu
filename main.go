package main

import "fmt"

func main() {

	//fmt.Println("Hello World")

	/* Input (angka) -> menerima input / masukan
	dari luar ke dalam program */

	// Menerima input ANGKA dari user
	//var angka int
	//fmt.Scanln(&angka)

	//fmt.Println("Hasil inputya adalah", angka)

	//var nama string
	//fmt.Scanln(&nama)

	//fmt.Println("Hasil inputnya adalah", nama)

	//var kondisi bool
	//fmt.Println("Isi variable kondisi", kondisi)

	var angka int
	fmt.Print("masukkan angka untuk dicek: ")
	fmt.Scanln(&angka)

	if angka%2 == 0 {
		fmt.Println("Genap")
	} else {
		fmt.Println("ganjil")
	}

}
