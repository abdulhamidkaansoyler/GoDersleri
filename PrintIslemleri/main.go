package main

import "fmt"

func main(){

	fmt.Println("Merhaba\n\n\nDünya") // \n \t
	fmt.Println("Hello\t\t\tWorld")

	fmt.Println("Abdulhamid Kaan Söyler'in \"Dünyası\" ")

	fmt.Println("Malatya"+"Kayısısı"+"Ünlüdür")
	fmt.Println("Malatya","Kayısısı","Ünlüdür")

	fmt.Println("Lorem İpsum" +
		"Merhaba Dünya" +
		"Hello World")


	isim:="Abdulhamid Kaan"
	soyisim:="Söyler"
	sehir:="Malatya"

	fmt.Println(isim,soyisim,sehir)

	//fmt.Println(len(isim))

	var stringUzunlugu=len(isim)

	fmt.Println(stringUzunlugu)








}
