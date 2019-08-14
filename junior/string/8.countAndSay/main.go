package main

import (
	"fmt"
	"strconv"
)

/*
报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221

1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。
1211
给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。



示例 1:

输入: 1
输出: "1"

示例 2:

输入: 4
输出: "1211"

Go

*/
func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	read := countAndSay(n - 1)
	runes := []rune(read)
	count := 0
	temp := runes[0]
	result := ""
	for i, r := range runes {
		if r == temp {
			count++

		} else {
			result += strconv.Itoa(count) + string(temp)
			temp = r
			count = 1
		}
		if i == len(runes)-1 {
			result += strconv.Itoa(count) + string(temp)
		}
	}
	return result
}

func main() {
	data := 5
	result := countAndSay(data)
	fmt.Println(result)
}
