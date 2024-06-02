package parts

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func writeStueckliste(pfad string, lagerbestand map[string]*BETRIEBSMITELL) {
	file2 := excelize.NewFile()
	defer func() {
		if err := file2.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rowNum := 1
	colNum := 0
	fmt.Println("starting excel")
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "==")
	file2.SetColWidth("Sheet1", "A", "A", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "=")
	file2.SetColWidth("Sheet1", "B", "B", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "++")
	file2.SetColWidth("Sheet1", "C", "C", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "+")
	file2.SetColWidth("Sheet1", "D", "D", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "-")
	file2.SetColWidth("Sheet1", "E", "E", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellnummer")
	file2.SetColWidth("Sheet1", "F", "F", 30)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "ERP")
	file2.SetColWidth("Sheet1", "G", "G", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "ERP KNT")
	file2.SetColWidth("Sheet1", "H", "H", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Hersteller")
	file2.SetColWidth("Sheet1", "I", "I", 10)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Menge")
	file2.SetColWidth("Sheet1", "J", "J", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung Moeller")
	file2.SetColWidth("Sheet1", "K", "K", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung KNT")
	file2.SetColWidth("Sheet1", "L", "L", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung Siteca")
	file2.SetColWidth("Sheet1", "M", "M", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Beisteller")
	file2.SetColWidth("Sheet1", "N", "N", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Quelle")
	file2.SetColWidth("Sheet1", "O", "O", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Beschreibung")
	file2.SetColWidth("Sheet1", "P", "P", 70)
	rowNum++
	for _, b := range lagerbestand {
		for _, bb := range b.Artikel {
			colNum = 0
			lineWriter(file2, "Sheet1", &colNum, &rowNum, b.BMK.FunktionaleZuordnung)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, b.BMK.Funktionskennzeichen)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, b.BMK.Aufstellungsort)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, b.BMK.Ortskennzeichen)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, b.BMK.BMK)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.Bestellnummer)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.ERP)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.ERP_KNT)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.Hersteller)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", bb.Stueckzahl))
			lineWriter(file2, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", bb.Bestellung_Moeller))
			lineWriter(file2, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", bb.Bestellung_KNT))
			lineWriter(file2, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", bb.Bestellung_Siteca))
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.Beistellung)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.Quelle)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.Beschreibung)
			rowNum++
		}
	}
	disable := false
	file2.AddTable("Sheet1", &excelize.Table{
		Range:             "A1:P" + fmt.Sprintf("%d", rowNum),
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowRowStripes:    &disable,
		ShowColumnStripes: true,
	})

	if err := file2.SaveAs(pfad + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val string) {

	file.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum), val)
	*colNum++
}
