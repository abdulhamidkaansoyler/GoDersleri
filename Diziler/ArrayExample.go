package main

import "fmt"

//Dizi Elemanlarını Çarpma

func main(){

	var dizi1 [2][3]int
	var dizi2 [2][3]int

	var dizisonuc [2][3]int

	fmt.Print("1.Dizinin Elemanlarını Giriniz:\n")

	for k,r:=range dizi1{
		for l:=range r{
			fmt.Scan(&dizi1[k][l])
			//dizi1[0][0] [0][1]

		}

	}

	fmt.Println("2.Dizinin Elemanlarını Giriniz:\n")

	for m,rr:=range dizi2{
		for n:=range rr{
			fmt.Scan(&dizi2[m][n])
		}
	}


	fmt.Println("Dizilerin Çarpımı Sonucu:\n")

	for i,satir:=range dizisonuc{
		for j:=range satir{
			dizisonuc[i][j]=dizi1[i][j]*dizi2[i][j]

			fmt.Print(dizisonuc[i][j],"\t")
		}

		fmt.Println()
	}




}