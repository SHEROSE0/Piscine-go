package piscine

func LoafOfBread(str string) string {
	count := 0
	for _, c := range str {
		if c != ' ' {
			count++
		}
	}
	if count == 0 {
		return "\n"
	}
	if count < 5 {
		return "Invalid Output\n"
	}
	result := ""
	current := ""
	i := 0
	for i < len(str) {
		if str[i] != ' ' {
			current += string(str[i])
		}
		if len(current) == 5 {
			if result != "" {
				result += " "
			}
			result += current
			current = ""
			if i+1 < len(str) {
				i++
			}
		}
		i++
	}
	if len(current) > 0 {
		if result != "" {
			result += " "
		}
		result += current
	}
	result += "\n"
	return result
}
