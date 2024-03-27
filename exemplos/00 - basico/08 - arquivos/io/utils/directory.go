package utils

import (
	"fmt"
	"io"
	"log"
	"os"
)

func CriarPasta(dirname string) {
	err := os.Mkdir(dirname, 0755)

	if err != nil {
		panic(err)
	}
}

func RemoverPasta(dirname string) {
	err := os.Remove(dirname)

	if err != nil {
		panic(err)
	}
}

func MostrarArquivos(dirname string) {
	files, err := os.ReadDir(dirname)

	if err != nil {
		panic(err)
	}

	for _, file := range files {
		fmt.Printf("\t%v - %v\n", file.Name(), file.IsDir())
	}
}

func ListarArquivos(dirname string) []string {
	files, err := os.ReadDir(dirname)

	if err != nil {
		panic(err)
	}

	var fileNames []string

	for _, file := range files {
		fileNames = append(fileNames, file.Name())
	}

	return fileNames
}

// Referencia: https://zetcode.com/golang/copyfile/
func CopiarArquivo(sourceFilePath string, destinationFilePath string) {

	fin, err := os.Open(sourceFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fin.Close()

	fout, err := os.Create(destinationFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer fout.Close()

	_, err = io.Copy(fout, fin)

	if err != nil {
		log.Fatal(err)
	}
}
