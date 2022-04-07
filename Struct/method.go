package main

import "fmt"

/*

func(alıcı_adı Türü) fonksiyon_adı(parametre_listesi)(dönüş_türü){
 Kod
}

 */

type Ogrenci struct{
	isim string
	yas int
	notlar []int
	gectiMi bool
}

func(o Ogrenci) AdGoster() string{
	return o.isim
}

func(ogr Ogrenci) BilgileriGoster(){
	fmt.Println("Öğrenci Adı:",ogr.isim)
	fmt.Println("Öğrenci Yas:",ogr.yas)
	fmt.Println("Öğrenci Notlar:",ogr.notlar)
	fmt.Println("Öğrenci Geçti Mi:",ogr.gectiMi)
}

func(not Ogrenci) Ortalama() float64{
	toplam:=0
	for _,i:=range not.notlar{
		toplam+=i
	}

	return float64(toplam)/float64(len(not.notlar))

}

func(gec Ogrenci) GectiMiKaldiMi() bool{
	if gec.Ortalama()<=50{
		return false
	}
	return true
}


func main(){

	o1:=Ogrenci{"AHKS",29,[]int{1,6,7,84,95},false}

	fmt.Println(o1.AdGoster())

	o1.BilgileriGoster()

	fmt.Println(o1.Ortalama())

	fmt.Println(o1.GectiMiKaldiMi())




}
