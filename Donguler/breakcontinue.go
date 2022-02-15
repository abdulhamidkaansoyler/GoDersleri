package main

import "fmt"

func main(){


	/* break
	for i:=1;i<=5;i++{
		if i==3{
			break
		}

		fmt.Println(i)
	}

	 */

	//continue

	for i:=1;i<=5;i++{
		if i==3{
			continue
		}

		fmt.Println(i)
	}


}