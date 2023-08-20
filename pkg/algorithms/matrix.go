package main

import "fmt"

//filling a multidimensional array

func main() {
	matrix := [2][3][3]int{}
	for _, h := range matrix {
		for _, cell := range h {
			cell = [3]int{1, 2, 3}
			fmt.Print(cell, " ")
		}
		fmt.Println()
	}
}
