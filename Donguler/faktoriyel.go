package main

import "fmt"

/* 1'den 10'a kadar olan sayıların Faktoriyel


	5!=5x4x3x2x1  = 120
	n!=n x (n-1)!
	1!=1
	0!=1


 */



func main(){

	for i:=1;i<=10;i++{
		faktoriyel:=1
		for j:=1;j<=i;j++{
			faktoriyel*=j  //faktoriyel=faktoriyel*j
		}
		fmt.Println(i,"! =",faktoriyel)
	}









}
