package main

import "fmt"

//declaring global variables
var TopLeftCorner string
var TopRightCorner string
var DownLeftCorner string
var DownRightCorner string
var UpFill string
var DownFill string
var LeftFill string
var RightFill string
var Fill string = " " //default filling character

//assigning value to global variables
func Assign(topLeftCorner, topRightCorner, downLeftCorner, downRightCorner, upFill, downFill, leftFill, rightFill string) {
	TopLeftCorner = topLeftCorner
	TopRightCorner = topRightCorner
	DownLeftCorner = downLeftCorner
	DownRightCorner = downRightCorner
	UpFill = upFill
	DownFill = downFill
	LeftFill = leftFill
	RightFill = rightFill
}

//printing functions
func CreateFirstRow(width int) { //rectangle width is number of characters in a line
	for i := 0; i < width; i++ {
		if i == 0 {
			fmt.Print(TopLeftCorner)
		} else if i == width-1 {
			fmt.Print(TopRightCorner)
		} else {
			fmt.Print(UpFill)
		}
	}
	fmt.Println()
}

func CreateLastRow(width int) {
	for i := 0; i < width; i++ {
		if i == 0 {
			fmt.Print(DownLeftCorner)
		} else if i == width-1 {
			fmt.Print(DownRightCorner)
		} else {
			fmt.Print(DownFill)
		}
	}
	fmt.Println()
}

func CreateRow(width int) { //prints mid rows
	for i := 0; i < width; i++ {
		if i == 0 {
			fmt.Print(LeftFill)
		} else if i == width-1 {
			fmt.Print(RightFill)
		} else {
			fmt.Print(Fill)
		}
	}
	fmt.Println()
}

func Draw(width, height int) {
	if width <= 0 || height <= 0 { //return nothing if arguments are not valid
		return
	} else {
		for i := 0; i < height; i++ { //height defines how many lines we need to print
			if i == 0 {
				CreateFirstRow(width)
			} else if i == height-1 {
				CreateLastRow(width)
			} else {
				CreateRow(width)
			}
		}
	}
}

//calling functions
func QuadA(x, y int) {
	Assign("o", "o", "o", "o", "-", "-", "|", "|") //topLeftCorner, topRightCorner, downLeftCorner, downRightCorner, upFill, downFill, leftFill, rightFill
	Draw(x, y)
}

func QuadB(x, y int) {
	Assign("/", "\\", "\\", "/", "*", "*", "*", "*")
	Draw(x, y)
}

func QuadC(x, y int) {
	Assign("A", "A", "C", "C", "B", "B", "B", "B")
	Draw(x, y)
}

func QuadD(x, y int) {
	Assign("A", "C", "A", "C", "B", "B", "B", "B")
	Draw(x, y)
}
func QuadE(x, y int) {
	Assign("A", "C", "C", "A", "B", "B", "B", "B")
	Draw(x, y)
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
