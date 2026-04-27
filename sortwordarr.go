package piscine

func SortWordArr(a []string) {
	lenA := len(a)

	for i := 0; i < lenA-1; i++ {
		for j := 0; j < lenA-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
}
