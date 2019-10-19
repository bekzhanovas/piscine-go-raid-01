package main

import (
	"github.com/01-edu/z01"
)

func Raid1a(x, y int) {
	z01.PrintRune(111)
	for i := 1; i <= x-2; i++ {
		z01.PrintRune(45)
	}
	z01.PrintRune(111)
	z01.PrintRune(10)
	for j := 1; j <= y-2; j++ {
		z01.PrintRune(124)
		for i := 1; i <= x-2; i++ {
			z01.PrintRune(32)
		}
		z01.PrintRune(124)
		z01.PrintRune(10)
	}
	z01.PrintRune(111)
	for i := 1; i <= x-2; i++ {
		z01.PrintRune(45)
	}
	z01.PrintRune(111)
	z01.PrintRune(10)
  }
  
