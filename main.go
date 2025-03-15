package main

import (
	"fmt"
)

func main() {
	var sair bool = false

	for sair != true {
		sair = inicio()
	}
}

func inicio() bool {
	fmt.Println("Qual operação deseja realizar:")
	fmt.Println("1 - Imagem Aleatória")
	fmt.Println("2 - Imagem Aleatória por raça")
	fmt.Println("9 - Sair")

	var input int8
	fmt.Scanln(&input)

	switch input {
	case 1:
		mostrarImagemAleatória()
		return false
	case 2:
		fmt.Println("Digite a raça do cachorro:")
		var raca string
		fmt.Scanln(&raca)
		mostrarImagemAleatoriaPorRaca(raca)
		return false
	case 9:
		fmt.Println("Saindo...")
		return true
	default:
		fmt.Println("Entrada inválida! Tente novamente.")
		return false
	}
}
