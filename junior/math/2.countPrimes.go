package main

import "fmt"

func run_2()  {
	fmt.Println(countPrimes(499979))
}

func countPrimes(n int) int {
	primes := []bool {}
	count := -2
	for i := 0; i < n; i++ {
		if i < 4 {
			primes = append(primes, true)
			count ++
			continue
		}
		if i % 2 == 0 || i % 3 == 0{
			primes = append(primes, false)
			continue
		}

		for j:=2;j< i/2;j++  {
			if primes[j] == true && i % j == 0 {
				primes = append(primes, false)
				break
			}
		}
		if len(primes) < i + 1 {
			primes = append(primes, true)
			count ++
		}
	}
	if count < 0 {
		return 0
	}
	return count
}
