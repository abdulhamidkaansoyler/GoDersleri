package main

import "fmt"

//struct  structure

type ogrenci struct{
	//field fieldtype

	isim string
	soyisim string
	yas int
	notlar []int
	isMezunMu bool

}

func main(){

	var o1 ogrenci //var sayi int

	o1.isim="Abdulhamid Kaan"
	o1.soyisim="Söyler"
	o1.yas=29
	o1.notlar=[]int{55,66,77}
	o1.isMezunMu=true

	fmt.Println(o1)

	//2. gösteriliş şekli

	o2:=ogrenci{
		isim:"Betül",
		//soyisim:"Gürbüz",
		//yas:28,
		notlar:[]int{80,90,95},
		isMezunMu:false,
	}

	fmt.Println(o2)

	o3:=ogrenci{"Ahmet","Yılmaz",49,[]int{5,6,7},false}

	fmt.Printf("%v",o3)




}
