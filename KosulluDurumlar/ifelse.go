package main

import "fmt"

/* if (kosul){
	islemler
	}
   else{
	islemler
	}

 */


func main(){

	/*

	isim:="Ali"

	if isim=="Ahmet"{
		fmt.Println("Hoşgeldin Ahmet")
	}else if isim=="Mehmet"{
		fmt.Println("Hoşgeldin Mehmet")
	}else{
		fmt.Println("Yanlış Kişi")
	}


	 */

	yas:=16

	fmt.Println("18 yaşından küçükler giremez")

	if yas>=18{
		fmt.Println("Mekana Hoşgeldiniz")

	}else if yas<18 && yas>=14{
		fmt.Println("Ailenle girebilirsin")
	}else{
		fmt.Println("Mekana giremezsin")
	}


















}
