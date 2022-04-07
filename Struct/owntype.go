package main

import (
	"fmt"
	"time"
)

type ustunf float64

func (y ustunf) CarpBes() ustunf{
	return y*5.2
}

type Aylar int

const(
	Ocak Aylar=1+iota
	Subat
	Mart
	Nisan
	Mayis
	Haziran
	Temmuz
	Agustos
	Eylul
	Ekim
	Kasim
	Aralik

)




// abdulhamid kaan bdlhmd kn

func main(){

	var u1 ustunf=5.4

	fmt.Printf("%T,%v\n",u1,u1)

	var u2 float64=6.77

	fmt.Println(u1+ustunf(u2))

	y1:=ustunf(50.5)

	fmt.Println(y1.CarpBes())

	_,ay,_:=time.Now().Date()
	fmt.Println(ay)

	fmt.Println(Aylar(ay))

	fmt.Println(Aylar(Nisan))

	fmt.Printf("%T,%v",Aylar(ay),Aylar(ay))













}
