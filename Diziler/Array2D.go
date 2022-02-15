package main

import "fmt"

func main(){

	//2 Boyutlu Dizi

	/*

	var dizi=[2][2]int{{0,1},{2,4}}

	fmt.Println(dizi)

	fmt.Println(dizi[0][1])
	fmt.Println(dizi[1][0])


	 */

	var dizi2D=[5][2]int{{0,0},
		{1,2},
		{2,4},
		{3,6},
		{4,8}}

	fmt.Println(dizi2D)

	fmt.Println(dizi2D[3][0])
	fmt.Println(dizi2D[4][1])

	for i:=0;i<5;i++{
		for j:=0;j<2;j++{
			fmt.Printf("Dizi2D[%d][%d] = %d\n",i,j,dizi2D[i][j])
		}
	}


	//3 Boyutlu Diziler

	var dizi3D=[3][3][3]int{{{1,2,3},{1,2,3},{1,2,3}},{{1,2,3},{1,2,3},{1,2,3}},{{1,2,3},{1,2,3},{1,2,3}}}

	fmt.Println(dizi3D[0][2][1])






}