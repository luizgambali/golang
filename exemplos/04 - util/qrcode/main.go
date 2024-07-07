package main

import (
	"fmt"

	"github.com/skip2/go-qrcode"
)

func main() {

	err := qrcode.WriteFile("https://github.com/luizgambali/golang", qrcode.Medium, 256, "qr.png")

	if err != nil {
		fmt.Printf("Erro ao gerar o codigo QR!")
	}
}
