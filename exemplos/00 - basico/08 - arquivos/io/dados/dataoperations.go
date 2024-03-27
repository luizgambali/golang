package dados

import (
	"os"
)

// Attempt to name the functions with first letter capital!

func LerDados() string {
	data, error := os.ReadFile("./data.txt")

	if error != nil {
		panic("Erro lendo arquivo")
	}

	return string(data)
}

func GravarDados(text string) {
	data := []byte(text + "\n")
	error := os.WriteFile("./newdata.txt", data, 0644)

	if error != nil {
		panic("Erro gravando dados no arquivo")
	}
}

func AdicionarDados(text string) {
	f, err := os.OpenFile("./newdata.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	defer f.Close()

	if err != nil {
		panic(err)
	}

	if _, err := f.Write([]byte(text + "\n")); err != nil {
		panic(err)
	}
}
