package main

import "fmt"

func main() {
	scores := [5][3]int{
		{20, 98, 40},
		{70, 65, 62},
		{35, 39, 38},
		{82, 96, 48},
		{92, 87, 85},
	}

	for i := 0; i < 5; i++ {
		total := scores[i][0]+scores[i][1]+scores[i][2] > 179 //総得点が180点以上
		balance := scores[i][0] > 59 && scores[i][1] > 59 &&  //全ての課題が60点以上（論理積）
			scores[i][2] > 59
		talent := scores[i][0] > 89 || scores[i][1] > 89 || //どれか一つの課題が90点以上（論理和）
			scores[1][2] > 89

		fmt.Printf("受験者%dさんは\n", i+1)
		if total { //条件を満たさなければ何もしない
			fmt.Println("総合得点で合格しています")
		} //↑↓両方を満たせば両方行う。
		if balance { //条件を満たさ根ければ何もしない
			fmt.Println("全ての課題で合格しています")
		}
		if !total && !balance && talent { //条件tallentのみ満たす
			fmt.Println("一芸に秀でています")
		}

		if total && balance {
			fmt.Println("おめでとうございます") //総得点だけでなく全課題条件もクリア
		} else if total || talent {
			fmt.Println("来週、再挑戦できます") //総得点か、さもなければ一芸を満たす。
		} else {
			fmt.Println("来年、また挑戦してください") //どの条件も満たさない。
		}
		fmt.Println()
	}
}
