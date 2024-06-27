package	piscine

import "ft"

func PrintComb2() {
	left_tens := '0'
	left_units := '0'
	right_tens := '0'
	right_units := left_units + 1

	for left_tens <= '9' {
		for left_units <= '9' {
			for right_tens <= '9' {
				for right_units <= '9' {
					ft.PrintRune(left_tens)
					ft.PrintRune(left_units)
					ft.PrintRune(' ')
					ft.PrintRune(right_tens)
					ft.PrintRune(right_units)
					if (left_tens != '9' || left_units != '8') {
						ft.PrintRune(',')
						ft.PrintRune(' ')
					}
					right_units++
				}
				right_tens++
				right_units = '0'
			}
			left_units++
			right_tens = left_tens
			right_units = left_units + 1
		}
		left_tens++
		left_units = '0'
		right_tens = left_tens
		right_units = left_units + 1
	}
	ft.PrintRune('\n')
}
