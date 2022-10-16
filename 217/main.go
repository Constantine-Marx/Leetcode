package main

import (
	"fmt"
	"sort"
)

func ExistSame(nums []int) bool {
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1]{
			return true
		} 
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	if ExistSame(nums) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
