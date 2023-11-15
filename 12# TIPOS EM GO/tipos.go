package main

import "fmt"

var x int
var y float64

func main() {

	x = 10
	y = 10.8
	fmt.Println("x =", x)
	fmt.Println("y =", y)
	fmt.Printf("X = %d,%T e y %g,%T", x, x, y, y)
}
