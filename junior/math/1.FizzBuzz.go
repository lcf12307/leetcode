package main

import (
	"strconv"
	"fmt"
)

func run_1()  {

	n := 15
	fmt.Println(fizzBuzz(n))
}

func fizzBuzz(n int) []string {
	result := []string {}
	for i:=0; i < n; i++ {
		if (i+1) % 3 == 0 && (i + 1) % 5 == 0 {
			result = append(result, "FizzBuzz")
		} else if (i + 1) % 3 == 0 {
			result = append(result, "Fizz")
		} else if (i + 1) % 5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, strconv.Itoa(i+1))
		}
	}
	return result
}