package main

import (
	"fmt"
	"strings"
)

func main(){

	fmt.Println(strings.ToLower("ABDULHAMID KAAN"))
	fmt.Println(strings.ToUpper("abdulhamid kaan"))

	//Index

	fmt.Println(strings.Index("Turkiye","r"))

	//Count

	fmt.Println(strings.Count("Adana","a"))
	fmt.Println(strings.Count("Adana",""))

	//Contains

	fmt.Println(strings.Contains("info@gmail.com","@"))

	//Compare
	k1:="Merhaba"
	k2:="Merhaba Dunya"
	k3:="Merhaba"

	fmt.Println(strings.Compare(k1,k2))
	fmt.Println(strings.Compare(k2,k3))
	fmt.Println(strings.Compare(k3,k1))

	//Return 0, if str1 == str2
	//Return 1, if str1 > str2
	//Return -1, if str1 < str2

	//Replace

	fmt.Println(strings.Replace("Olmak yadak Olmamak","k","z",2))

	//ReplaceAll

	fmt.Println(strings.ReplaceAll("Hey Hey Hey","Hey","Mey"))

	//Fields

	data:="  username  password       email   date"
	fields:=strings.Fields(data)
	fmt.Printf("%q\n",fields)

	//Split

	fmt.Println(strings.Split("a,b,c,d",","))
	fmt.Println(strings.Split("19-11-1987","-"))

	//Split Slice

	tarih:="19-11-1879"
	ayir:=strings.Split(tarih,"-")
	fmt.Println(ayir[2])


	//SplitAfter

	sira:="a,b,c,d"
	ayri:=strings.SplitAfter(sira,",")
	fmt.Println(ayri)
	fmt.Println(ayri[1])









}