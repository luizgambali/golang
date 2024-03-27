// reference: https://xuri.medium.com/golang-library-for-reading-and-writing-microsoft-excel-xlsx-files-a4a63796c98a

package main

import (
	"fmt"

	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"

	"github.com/xuri/excelize/v2"
)

func main() {
	CreateExcelFile()
	ReadingExcelFile()
	AddChartToExcelFile()
	AddPictureToExcelFile()
}

func CreateExcelFile() {

	fmt.Println("Creating Excel file...")

	f := excelize.NewFile()

	index, err := f.NewSheet("Sheet2")

	if err != nil {
		fmt.Println(err)
	}

	f.SetCellValue("Sheet2", "A2", "Hello world.")
	f.SetCellValue("Sheet1", "B2", 100)
	f.SetActiveSheet(index)

	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func ReadingExcelFile() {

	fmt.Println("Reading data from Excel file...")

	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	// Get value from cell by given worksheet name and axis.
	cell, err := f.GetCellValue("Sheet1", "B2")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(cell)
	// Get all the rows in the Sheet1.
	rows, _ := f.GetRows("Sheet1")
	for _, row := range rows {
		for _, colCell := range row {
			fmt.Print(colCell, "\t")
		}
		fmt.Println()
	}
}

func AddChartToExcelFile() {
	categories := map[string]string{
		"A2": "Small", "A3": "Normal", "A4": "Large",
		"B1": "Apple", "C1": "Orange", "D1": "Pear"}
	values := map[string]int{"B2": 2, "C2": 3, "D2": 3, "B3": 5, "C3": 2, "D3": 4, "B4": 6, "C4": 7, "D4": 8}
	f := excelize.NewFile()

	for k, v := range categories {
		f.SetCellValue("Sheet1", k, v)
	}
	for k, v := range values {
		f.SetCellValue("Sheet1", k, v)
	}

	chart := excelize.Chart{
		Type: excelize.Col3DClustered,
		Series: []excelize.ChartSeries{
			{
				Name:       `Sheet1!$A$2`,
				Categories: `Sheet1!$B$1:$D$1`,
				Values:     `Sheet1!$B$2:$D$2`,
			},
			{
				Name:       `Sheet1!$A$3`,
				Categories: `Sheet1!$B$1:$D$1`,
				Values:     `Sheet1!$B$3:$D$3`,
			},
			{
				Name:       `Sheet1!$A$4`,
				Categories: `Sheet1!$B$1:$D$1`,
				Values:     `Sheet1!$B$4:$D$4`,
			},
		},
		Title: []excelize.RichTextRun{{Text: "Fruit 3D Clustered Column Chart"}},
	}

	if err := f.AddChart("Sheet1", "E1", &chart); err != nil {
		fmt.Println(err)
		return
	}
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func AddPictureToExcelFile() {
	f, err := excelize.OpenFile("Book1.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}

	printObject := true
	locked := false

	options := excelize.GraphicOptions{
		OffsetX:         15,
		OffsetY:         10,
		PrintObject:     &printObject,
		LockAspectRatio: false,
		Locked:          &locked,
	}

	// Insert a picture.
	if err := f.AddPicture("Sheet1", "M2", "image.png", &options); err != nil {
		fmt.Println(err)
	}

	// Save the spreadsheet with the origin path.
	if err = f.Save(); err != nil {
		fmt.Println(err)
	}
}
