package main

import "fmt"

// defer へ渡した関数の実行を、呼び出し元の関数の終わり(returnする)まで遅延させる。
// defer へ渡した関数の引数は、すぐに評価されるが、その関数自体は呼び出し元の関数がreturnするまで実行されない。
func main() {
	defer fmt.Println("World!")

	fmt.Println("Hello")
}
