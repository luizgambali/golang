// reference: https://unidoc.io/post/write-pdfs-in-golang-beginners-guide/

package main

import (
	"fmt"

	"github.com/go-pdf/fpdf"
)

func main() {
	CriarPDF()
}

func CriarPDF() {
	fmt.Println("Criando PDF...")

	pdf := fpdf.New("P", "mm", "A4", "")

	//pdf.AddUTF8Font("Arial", "", "arial.ttf")

	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)

	pdf.Cell(40, 10, "Table title")

	// Table headers
	headers := []string{"Country", "Capital", "Population"}

	// Table data
	data := [][]string{
		[]string{"China", "Beijing", "1,403M"},
		[]string{"India", "New Delhi", "1,366M"},
		[]string{"United States", "Washington, D.C.", "331M"},
	}

	pdf.Ln(-1)
	pdf.SetFont("Arial", "B", 12)
	pdf.Cell(40, 10, "Main cities and their populations:")
	pdf.Ln(-1)

	pdf.SetFont("Arial", "B", 12)

	// Print table headers for
	for i, _ := range headers {
		pdf.CellFormat(40, 10, headers[i], "1", 0, "", false, 0, "")
	}

	pdf.Ln(-1)

	// Print table data
	pdf.SetFont("Arial", "", 12)
	for _, line := range data {
		for _, cell := range line {
			pdf.CellFormat(40, 10, cell, "1", 0, "", false, 0, "")
		}
		pdf.Ln(-1)
	}

	pdf.AddPage()
	pdf.Image("image.png", 10, 10, 100, 0, false, "", 0, "")

	err := pdf.OutputFileAndClose("hello.pdf")

	if err != nil {
		fmt.Println("Erro ao gerar pdf")
	}

}
