package piscine

func ShoppingListSort(slice []string) []string {
	n := len(slice)

	for i := 0; i < n; i++ {
		for j := 0; j < n-1-i; j++ {
			if len(slice[j]) > len(slice[j+1]) {
				temp := slice[j]
				slice[j] = slice[j+1]
				slice[j+1] = temp
			}
		}
	}
	return slice
}
