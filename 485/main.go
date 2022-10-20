package main

import "fmt"

func findMaxConsecutiveOnes(nums []int) int {
	cnt := 0
	max := 0
	for _, v := range nums {
		if v == 1 {
			cnt++
		} else {
			cnt = 0
		}
		if max <cnt{
			max = cnt
		}
	}
	return max
}

func main() {
	fmt.Println(findMaxConsecutiveOnes([]int{1,0,1,1,0,1}))
}