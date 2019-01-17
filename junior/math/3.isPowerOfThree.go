package main

import (
	"math"
	"fmt"
)

/**
给定一个整数，写一个函数来判断它是否是 3 的幂次方。

示例 1:

输入: 27
输出: true

示例 2:

输入: 0
输出: false

示例 3:

输入: 9
输出: true

示例 4:

输入: 45
输出: false

进阶：
你能不使用循环或者递归来完成本题吗？

 */
func run_3()  {
	fmt.Println(isPowerOfThree(1))
}

func isPowerOfThree1(n int) bool {
	if n == 0 {
		return false
	}
	temp := math.Log((float64(n)))
	fmt.Println(temp, float64(n))
	if temp - float64(int(temp)) == 0{
		return true
	}
	return false
}


func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	if n == 1{
		return true
	}
	if n % 3 != 0 {
		return false
	}
	return isPowerOfThree(n/3)
}