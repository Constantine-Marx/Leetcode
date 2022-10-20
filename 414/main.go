package main

import (
	"fmt"
	"sort"
)

func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	cnt := 1
	tmp := nums[0]
	for _, v := range nums {
		if cnt == 3 {
			break
		}
		if v != tmp {
			cnt++
		}
		tmp = v
	}
	return tmp
}

func main() {
	fmt.Println(thirdMax([]int{2, 2, 3, 1}))
}
