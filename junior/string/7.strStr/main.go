package main

import (
	"fmt"
	)
/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2

示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1

说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。



 */
 //使用kmp算法
func strStr(haystack string, needle string) int {

	hRune := []rune(haystack)
	nRune := []rune(needle)
	if len(nRune) == 0 {
		return 0
	}
	next := findNext(nRune)
	temp := 0
	for i:=0; i < len(hRune); i++ {
		if hRune[i] == nRune[temp] {
			index := i+1
			temp ++
			if temp == len(nRune) {
				return i
			}
			if index > len(hRune) - 1 {
				return -1
			}
			for hRune[index] == nRune[temp]  {
				index ++
				temp ++
				if temp > len(nRune) - 1 {
					return i
				}
				if index > len(hRune) - 1 {
					return -1
				}
			}
			i = i + next[temp]
			temp = 0
		}
	}
	return -1
}

//计算模式串的next
func findNext(runes []rune) []int {
	var next  = []int {0}
	temp := 0
	for index := 1; index < len(runes); index++{
		if runes[index] != runes[0] {
			next = append(next, 0)
			continue
		}
		for index < len(runes) && runes[index] == runes[temp]  {
			temp ++
			index ++
			next = append(next, temp)
		}
		temp = 0
		index --
	}
	return next
}
func main()  {
	data2 := "aaa"
	data1 := "aaa"
	result := strStr(data2, data1)
	fmt.Println(result)
}