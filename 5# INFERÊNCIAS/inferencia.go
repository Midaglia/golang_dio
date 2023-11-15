package main

import "fmt"

// Com tipagem a mostra.

func main() {

	var nome string = "Midaglia"
	var idade int = 22
	var versao float32 = 3.2
	fmt.Println("Meu nome é", nome, "e minha idade é", idade)
	fmt.Println("Estou usando o programa de versão:", versao)

}

// Ou pode ser feito sem tipagem a mostra.

func main2() {

	var nome = "Midaglia"
	var idade = 22
	var versao = 3.2
	fmt.Println("Meu nome é", nome, "e minha idade é", idade)
	fmt.Println("Estou usando o programa de versão:", versao)

}
