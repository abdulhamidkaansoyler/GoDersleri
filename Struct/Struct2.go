package main

import "fmt"

type calisan struct{
	isim string
	yas int
	maas int
	isHaveKid bool //isMarried

}

type mudur struct{
	calisan
	unvan string

}


func main(){

	a1:=calisan{
		isim:"Ahmet",
		yas:46,
		maas:4250,
		isHaveKid:true,
	}

	fmt.Println(a1)

	y1:=mudur{
		calisan:calisan{
			isim:"Volkan",
			yas:36,
			maas:9750,
			isHaveKid:false,

		},
		unvan:"Personel Muduru",
	}

	fmt.Println(y1)

	kurucu:=struct{
		isim string
		sermaye int
	}{isim:"Tahsin",sermaye:750000}

	fmt.Println(kurucu)

	fmt.Println(kurucu.sermaye)


}
