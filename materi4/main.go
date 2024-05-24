package main

import (
	"fmt"
	//"strings"
)

func main() {
	// print 1 - 5
	//fmt.Println("1 2 3 4 5")

	// variable=nilai_awal; kondisi/syarat kegiatan berlangsung; increment(aktifitas variable dalam setiap repitisi)
	// i=i+1 -> i++   //  i = i + 2 -> +=2
	//for i := 1; i <= 100; i++ {
	//	fmt.Print(i, " ")
	//}

	//cek bilangan prima
	
	//ar angka int
	//var bantuan = 2
	//var isPrima bool = true
	//fmt.Scanln(&angka)

	//for i := 2; i < angka; i++ {
	//	if angka % i == 0 {
	//		isPrima = false
	//	}
	//}
	
	//for bantuan < angka{
	//	if angka%bantuan == 0 {
	//		isPrima = false
	//	}
	//	bantuan++
	//}
	//fmt.Println("apakah", angka, "adalah bilangan prima ?", isPrima)


	//var kota string ="surabaya" //-> 8
	//i = o
	// huruf S ada diposisi ke brpa? ke 0
	//for i := 0; i < len(kota); i++ {
	//	fmt.Print("posisi i-", i)
	//	fmt.Println(": ", string(kota[i+3]))
	//}

	//var angka int
	//var bantuan = 2
	//var isPrima bool = true
	//fmt.Scanln(&angka)

	//for bantuan < angka{
	//	fmt.Println("nilai bantuan saat ini -", bantuan)
	//	if angka%bantuan == 0 {
	//		isPrima = false
	//		break
	//	}
	//	bantuan++
	//}
	//fmt.Println("apakah", angka, "adalah bilangan prima ?", isPrima)
	
	var nomor = 1
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(nomor, " ")
			nomor++
		}
		fmt.Println()
	}


}