package main

import "fmt"

func main(){

	//var sayi1 int=5

	//fmt.Printf("Benim Yaşım: %d",sayi1)

	//var sayi2 float32=3.1415925

	//fmt.Printf("PI SAYISI: %.2f",sayi2)

	//var dogruMu bool=true

	//fmt.Printf("%t",dogruMu)

	// %c %p

	/*

	var isim="Ahmet"
	var sehir="Ankara"
	var yas = 35

	fmt.Printf("Benim adım %s, %d yaşındayım ve %s doğdum",isim,yas,sehir)
	*/


	var sayi1 int=54
	var sayi2 float32=5.6434
	var yanlisMi bool=false
	var str string="Betül"

	fmt.Printf("%T %T %T %T",sayi1,sayi2,yanlisMi,str)


	var mesaj string="Merhaba Dünya"
	var yil int=2022

	var tummesaj=fmt.Sprintf("%d yılına %s",yil,mesaj)
	fmt.Println(tummesaj)




}
