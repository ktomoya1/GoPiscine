package	piscine

import "ft"

func PrintComb() {
	hundreds := '0'
	tens := '0'
	units := '0'

	for hundreds <= '7' {
		tens = hundreds + 1
		for tens <= '8' {
			units = tens + 1
			for units <= '9' {
				ft.PrintRune(hundreds)
				ft.PrintRune(tens)
				ft.PrintRune(units)
				if (hundreds != '7') {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
				units++
			}
			tens++
		}
		hundreds++
	}
	ft.PrintRune('\n')
}
