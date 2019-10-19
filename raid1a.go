package student

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	if y > 1 {
		if x > 0 {
			z01.PrintRune(111)
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(45)
		}
		if x > 1 {
			z01.PrintRune(111)
		}
		if x > 0 {
			z01.PrintRune(10)
		}

	}

	if y > 1 {
		if x > 0 {
			for j := 1; j <= y-2; j++ {
				z01.PrintRune(124)
				for i := 1; i <= x-2; i++ {
					z01.PrintRune(32)
				}
				if x > 1 {
					z01.PrintRune(124)
				}
				z01.PrintRune(10)
			}
		}
	}

	if y >= 1 {
		if x > 0 {
			z01.PrintRune(111)
		}
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(45)
		}
		if x > 1 {
			z01.PrintRune(111)
		}
		if x > 0 {
			z01.PrintRune(10)
		}
	}
}
