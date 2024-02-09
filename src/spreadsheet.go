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

func WriteData(f *excelize.File, titles []string, descriptions []string, prices []string, links []string) {
	f.SetCellValue("Sheet1", "A1", "Product")
	f.SetCellValue("Sheet1", "B1", "Description")
	f.SetCellValue("Sheet1", "C1", "Price")
	f.SetCellValue("Sheet1", "D1", "Link")

	for i := range titles {
		dataRow := i + 2

		titleColumn := "A"
		descriptionColumn := "B"
		priceColumn := "C"
		linkColumn := "D"

		titleValue := titles[i]
		descriptionValue := descriptions[i]
		priceValue := prices[i]
		linkValue := links[i]

		titleCell := fmt.Sprintf("%s%d", titleColumn, dataRow)
		descriptionCell := fmt.Sprintf("%s%d", descriptionColumn, dataRow)
		priceCell := fmt.Sprintf("%s%d", priceColumn, dataRow)
		linkCell := fmt.Sprintf("%s%d", linkColumn, dataRow)

		f.SetCellValue("Sheet1", titleCell, titleValue)
		f.SetCellValue("Sheet1", descriptionCell, descriptionValue)
		f.SetCellValue("Sheet1", priceCell, priceValue)
		f.SetCellValue("Sheet1", linkCell, linkValue)
	}

	SaveDocument(f)
}
