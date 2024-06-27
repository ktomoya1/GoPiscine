package	piscine

import "ft"

func PrintAlphabet() {
	char := 'a'
	for char <= 'z' {
		ft.PrintRune(char);
		char++
	}
	ft.PrintRune('\n');
}
