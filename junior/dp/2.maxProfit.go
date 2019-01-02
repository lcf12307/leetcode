package main

import "fmt"

/**
给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票），设计一个算法来计算你所能获取的最大利润。

注意你不能在买入股票前卖出股票。

示例 1:

输入: [7,1,5,3,6,4]
输出: 5
解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
     注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。

示例 2:

输入: [7,6,4,3,1]
输出: 0
解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。


 */

func run_2()  {
	data := []int{7,1,5,3,6,4}
	fmt.Println(maxProfit(data))
}



func maxProfit(prices []int) int {
	profit := 0
	if len(prices) == 0 {
		return 0
	}

	min, max := prices[0], prices[0]
	for _, d := range prices {
		if d > max {
			max = d
		}
		if d < min {
			min, max = d, d
		}
		if profit < max - min {
			profit = max - min
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	min, max := prices[0], prices[0]
	profit := 0
	for i:=0; i<len(prices) - 1; i++ {
		if prices[i] > prices[i + 1]{
			max = prices[i]
			profit += max - min
			min = prices[i + 1]
		} else if i == len(prices) - 2 {
			max = prices[i + 1]
			profit += max - min
		}
	}
	return profit
}