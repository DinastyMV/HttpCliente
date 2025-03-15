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

	var input string
	fmt.Scanln(&input)

	switch input {
	case "1":
		mostrarImagemAleatória()
		return false
	case "2":
		fmt.Println("Digite a raça do cachorro:")
		var raca string
		var subRaca string
		fmt.Scanln(&raca, &subRaca)
		mostrarImagemAleatoriaPorRaca(raca, subRaca)
		return false
	case "sair":
		fmt.Println("Saindo...")
		return true
	default:
		fmt.Println("Entrada inválida! Tente novamente.")
		return false
	}
}
