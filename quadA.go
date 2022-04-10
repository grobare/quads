package main

import "fmt"

func createFirstAndLastRow(w int) {
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

func createRow(w int) {
	//w is width of rectange as in characters used to do it
	for i := 0; i < w; i++ {
		if i == 0 || i == w-1 {
			fmt.Print("|")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Print("\n")
}

func quadA(w, h int) {
	// w - width of rectangle
	// h - height
	for i := 0; i < h; i++ {
		if i == 0 || i == h-1 {
			createFirstAndLastRow(w)
		} else {
			createRow(w)
		}
	}
}
