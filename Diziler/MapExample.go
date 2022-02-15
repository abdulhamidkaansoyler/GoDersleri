package main

import "fmt"

func main(){

	sozluk:=make(map[string]string)

	var eng string
	var tr string


	for i:=1;i<=5;i++{

		fmt.Print("İngilizce Kelime Giriniz:")
		fmt.Scan(&eng)

		fmt.Print("Türkçe Karşılığını Giriniz:")
		fmt.Scan(&tr)


		sozluk[eng]=tr


	}

	fmt.Println(sozluk)

	fmt.Print("Silmek İstediğiniz İngilizce Kelimeyi Giriniz:")
	fmt.Scan(&eng)

	delete(sozluk,eng)

	for k,v:=range sozluk{
		fmt.Printf("ENG:%s \t TR:%s \n",k,v)
	}






}