package main

import "fmt"

func moveZeroes(nums []int) {
	for i := len(nums) - 1; i >= 0; i-- {
		for j := len(nums) - 2; j >= 0; j-- {
			if nums[j] == 0 && nums[j+1] != 0 {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

func main() {
	nums := []int{0}
	moveZeroes(nums)
	fmt.Println(nums)
}
