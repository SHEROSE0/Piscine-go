package piscine

func Rot14(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		if r >= 'a' && r <= 'z' {
			runes[i] = (r-'a'+14)%26 + 'a'
		} else if r >= 'A' && r <= 'Z' {
			runes[i] = (r-'A'+14)%26 + 'A'
		}
	}
	return string(runes)
}
