package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const urlAPI = "https://dog.ceo/api/breeds/image/random"

func mostrarImagem() {
	response, err := http.Get(urlAPI)
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

func main() {
	mostrarImagem()
}
