package main

import "fmt"

func main() {
	var umur int
	var nilai int
	fmt.Print("masukkan umur: ")
	fmt.Scanln(&umur)
	fmt.Print("masukkan nilai: ")
	fmt.Scanln(&nilai)
	
	//if nilai > 80{
	//	fmt.Println("nilai bagus")
	//}else{
	//	fmt.Println("ayo lebih semangat")
	//}
	
	
	if umur > 17 && nilai > 90 {
		fmt.Println("mantap")
	} else if umur > 17 {
		fmt.Println("sudah dapat KTP")
	}


	//if umur > 17 && nilai > 90 {
	//	fmt.Println("mantap")
	//} 

	//if umur > 17 {
	//	fmt.Println("sudah dapat KTP")
	//}

	var warna string
	fmt.Scanln(&warna)
	switch warna {
	case "red":
		fmt.Println("ini warna merah")
	case "blue":
		fmt.Println("ini warrna biru")
	default :
		fmt.Println("warna ", warna)
	}
	
}