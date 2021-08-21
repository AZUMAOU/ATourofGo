package main

import "fmt"

func main() {

	pattern := [3]string{"グー", "チョキ", "パー"}

	fmt.Println("****じゃんけん勝敗リスト****")

	for me := 0; me < 3; me++ {

		fmt.Printf("わたしが%sのとき:\n", pattern[me])

		for you := 0; you < 3; you++ {

			score := (3 + me - you) % 3 //%=余り

			if score == 2 {
				fmt.Printf("あなたが%sなら私の勝ち\n", pattern[you])
			} else if score == 1 { //else=さもなくば
				fmt.Printf("あなたが%sなら私の負け\n", pattern[you])
			} else {
				fmt.Printf("あなたが%sならあいこ\n", pattern[you])
			}
		}
		fmt.Println()
	}

}
