package piscine

import "ft"

func PrintAlphabet(){
	var r rune = 'a'
	for r <= 'z' {
		ft.PrintRune(r);
		r++;
	}
	ft.PrintRune('\n')
}
