package main

import "fmt"

func main() {
	prices := []int{98, 125, 232, 147, 486}
	items := []string{"消しゴム", "ボールペン", "ノート", "付箋紙", "ボールペン"}

	for i := 0; i < len(prices); i++ {

		fmt.Printf("%s: %d円\n", items[i], prices[i])
	}

}
