package main

import "fmt"

func main(){

	//var x int=10
	x:=10

	fmt.Println(x)
	fmt.Println(&x)



	p:=&x

	//var p *int=&x

	fmt.Println(p)

	// *

	fmt.Println(*p)

	*p=44

	fmt.Println(x)

	fmt.Printf("%T,%v\n",x,x)
	fmt.Printf("%T,%v\n",p,p)

	var z string="Merhaba"

	var pz *string=&z

	fmt.Println(z,pz)

	//Google Etkisi

	//v2 rocket

	//2.dünya savaşı,alman rocket








}
