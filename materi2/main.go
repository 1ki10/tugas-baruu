package main

import (
	"fmt"
	"strconv"
)

func main() {
	//var isBenar bool = true
	//var nama string
	var umur int = 10
	var pi float32 = 3.14

	var hasil = umur + int(pi) // type cassting -> melakukan perubahan tioe data secara paksa
	fmt.Println(hasil)

	var hasil2 = float32(umur) + pi
	fmt.Println(hasil2)

	var cnv = strconv.Itoa(umur)

	var nama = "tono" + string(cnv)
	fmt.Println(nama)
}
