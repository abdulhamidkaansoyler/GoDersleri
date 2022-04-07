package main

import (
	"fmt"
	"os"
)

//try catch go'da yok
//errors,fmt

//nil-->null

func main(){

	//var sayi int
	//
	//fmt.Println(sayi)

	//scanner:=bufio.NewScanner(os.Stdin)
	//
	//fmt.Print("Sayı Giriniz:")
	//scanner.Scan()
	//
	//a,err:=strconv.ParseInt(scanner.Text(),10,64)
	//
	//if err != nil{
	//	//fmt.Println(err.Error())
	//	log.Fatal("HATA",err)
	//
	//}
	//
	//fmt.Println(a)


	file,err:=os.Open("/main.go")
	if err !=nil{
		fmt.Println(err)
		//fmt.Println(errors.New("DOSYA YOK"))
	}

	fmt.Println(file.Name(),"Başarılı bir şekilde açıldı.")











}
