package piscine

// package main

// func Max(a []int) int {
// 	if len(a) == 0 {
// 		return 0
// 	}
// 	ans := a[0]
// 	for _, char := range a {
// 		if char > ans {
// 			ans = char
// 		}
// 	}
// 	return ans
// }

func Max(a []int) int {
	if len(a) == 0 {
		return 0
	}
	ans := a[0]
	for i := 0; i < len(a)-1; i++ {
		if a[i] > ans {
			ans = a[i]
		}
	}
	return ans
}
