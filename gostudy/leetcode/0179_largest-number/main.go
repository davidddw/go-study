package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

/*
给定一组非负整数，重新排列它们的顺序使之组成一个最大的整数。

示例 1:

输入: [10,2]
输出: 210
示例 2:

输入: [3,30,34,5,9]
输出: 9534330
*/

func main() {
	num1 := []int{3, 30, 34, 5, 9}
	fmt.Println(largestNumber(num1))
}

func largestNumber(nums []int) string {
	strs := make([]string, len(nums))
	for i := 0; i < len(nums); i++ {
		strs[i] = strconv.Itoa(nums[i])
	}
	sort.Slice(strs, func(i, j int) bool {
		return (strs[i] + strs[j]) > (strs[j] + strs[i])
	})
	numsStr := strings.Join(strs, "")
	numsStr = strings.TrimLeft(numsStr, "0")
	if numsStr == "" {
		return "0"
	}
	return numsStr
}
