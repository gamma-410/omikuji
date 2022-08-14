package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	const (
		str  string = "あなたの運勢は"
		desu string = "です"
	)
	var i int

	for i < 10 {
		rand.Seed(time.Now().UnixNano())
		var randomInt int = rand.Intn(5)

		fmt.Println("🥠", i+1, "回目のおみくじ")

		switch randomInt {
		case 0:
			fmt.Println(str + "「凶」" + desu + "\n")
		case 1:
			fmt.Println(str + "「小吉」" + desu + "\n")
		case 2:
			fmt.Println(str + "「中吉」" + desu + "\n")
		case 3:
			fmt.Println(str + "「吉」" + desu + "\n")
		case 4:
			fmt.Println(str + "「大吉」" + desu + "\n")
		default:
			fmt.Println("Error")
		}
		i++
	}

}
