package	piscine

import "ft"

func PrintDigits() {
	char := '0'
	for char <= '9' {
		ft.PrintRune(char);
		char++
	}
	ft.PrintRune('\n');
}
