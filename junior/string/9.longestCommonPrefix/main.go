package main

import (
	"fmt"
	)
/*
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"

示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。

说明:

所有输入只包含小写字母 a-z 。

 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	var temp [][]rune
	min := len(strs[0])
	for index, data := range strs {
		temp = append(temp, []rune(data))
		if min > len(temp[index]){
			min = len(temp[index])
		}
	}
	var i int
	for i = 0; i < min; i++ {
		r := temp[0][i]
		for j := 0; j < len(temp); j++{
			if r != temp[j][i] {
				return string(temp[0][:i])
			}
		}
	}
	return string(temp[0][:i])
}

func main()  {
	data := []string{"flower", "flow", "flight"}
	result := longestCommonPrefix(data)
	fmt.Println(result)
}