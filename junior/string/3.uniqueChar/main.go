package main

import (
	"fmt"
)

/*
给定一个字符串，找到它的第一个不重复的字符，并返回它的索引。如果不存在，则返回 -1。

案例:

s = "leetcode"
返回 0.

s = "loveleetcode",
返回 2.


注意事项：您可以假定该字符串只包含小写字母。
*/
func firstUniqChar(s string) int {
	if s == "" {
		return -1
	}
	runes := []rune(s)
	var index int
	temp := make(map[rune]bool)
	for _, r := range runes {
		if _, ok := temp[r]; ok {
			temp[r] = true
			if r == runes[index] {
				for temp[runes[index]] {
					index++
					if index >= len(runes) {
						return -1
					}
				}
			}
		} else {
			temp[r] = false
		}
	}
	return index
}

func main() {
	data := ""
	result := firstUniqChar(data)
	fmt.Println(result)
}
