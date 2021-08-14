package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	playGame()
}

func playGame() {
	target := generateRandNum(10, 100)
	fmt.Println("請輸入隨機數:")
	fmt.Println("--------------------")

	count := 0
	for {
		count++
		yourNum := 0
		fmt.Scanln(&yourNum)

		if yourNum < target {
			fmt.Println("Too small\n")
		} else if yourNum > target {
			fmt.Println("Too big\n")
		} else {
			fmt.Println("Right!")
			fmt.Printf("您一共猜測了%d次!\n", count)
			break
		}
		alertInfo(count, target)
	}
}

func alertInfo(count, target int) {
	if count >= 6 {
		fmt.Printf("\n呵呵，猜了%d次都沒猜中!\n", count)
		fmt.Println("正確數字是:", target)
		os.Exit(0)
	}
}

func generateRandNum(min, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max - min + 1) + min
}