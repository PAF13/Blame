package parts

import (
	"fmt"
	"log"

	"github.com/xuri/excelize/v2"
)

func readExcel(pfad string) *[][]string {
	file, err := excelize.OpenFile(pfad)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		if err := file.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}
	var maxLength int
	for _, b := range rows {
		if len(b) > maxLength {
			maxLength = len(b)
		}
	}
	fmt.Println(maxLength)
	var rowsClean [][]string
	for _, b := range rows {
		row := make([]string, maxLength)
		copy(row, b)
		rowsClean = append(rowsClean, row)
		//fmt.Println(b)
	}

	return &rowsClean
}
