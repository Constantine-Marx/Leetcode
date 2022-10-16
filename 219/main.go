package main

import (
	"fmt"
	"math"
	"sort"
)

type Number struct {
	Value int
	Count int
}

type Numbers []Number

func (n Numbers) Len() int {
	return len(n)
}

func (n Numbers) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func (n Numbers) Less(i, j int) bool {
	return n[i].Value < n[j].Value
}

func ExistSame(nums []int, k int) bool {
	var Data = make(Numbers, len(nums))
	for i := 0; i < len(nums); i++ {
		Data[i].Count = i
		Data[i].Value = nums[i]
	}
	sort.Sort(Numbers(Data))
	fmt.Println(Data)
	for i := 0; i < len(nums)-1; i++ {
		if Data[i].Value == Data[i+1].Value {
			if int(math.Abs(float64(Data[i].Count-Data[i+1].Count))) <= k {
				fmt.Println(Data[i].Count,Data[i].Value)
				fmt.Println(Data[i+1].Count,Data[i+1].Value)
				return true
			}
		} else {
			fmt.Println(Data[i].Count,Data[i].Value)
			fmt.Println(Data[i+1].Count,Data[i+1].Value)
			fmt.Println("*************")
			continue
		}
	}
	return false
}

func main() {
	nums := []int{0,1,2,3,4,0,0,7,8,9,10,11,12,0}
	k := 1
	if ExistSame(nums,k) {
		fmt.Println("YES!")
	}else {
		fmt.Println("NO!")
	}
}
