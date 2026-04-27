package piscine

// package main

func Join(strs []string, sep string) string {
	if len(strs) == 0 {
		return "\n"
	}
	result := ""
	for i := 0; i < len(strs); i++ {
		if i == len(strs)-1 {
			result += string(strs[i])
		} else {
			result += string(strs[i]) + sep
		}
	}
	return result
}
