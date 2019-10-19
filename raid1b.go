package student

import (
	"github.com/01-edu/z01"
)

func Raid1b(x, y int) {
	if y > 1 {
		if x >= 0 {
			z01.PrintRune(rune('/'))
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(rune('*'))
		}
		if x > 1 {
			z01.PrintRune(rune('\\'))
		}
		if x > 0 {
			z01.PrintRune(10)
		}

	}

	if y > 1 {
		if x > 0 {
			for j := 1; j <= y-2; j++ {
				z01.PrintRune(rune('*'))
				for i := 1; i <= x-2; i++ {
					z01.PrintRune(32)
				}
				if x > 1 {
					z01.PrintRune(rune('*'))
				}
				z01.PrintRune(10)
			}
		}
	}

	if y > 1 {
		if x > 0 {
			z01.PrintRune(rune('\\'))
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(rune('*'))
		}
		if x > 1 {
			z01.PrintRune(rune('/'))
		}
		if x > 0 {
			z01.PrintRune(10)
		}
	}
}
