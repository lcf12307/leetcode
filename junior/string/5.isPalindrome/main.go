package main

import (
	"fmt"
)
/*
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false
 */
func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	runes := []rune(s)
	for start, end := 0, len(runes) - 1; start < end; start,end = start+1,end-1{
		start = nextValidChar(start, runes, 1, end)
		end = nextValidChar(end, runes, -1, start)
		if start == -1 || end == -1 {
			break
		}
		if runes[start] != runes[end] {
			return false
		}
	}
	return true
}

//获取下一个有效字符，并且替换
func nextValidChar(index int, runes []rune, incr int, limit int) int {
	runes[index] = isValidChar(runes[index])
	for runes[index] == 0 {

		index += incr
		fmt.Println(index, limit)
		if (limit - index) / incr < 0 || index < 0 || index > len(runes) - 1{
			return -1
		}
		runes[index] = isValidChar(runes[index])
	}
	return index
}

//判断是否为有效字符
func isValidChar(c rune) rune  {
	if !(c > 47 && c < 58 || c > 64 && c < 91 || c > 96 && c < 123) {
		return 0
	}
	if c > 64 && c < 91 {
		c += 32
	}
	return c
}

func main()  {
	data := ".,"
	result := isPalindrome(data)
	fmt.Println(result)
}