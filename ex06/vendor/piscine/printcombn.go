package piscine

import "ft"

func power(num int, power int) int{
	var ans int = 1;
	for power > 0 {
		ans = ans * num
		power -= 1
	}
	return ans
}

func ascendingcheck(num int, n int) bool {
	var lastnum int = -100
	var curnum int = -100
	for n > 0 {
		curnum = (num / (power(10, n - 1))) % 10
		if curnum <= lastnum {
			return (false)
		}
		lastnum = curnum
		n -= 1
	}
	return (true)
}

func putnum(num int, n int) {
	var dig int = -1
	for n > 0 {
		dig = (num / (power(10, n - 1))) % 10
		ft.PrintRune(rune(dig + '0'))
		n -= 1
	}
}

func putlast(n int) {
	var i int = 0
	for i < n {
		ft.PrintRune(rune(10 - (n - i) + '0'))
		i++
	}
	ft.PrintRune('\n')
}

func PrintCombn(n int) {
	var num int = 0
	for ((num / (power(10, n - 1))) % 10) < (10 - n) {
		if (ascendingcheck(num, n) == false) {
			num++
		} else {
			putnum(num, n) 
			ft.PrintRune(',')
			ft.PrintRune(' ')
			num++
		}
	}
	putlast(n)
}