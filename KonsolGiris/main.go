package main

import(
	"os"
	"bufio"
	"fmt"
	"strconv"
)



func main()  {

	/*
	tarayici:=bufio.NewScanner(os.Stdin)

	fmt.Print("Birşeyler Yazınız:")

	tarayici.Scan()

	verigirisi:=tarayici.Text()

	fmt.Printf("Bunu yazdınız: %s",verigirisi)
	*/

	scanner:=bufio.NewScanner(os.Stdin)

	fmt.Print("Hangi Yıl Doğdunuz: ")

	scanner.Scan()

	verigirisi,_:=strconv.ParseInt(scanner.Text(),10,64)

	fmt.Printf("Şuanda %d yaşındasın.",2021-verigirisi)







}
