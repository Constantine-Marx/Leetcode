package main

import (
	"fmt"
)

type Table struct {
	value bool
	key   int
}

func findDisappearedNumbers(nums []int) []int {
	res := []int{}
	ln := len(nums)
	var tmp []Table
	for i := 0; i < ln; i++ {
		var t Table = Table{false, i + 1}
		tmp = append(tmp, t)
	}
	for i := 0; i < ln; i++ {
		tmp[nums[i]-1].value = true
	}
	for _, v := range tmp {
		if !v.value {
			res = append(res, v.key)
		}
	}
	return res
}

func main() {
	fmt.Println(findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1}))
}
