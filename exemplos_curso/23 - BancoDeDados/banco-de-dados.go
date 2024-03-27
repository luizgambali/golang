package main

import (
	"database/sql"
	"fmt"
	"log"
)

func main() {
	stringConexao := "golang:golang@/devbook?charset=utf8&parseTime=True&loc=Local"

	db, erro := sql.Open("mysql", stringConexao)

	//o erro aqui só pega se houve alguma coisa que impedisse a conexão
	//se a senha está errada, ou algo assim, não vai cair aqui
	if erro != nil {
		log.Fatal(erro)
	}

	defer db.Close()

	fmt.Println("Verificando a conexão com o banco...")
	//aqui sim, é o teste se a conexão foi aberta ou não
	//não usei o erro := db.Ping(), pq estou reaproveitando o erro declarado lá em cima
	if erro = db.Ping(); erro != nil {
		log.Fatal(erro)
	}

	fmt.Println("Conectado ok ao banco!")

	linhas, erro := db.Query("select * from usuarios")

	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(linhas)

}
