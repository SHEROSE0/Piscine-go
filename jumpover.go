package piscine

// package main

func JumpOver(str string) string {
	if str == "" {
		return "\n"
	}
	if len(str) < 3 {
		return "\n"
	}
	result := ""
	for i := 2; i < len(str); i = i + 3 {
		result += string(str[i])
	}
	return result + "\n"
}
