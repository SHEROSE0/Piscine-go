package piscine

import "github.com/01-edu/z01"

func DescendComb() {
	for num := 99; num >= 0; num-- {
		for score := num - 1; score >= 0; score-- {
			z01.PrintRune(rune(num/10 + '0'))
			z01.PrintRune(rune(num%10 + '0'))
			z01.PrintRune(' ')
			z01.PrintRune(rune(score/10 + '0'))
			z01.PrintRune(rune(score%10 + '0'))

			if !(num == 1 && score == 0) {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
	}
}
