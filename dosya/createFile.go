package main

import (
	"log"
	"os"
)

func main(){

	dosya,err:=os.Create("veri.txt") //  ../../veri.html
	if err !=nil{
		log.Fatal(err)
	}

	dosya.Close()



















}
