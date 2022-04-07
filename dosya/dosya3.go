package main

import (
	"log"
	"os"
)

func main(){

	//os.Remove("veri.txt")

	err:=os.Remove("ahks/veri.html")
	if err !=nil{
		log.Fatal(err)
	}




}
