package main

import "fmt"

// fizzbuzz は以下のルールに基づいて print する
//   1. n が 3 の倍数の場合、"fizz" を表示
//   2. n が 5 の倍数の場合、"buzz" を表示
//   3. n が 3 と 5 の公倍数の場合、"fizzbuzz" を表示
//   4. それ以外の場合、数字をそのまま表示
// 例)
// 1 -> "1"
// 2 -> "2"
// 3 -> "fizz"
// 5 -> "buzz"
// 15 -> "fizzbuzz"
func main() {
	for i := 1; i <= 100; i++ {

		score3 := (i) % 3
		score5 := (i) % 5

		if score3 == 0 && score5 == 0 {
			fmt.Printf("%d.fizzbuzz\n", i)
		} else if score3 == 0 {
			fmt.Printf("%d.fizz\n", i)
		} else if score5 == 0 {
			fmt.Printf("%d.buzz\n", i)
		} else {
			fmt.Printf("%d\n", i)
		}
	}
	// 1 から 100 までの fizzbuzz をここに書く
}
