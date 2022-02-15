package main

import (
	"fmt"
	"strconv"
)

func main()  {

	//TYPE CASTING


	/*
	var toplam int=984
	var sayi int=17

	var sonuc=float32(toplam)/float32(sayi)

	var s1=int(sonuc)

	fmt.Println(sonuc)

	fmt.Println(s1)


	 */


	//TYPE CONVERSION


	/*
	var str="1"

	//String Int'e Dönüşüm

	var sayi,_=strconv.Atoi(str)

	fmt.Println(sayi)

	var sonuc=sayi+7

	fmt.Println(sonuc)


	 */

	//Int String'e Dönüşüm

	var sayi=1

	var str=strconv.Itoa(sayi)

	fmt.Println(str)

	fmt.Println("NASILSIN"+str)

























}
