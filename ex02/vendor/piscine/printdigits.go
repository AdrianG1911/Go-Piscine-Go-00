package piscine

import "ft"

func PrintDigits() {
	var r rune = '0'
	for r <= '9' {
		ft.PrintRune(r)
		r++
	}
	ft.PrintRune('\n')
}