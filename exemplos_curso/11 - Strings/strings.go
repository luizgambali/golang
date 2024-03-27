package main

import (
	"fmt"
	"strings"
)

var p = fmt.Println

func main() {

	var texto string = "Esta é a string de teste, para manipulação de funções de string"

	fmt.Println("Manipulação de strings")

	p("Contains:  ", strings.Contains(texto, "teste"))
	p("Count:     ", strings.Count(texto, "t"))
	p("HasPrefix: ", strings.HasPrefix(texto, "te"))
	p("HasSuffix: ", strings.HasSuffix(texto, "st"))
	p("Index:     ", strings.Index(texto, "para"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace(texto, "e", "X", -1))
	p("Replace:   ", strings.Replace(texto, "e", "X", 1))
	p("Split:     ", strings.Split(texto, " "))
	p("ToLower:   ", strings.ToLower(texto))
	p("ToUpper:   ", strings.ToUpper(texto))
	p("Compare	teste Teste  ", strings.Compare("teste", "Teste"))
	p("Compare	Teste teste  ", strings.Compare("Teste", "teste"))
	p("Compare	teste teste  ", strings.Compare("teste", "teste"))
}
