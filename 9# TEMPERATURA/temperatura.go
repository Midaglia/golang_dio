package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	var tempF float64 = ebulicaoF
	var tempC float64 = (tempF - 32) * 5 / 9

	fmt.Println("A temperatura de ebulição da água em graus Farenheit =", tempF)
	fmt.Println("A temperatura de ebulição da água em graus Celsius =", tempC)

}
