package main

import "fmt"

func main(){

	/*
	x:=0

	for x<=5{
		fmt.Println(x)
		x++
	}

	 */

	/*

	for x:=0;x<=10;x++{
		fmt.Println(x)
	}

	 */

	/*
	for i:=0;i<=100;i++{
		if i%2==0{
			fmt.Printf("Çift Sayılar: %d\n",i)
		}
	}

	 */

	toplam:=0

	for i:=0;i<10;i++{
		toplam+=i
	}
	fmt.Println(toplam) // 1+2+3+4+5+6...


}

