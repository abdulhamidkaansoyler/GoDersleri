package main

import "fmt"

/*
	MÜKEMMEL SAYI;

	Kendi bölünenlerinin toplamının yine kendisine eşit olduğu sayılardır.

	28 = 14+7+4+2+1

	6 = 3+2+1

	496,8128,33550336


 */


func main(){

	var sayi int

	fmt.Print("Sayı Giriniz:")
	fmt.Scanf("%d",&sayi)

	toplam:=0

	for i:=1;i<sayi;i++{
		if sayi%i==0{
			toplam+=i
		}
	}

	if toplam==sayi{
		fmt.Printf("%d sayısı Mükemmel Sayıdır.",sayi)
	}else{
		fmt.Printf("%d sayısı Mükemmel Sayı Değildir.",sayi)
	}


}
