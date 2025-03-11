package piscine

import "ft"

func putnum(num int) {
	if num < 10 {
		ft.PrintRune('0')
	} else {
		ft.PrintRune(rune((num / 10) + '0'))
	}
	ft.PrintRune(rune((num % 10) + '0'))
}

func PrintComb2() {
	var num1 int = 0
	var num2 int = 0

	for num1 < 98 {
		if num1 >= num2 {
			num2++;
		}
		if num2 == 100 {
			num2 = 0
			num1++;
		}
		if num1 < num2 {
			putnum(num1)
			ft.PrintRune(' ')
			putnum(num2)
			ft.PrintRune(',')
			ft.PrintRune(' ')
			num2++
		}
	}
	putnum(98)
	ft.PrintRune(' ')
	putnum(99)
	ft.PrintRune('\n')
}