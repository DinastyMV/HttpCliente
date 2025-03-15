package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func mostrarImagemAleatória() {
	const url string = "https://dog.ceo/api/breeds/image/random"
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("Erro ao buscar imagem: ", err)
		return
	}
	defer response.Body.Close()

	var imagem DogResponse

	if err := json.NewDecoder(response.Body).Decode(&imagem); err != nil {
		fmt.Println("Erro ao decodificar imagem: ", err)
		return
	}

	fmt.Printf("Messagem: %s\nStatus: %s\n", imagem.Message, imagem.Status)

}

func mostrarImagemAleatoriaPorRaca(raca string, subRaca string) {
	const url string = "https://dog.ceo/api/breed/"
	urlCompleta := url + raca + "/" + subRaca + "/images/random"
	if subRaca == "" {
		urlCompleta = url + raca + "/images/random"
	}
	response, err := http.Get(urlCompleta)
	if err != nil {
		fmt.Println("Erro ao buscar imagem por raça: ", err)
		return
	}
	defer response.Body.Close()

	var imagem DogResponse

	if err := json.NewDecoder(response.Body).Decode(&imagem); err != nil {
		fmt.Println("Erro decodificar imagem por raça: ", err)
		return
	}

	fmt.Printf("Message: %s\nStatus: %s", imagem.Message, imagem.Status)
}
