package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Usuarios struct {
	Login  string `json:"login"`
	Nome   string `json:"email"`
	Perfil string `json:"profile"`
}

//resposta simples
func Home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá Mundo!"))
}

//slice convertido para lista json
func ListarUsuarios(w http.ResponseWriter, r *http.Request) {
	lista := []Usuarios{{"luizgambali@yahoo.com.br", "Luiz Gambali", "administrador"},
		{"cleogambali@yahoo.com.br", "Cleo Gambali", "usuario"},
	}

	retornoJSon, erro := json.Marshal(lista)

	if erro != nil {
		w.Write([]byte("Ocorreu um erro ao processar a lista de usuarios"))
	}

	w.Write([]byte(retornoJSon))
}

func main() {

	//define a rota, e a função que atende por essa rota. No caso abaixo, uma função anonima

	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Página raiz!"))
	})

	//chama a função declarada
	http.HandleFunc("/home", Home)

	//chama uma função anonima
	http.HandleFunc("/usuarios", ListarUsuarios)

	//inicia a aplicação em http://localhost:5000/
	log.Fatal(http.ListenAndServe(":5000", nil))
}
