package main

import (
	"fmt"
)

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/

func main() {
	num1 := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(num1))
}

func maxSubArray(nums []int) int {
	ret := nums[0]
	var sum int
	for _, v := range nums {
		if sum > 0 {
			sum += v
		} else {
			sum = v
		}
		if ret < sum {
			ret = sum
		}
	}
	return ret
}
