package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main(){

	Pisayisi:=math.Pi

	fmt.Println(Pisayisi)

	fmt.Println(math.Pow(2,10))

	fmt.Println(math.Sqrt(64))

	fmt.Println(math.Pow(64,0.5))


	//Mutlak Değer Absolute Value-Float
	fmt.Println(math.Abs(-55.4))



	fmt.Println(math.Sin(37))
	fmt.Println(math.Cos(57))


	fmt.Println(math.Mod(44,3))

	//Yuvarlama
	fmt.Println(math.Round(22.5))

	//Tavan-- Yukarı Doğru Yuvarlama
	fmt.Println(math.Ceil(22.1))


	//Taban--Aşağı Doğru Yuvarlama
	fmt.Println(math.Floor(22.9))


	//Max-Min
	fmt.Println(math.Max(55,2))
	fmt.Println(math.Min(67,43))


	//Sayı Kontrol
	fmt.Println(math.IsNaN(400))


	fmt.Println(math.Log(100))
	fmt.Println(math.Log10(100))


	rand.Seed(time.Now().UnixNano())

	rastgele:=rand.Intn(100)

	fmt.Println(rastgele)



}
