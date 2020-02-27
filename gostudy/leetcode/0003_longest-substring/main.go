package main

import "fmt"

/*
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
示例 1:
输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:
输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:
输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。
*/

func main() {
	str := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(str))
	str = "abcdeaaaabcdaa"
	fmt.Println(lengthOfLongestSubstring(str))
	str = "pwwkew"
	fmt.Println(lengthOfLongestSubstring(str))
}

func lengthOfLongestSubstring(s string) int {
	var start, max int
	for i := 0; i < len(s); i++ {
		for j := start; j < i; j++ {
			if s[j] == s[i] {
				start = j + 1
				break
			}
		}
		if i-start+1 > max {
			max = i - start + 1
		}
	}
	return max
}
