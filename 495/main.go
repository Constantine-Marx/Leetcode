package main

import "fmt"

func findPoisonedDuration(timeSeries []int, duration int) int {
	res := 2
	for i := 1; i < len(timeSeries); i++ {
		res += 2
		if timeSeries[i] < timeSeries[i-1]+2 {
			res -= 2 - (timeSeries[i] - timeSeries[i-1])
		}
	}
	return res
}

func main() {
	fmt.Println(findPoisonedDuration([]int{1, 4}, 2))
}
