package main

import "fmt"

func missingNumber(nums []int) int {
	n := len(nums)
	sum := (1 + n) * n / 2
	sum1 := 0
	for _, v := range nums {
		sum1 += v
	}
	return sum - sum1
}

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(missingNumber(nums))
}
