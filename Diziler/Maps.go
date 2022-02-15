package main

import "fmt"

//MAPS
func main(){

	//Standart oluşturma yöntemi

	var ogrenci map[int]string=map[int]string{
		10:"Ahmet",
		21:"Mehmet",
		39:"Buse",
		89:"Cem",
	}

	fmt.Println(ogrenci)

	fmt.Println(ogrenci[39])

	ogrenci[21]="Uğur"

	fmt.Println(ogrenci)

	//key:int  value:string
	plaka:=make(map[int]string)

	plaka[44]="Malatya"
	plaka[35]="İzmir"
	plaka[01]="Adana"


	fmt.Println(plaka)

	fmt.Println(plaka[44])


	//key:string value:string

	unvanlar:=make(map[string]string)

	unvanlar["DR"]="DOKTOR"
	unvanlar["PRF"]="PROFESOR"
	unvanlar["DOC"]="DOCENT"

	fmt.Println(unvanlar["PRF"])

	a:=unvanlar["DR"]

	fmt.Println(a)

	//silme

	delete(unvanlar,"DR")

	fmt.Println(unvanlar)

	fmt.Println(len(unvanlar))


	for key,value:=range unvanlar{
		fmt.Printf("Kısaltılması:%v Tam Hali:%v\n",key,value)
	}






}