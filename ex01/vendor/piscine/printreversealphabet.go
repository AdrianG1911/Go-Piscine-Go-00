package piscine

import "ft"

func PrintReverseAlphabet() {
	var r rune = 'z'
	for r >= 'a' {
		ft.PrintRune(r)
		r -= 1;
	}
	ft.PrintRune('\n')
}