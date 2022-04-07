package main

import (
	"log"
	"os"
)

func main(){

	err:=os.Mkdir("../xml",os.ModePerm)

	if err !=nil{
		log.Fatal(err)
	}













}