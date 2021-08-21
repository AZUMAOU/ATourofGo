package main

import "fmt"

func main() {
	scores := [][3]int{
		{50, 55, 48},  //0
		{72, 85, 66},  //1
		{31, 100, 54}, //2
		{64, 68, 56},  //3
		{87, 67, 73}}  //4

	for i := 0; i < len(scores); i++ {
		sum := 0
		for k := 0; k < len(scores[i]); k++ {
			sum += scores[i][k]

		}
		fmt.Printf("受験者%d:平均値%d点\n", i+1, sum/3)
	}
}
