package piscine

func IsSorted(f func(a, b int) int, a []int) bool {
	hasNegative := false
	hasPositive := false

	for i := 0; i < len(a)-1; i++ {
		s := f(a[i], a[i+1])
		if s > 0 {
			hasPositive = true
		}
		if s < 0 {
			hasNegative = true
		}
		if hasNegative && hasPositive {
			return false
		}
	}
	return true
}
