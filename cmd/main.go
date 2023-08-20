package main

import "fmt"

func main() {
	type a int
	var b int

	c := a(5)
	fmt.Println(c, b)
}
