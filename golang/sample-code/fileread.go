package main

import "fmt"

func main()  {
	var N int
	fmt.Scanf("%d", &N)
	fmt.Println(N)
	//for i:=0; i<N; i++ {
	//	fmt.Println(i)
	//}
	var nextValue, left, right int
	fmt.Scanf("%d %d %d", &nextValue, &left, &right)
	fmt.Println(nextValue, nextValue, right)
}