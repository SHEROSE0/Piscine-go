package piscine

// package main

func ActiveBits(n int) int {
	count := 0
	if n > 0 {
		for {
			if n%2 == 1 {
				count++
			}
			n = n / 2
			if n == 0 {
				break
			}
		}
	}
	return count
}
