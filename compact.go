package piscine

func Compact(ptr *[]string) int {
	slice := *ptr
	count := 0

	for _, s := range slice {
		if s != "" {
			count++
		}
	}

	result := make([]string, count)
	index := 0
	for _, s := range slice {
		if s != "" {
			result[index] = s
			index++
		}
	}

	*ptr = result
	return count
}
