package main

import "fmt"

func getfactors(thenum int) { //func(関数を定義する際に初めに書く。)
	// getfactors(関数名)
	// (thenum(変数名)[]int(データ型を指定))
	//int(戻り値) {
	fmt.Println(thenum)
	for n := 0; n < thenum; n++ {
		switch {
		case n == 0:
		case n == 1:
		case thenum%n == 0:
			thenum /= n
			fmt.Printf("を%dで割ると、%d\n", n, thenum)
		default:
		}
	}
	fmt.Println()
}

func main() {
	getfactors(2310)
	getfactors(37789)
	getfactors(1256997)
}
