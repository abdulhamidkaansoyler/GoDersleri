package main

import "fmt"

func main(){

	//Array statik yapılıdır.List dinamik statik

	arr:=[10]int{1,2,3,4,5,6}

	//arr[15]=99

	arr[2]=99

	fmt.Println(arr)

	//len komutu:bir dizinin kaç elemanlı olduğunu söyler

	fmt.Println(len(arr))

	//cap komutu:bir dizinin kaç elemana kadar alabileceğini söyler.

	fmt.Println(cap(arr))

	for i:=0;i<=len(arr)-1;i++{
		fmt.Println(arr[i])
	}


	

}

