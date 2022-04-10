package piscine2

import (
	"fmt"
	"strings"
)

var TopLeftCorner string
var TopRightCorner string
var DownLeftCorner string
var DownRightCorner string
var UpFill string
var DownFill string
var LeftDoor string
var RightDoor string

func assign(a, b, c, d, e, f, g, h string) {
	TopLeftCorner = a
	TopRightCorner = b
	DownRightCorner = c
	DownLeftCorner = d
	UpFill = e
	DownFill = f
	LeftDoor = g
	RightDoor = h
	// a,    b,    c,  d,    e ,  f,   g ,	h
}

func QuadA(x, y int) {
	assign("o", "o", "o", "o", "-", "-", "|", "|")
	Draw(x, y)
}

func QuadB(x, y int) {
	assign("/", "\\", "\\", "/", "*", "*", "*", "*")
	Draw(x, y)
}

func QuadC(x, y int) {
	assign("A", "A", "C", "C", "B", "B", "B", "B")
	Draw(x, y)
}

func QuadD(x, y int) {
	assign("A", "C", "A", "C", "B", "B", "B", "B")
	Draw(x, y)
}
func QuadE(x, y int) {
	assign("A", "C", "C", "A", "B", "B", "B", "B")
	Draw(x, y)
}

func Draw(x, y int) {
	if x == 0 || y == 0 {
		return
	}

	if x == 1 && y == 1 {
		fmt.Println(TopLeftCorner)
		return
	}

	if x == 1 && y == 2 {
		fmt.Println(TopLeftCorner)
		fmt.Println(DownLeftCorner)
		return
	}

	if x == 2 && y == 2 {
		fmt.Printf("%v%v\n", TopLeftCorner, TopRightCorner)
		fmt.Printf("%v%v\n", DownLeftCorner, DownRightCorner)
		return
	}

	if y == 1 {
		fmt.Printf("%v%v%v\n", TopLeftCorner, Repeat(UpFill, x-2), TopRightCorner)
		return
	}

	if x == 1 {
		fmt.Printf("%v\n", TopLeftCorner)
		for i := 1; i < y; i++ {
			fmt.Printf("%v\n", LeftDoor)
		}
		fmt.Printf("%v\n", DownLeftCorner)
		return
	}

	fmt.Printf("%v%v%v\n", TopLeftCorner, Repeat(UpFill, x-2), TopRightCorner)
	for i := 2; i < y; i++ {
		fmt.Printf("%v%v%v\n", LeftDoor, Repeat(" ", x-2), RightDoor)
	}
	fmt.Printf("%v%v%v\n", DownLeftCorner, Repeat(DownFill, x-2), DownRightCorner)

}

func Repeat(s string, n int) string {
	if n > 0 {
		return strings.Repeat(s, n)
	}
	return ""
}

func main() {
	QuadA(5, 3)
	fmt.Println()
	QuadB(5, 1)
	fmt.Println()
	QuadC(1, 1)
	fmt.Println()
	QuadD(1, 5)
	fmt.Println()
}
