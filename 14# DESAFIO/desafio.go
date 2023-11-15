//C = K - 273

package main

import "fmt"

func main() {

	var kelvin float32

	fmt.Println(("Digite um valor é Kelvin: "))
	fmt.Scan(&kelvin)

	var celsius float32 = kelvin - 273

	fmt.Println("O valor de kelvin convertido em graus celsius é igual:", celsius, "°C")

}
