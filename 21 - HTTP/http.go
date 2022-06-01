package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Hello World</h1>")) //precisa passar a mensagem como byte
}

func usuarios(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Página de Usuários")) //precisa passar a mensagem como byte
}

func main() {
	// HTTP é um protocolo de comunicação - base da comunicação web

	// Cliente (Faz requisição) - Servidor (Processa Requisição e envia resposta)

	//Rotas
	//URI  - Identificação de um recurso
	//Método - GET, POST, PUT, DELETE

	http.HandleFunc("/home", home)
	http.HandleFunc("/usuarios", usuarios)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
