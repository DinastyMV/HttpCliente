package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Primeiro projeto em Go")
	SayHello()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!\n")
	})

	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem vindo a página de users!")
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error: ", err)
	}
}
