package main

import "fmt"

const ebulicaoF float64 = 212.0

func main() {

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9

	fmt.Printf("A temperatura de ebulição da água em graus Farenheit = %g e a temperatura de ebulição da água em graus Celsius = %g", tempF, tempC)

}
