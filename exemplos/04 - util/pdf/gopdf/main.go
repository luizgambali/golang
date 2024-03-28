/*

	Referencia: https://github.com/signintech/gopdf

	Outros exemplos: https://github.com/oneplus1000/gopdfsample

*/

package main

import (
	"log"

	"github.com/signintech/gopdf"
)

func main() {

	CriarPdfTextoSimples()
	CriarPdfComImagem()
	CriarPdfComHeader()

}

func CriarPdfTextoSimples() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	err := pdf.AddTTFFont("Arial", "./fonts/ARIALUNI.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.Cell(nil, "Teste de geração de PDF com a biblioteca gopdf.")
	pdf.WritePdf("hello.pdf")
}

func CriarPdfComImagem() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4})
	pdf.AddPage()
	var err error
	err = pdf.AddTTFFont("Arial", "./fonts/ARIALUNI.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.Image("./images/image.png", 200, 50, nil) //print image
	err = pdf.SetFont("Arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}
	pdf.SetXY(250, 200)                        //move current location
	pdf.Cell(nil, "exemplo de pdf com imagem") //print text

	pdf.WritePdf("image.pdf")
}

func CriarPdfComHeader() {
	pdf := gopdf.GoPdf{}
	pdf.Start(gopdf.Config{PageSize: *gopdf.PageSizeA4}) //595.28, 841.89 = A4

	err := pdf.AddTTFFont("Arial", "./fonts/ARIALUNI.ttf")
	if err != nil {
		log.Print(err.Error())
		return
	}

	err = pdf.SetFont("Arial", "", 14)
	if err != nil {
		log.Print(err.Error())
		return
	}

	pdf.AddHeader(func() {
		pdf.SetY(5)
		pdf.Cell(nil, "header")
	})
	pdf.AddFooter(func() {
		pdf.SetY(825)
		pdf.Cell(nil, "footer")
	})

	pdf.AddPage()
	pdf.SetY(400)
	pdf.Text("page 1 content")
	pdf.AddPage()
	pdf.SetY(400)
	pdf.Text("page 2 content")

	pdf.WritePdf("header-footer.pdf")
}
