package main

import (
	"fmt"
	"time"
)

func main(){

	//Şuanki Zaman
	zaman:=time.Now()
	fmt.Println(zaman)

	fmt.Println(zaman.Day())
	fmt.Println(zaman.Month())
	fmt.Println(zaman.Year())

	fmt.Println()

	fmt.Println(zaman.Hour())
	fmt.Println(zaman.Minute())
	fmt.Println(zaman.Second())

	zamanstring:=zaman.String()
	fmt.Println(zamanstring)


	//sleep

	//fmt.Println("Merhaba")
	//time.Sleep(time.Second*5)
	//fmt.Println("Merhaba")

	//UTC

	zamanUTC:=time.Now().UTC()
	fmt.Println(zamanUTC)

	//Date Kendi Tarihimizi Oluşturmak

	dogumgunu:=time.Date(1992,11,30,12,30,0,0,time.Local)

	fmt.Println(dogumgunu)

	fmt.Printf("%v-%v-%v\n",dogumgunu.Day(),dogumgunu.Month(),dogumgunu.Year())


	//2 tarih arasındaki fark
	baslangic:=time.Date(1915,3,18,0,0,0,0,time.UTC)
	bugun:=time.Now().UTC()

	fark:=bugun.Sub(baslangic).Hours()/24

	fmt.Println(fark)

	kacgunoldu:=bugun.Sub(dogumgunu).Hours()/24

	fmt.Println(kacgunoldu)


	var d time.Duration=1000000000

	fmt.Println(d.Hours())

	//Tarih Ekleme

	suankizaman:=time.Now()

	yarin:=suankizaman.AddDate(15,3,1)

	fmt.Println(yarin)

	ikigunsonra:=suankizaman.Add(time.Hour*48)

	fmt.Println(ikigunsonra)

	//Kendimize Özel Zaman Biçimi

	hafta:=time.Hour*7*24

	fmt.Println(hafta)








}
