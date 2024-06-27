package	piscine

import "ft"

func PrintReverseAlphabet() {
	char := 'z'
	for char >= 'a' {
		ft.PrintRune(char);
		char--
	}
	ft.PrintRune('\n');
}
