package piscine

import "fmt"

func CreateFirstAndLastRow(w int) {
	//w is width of rectangle
	for i := 0; i < w; i++ {
		if i == 0 || i == w-1 {
			fmt.Print("o")
		} else {
			fmt.Print("-")
		}
	}
	fmt.Print("\n")
}

func CreateRow(w int) {
	//w is width of rect
	for i := 0; i < w; i++ {
		if i == 0 || i == w-1 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func QuadA(w, h int) {
	// w - width of rectangle
	// h - height
	for i := 0; i < h; i++ {
		if i == 0 || i == h-1 {
			CreateFirstAndLastRow(w)
		} else {
			CreateRow(w)
		}
	}
}
