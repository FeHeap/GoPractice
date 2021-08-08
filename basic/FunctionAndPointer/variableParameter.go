package main

import "fmt"

func main() {
	sum, avg, count := GetScore(90, 82.5, 73, 64.8)
	fmt.Printf("學員共有%d門成績, 總成績為: %.2f, 平均成績為: %.2f", count, sum, avg)
	fmt.Println()
	scores := []float64{92, 72.5, 93, 74.5, 89, 87, 74}
	sum, avg, count = GetScore(scores...)
	fmt.Printf("學員共有%d門成績, 總成績為: %.2f, 平均成績為: %.2f", count, sum, avg)
	fmt.Println()
}

func GetScore(scores ...float64) (sum, avg float64, count int) {
	for _, value := range scores {
		sum += value
		count++
	}
	avg = sum / float64(count)
	return
}
