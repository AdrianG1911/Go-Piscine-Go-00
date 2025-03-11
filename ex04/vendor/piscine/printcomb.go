package piscine

import "ft"

func printnum(f int, s int, t int) {
	ft.PrintRune(rune('0' + f))
	ft.PrintRune(rune('0' + s))
	ft.PrintRune(rune('0' + t))
}

func PrintComb() {
	var f int = 0
	var s int = 0
	var t int = 0
	for f < 7 {
		if (t <= s || t <= f || s <= f) {
			t++;
		}
		if (t == 10) {
			t = 0
			s++;
		}
		if (s == 10) {
			s = 0
			f++;
		}
		if (f < s && s < t) {
			printnum(f, s, t)
			ft.PrintRune(',')
			ft.PrintRune(' ')
			t++
		}
	}
	printnum(7, 8, 9)
	ft.PrintRune('\n')
}