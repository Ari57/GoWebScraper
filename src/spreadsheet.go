package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func CreateDocument() *excelize.File {
	f := excelize.NewFile()

	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	return f
}

func SaveDocument(f *excelize.File) {
	if err := f.SaveAs("../Products.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func WriteData(f *excelize.File) {
	f.SetCellValue("Sheet1", "A1", "Product")
	f.SetCellValue("Sheet1", "B1", "Price")
	f.SetCellValue("Sheet1", "C1", "Link")

	data := [][]interface{}{
		{1, "John", 30},
		{2, "Alex", 20},
		{3, "Bob", 40},
	}

	for i := range data {
		dataRow := i + 2
		column := "A"

		cellRef := fmt.Sprintf("%s%d", column, dataRow)
		f.SetCellValue("Sheet1", cellRef, "Ejbu")
	}

	SaveDocument(f)
}
