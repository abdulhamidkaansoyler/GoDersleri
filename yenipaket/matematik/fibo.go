package matematik


func Fibo(n int)int{
	a:=0
	b:=1

	for i:=0;i<n;i++{
		gecici:=a
		a=b
		b=gecici+a
	}
	return a

}

