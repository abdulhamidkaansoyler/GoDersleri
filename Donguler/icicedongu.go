package main

import "fmt"

//İÇ İÇE DÖNGÜ(NESTED LOOPS)

func main(){

	/*

	for i:=1;i<=5;i++{
		for j:=1;j<=5;j++{
			fmt.Printf("İ:%d J:%d \n",i,j)
		}
	}


	 */

	for i:=1;i<=10;i++{
		fmt.Println("*********")
		for j:=1;j<=10;j++{
			fmt.Printf("%d X %d  = %d\n",i,j,i*j)
		}
	}





}
