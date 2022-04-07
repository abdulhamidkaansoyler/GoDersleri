package main

import "fmt"

//0,1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, …

//f(n)=f(n-1)+f(n-2)

func fibore(n int)int{
	if n<=1{
		return n
	}
   //n=5
	return fibore(n-1)+fibore(n-2)
}


func main(){

	for i:=0;i<=20;i++{
		sonuc:=fibore(i)

		fmt.Printf("Fibonacci %d = %d\n",i,sonuc)
	}


}