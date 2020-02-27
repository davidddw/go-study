package main

import "fmt"

/*
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个
整数，并返回他们的数组下标。
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
*/

func main() {
	nums := [...]int{2, 7, 7, 15}
	target := 9
	fmt.Println(twoSum(nums[:], target))
}

func twoSum(nums []int, target int) []int {
	ret := make([]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				ret = append(ret, i, j)
			}
		}
	}
	return ret
}

func twoSum1(nums []int, target int) []int {
	var ret []int
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		targetNum := target - nums[i]
		if v, ok := m[targetNum]; ok {
			ret = append(ret, v, i)
		}
		m[nums[i]] = i
	}
	return ret
}
