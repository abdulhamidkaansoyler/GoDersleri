package main

import "fmt"

type Vatandas struct{
	isim string
	tcno int
}

func (v Vatandas) AyarlaValue(){ //method on value
	v.isim="Betül"
	v.tcno=444555
}

func (v *Vatandas) AyarlaPointer(){//method on pointer
	v.isim="Cem"
	v.tcno=999999
}


func main(){

	vat1:=Vatandas{"Abdulhamid Kaan",123456}

	var ptr *Vatandas=&vat1

	fmt.Println(ptr)

	fmt.Println(&(ptr.isim))
	fmt.Println(&(ptr.tcno))

	fmt.Println(*ptr)

	ptr.isim="Ahmet"

	fmt.Println(vat1)

	fmt.Println()

	ptr.AyarlaValue()
	fmt.Println(vat1)

	ptr.AyarlaPointer()
	fmt.Println(vat1)





}
