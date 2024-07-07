// reference: https://www.practical-go-lessons.com/post/how-to-generate-a-barcode-with-golang-ccb1i6auai2c70kp5ja0
package main

import (
	"image/png"
	"log"
	"os"

	"github.com/makiuchi-d/gozxing"
	"github.com/makiuchi-d/gozxing/oned"
)

func main() {

	toEncode := "7891234561234"

	writer := oned.NewCode128Writer()

	img, err := writer.Encode(toEncode, gozxing.BarcodeFormat_CODE_128, 250, 50, nil)

	if err != nil {
		log.Fatalf("Não foi possível gerar o código de barras: %s", err)
	}

	file, err := os.Create("barcode.png")
	if err != nil {
		log.Fatalf("Não foi possível riar o arquivo: %s", err)
	}

	defer file.Close()

	err = png.Encode(file, img)
	if err != nil {
		log.Fatalf("Não foi possível codificar o código de barras: %s", err)
	}

}
