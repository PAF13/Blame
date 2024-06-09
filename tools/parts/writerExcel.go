package parts

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func writeStueckliste(pfad string, lagerbestand map[string]*BETRIEBSMITELL) {
	file2 := excelize.NewFile()
	defer func() {
		if err := file2.Close(); err != nil {
			fmt.Println("--1--")
			fmt.Println(err)
			fmt.Println("--1--")
		}
	}()
	style, err := file2.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			TextRotation: 90,
			Horizontal:   "left",
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
			Size:  14,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	style2, err := file2.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
			Size:  14,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	tab1 := "Sheet1"
	file2.SetCellStyle(tab1, "A1", "R1", style2)
	rowNum := 1
	colNum := 0
	fmt.Println("starting excel")

	var column int
	//var col string
	fmt.Println(excelMakeColumn(&column, true))

	file2.SetRowHeight(tab1, 1, 150)
	lineWriter(file2, tab1, &colNum, &rowNum, "==")
	file2.SetColWidth(tab1, "A", "A", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "=")
	file2.SetColWidth(tab1, "B", "B", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "++")
	file2.SetColWidth(tab1, "C", "C", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "+")
	file2.SetColWidth(tab1, "D", "D", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "-")
	file2.SetColWidth(tab1, "E", "E", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "Bestellnummer")
	file2.SetColWidth(tab1, "F", "F", 30)
	lineWriter(file2, tab1, &colNum, &rowNum, "ERP")
	file2.SetColWidth(tab1, "G", "G", 13)
	lineWriter(file2, tab1, &colNum, &rowNum, "ERP KNT")
	file2.SetColWidth(tab1, "H", "H", 20)
	lineWriter(file2, tab1, &colNum, &rowNum, "Hersteller")
	file2.SetColWidth(tab1, "I", "I", 20)
	lineWriter(file2, tab1, &colNum, &rowNum, "Menge")
	file2.SetColWidth(tab1, "J", "J", 8)
	file2.SetCellStyle(tab1, "J1", "J1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Beistellung Kunde")
	file2.SetColWidth(tab1, "K", "K", 8)
	file2.SetCellStyle(tab1, "K1", "K1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Bestellung Moeller")
	file2.SetColWidth(tab1, "L", "L", 8)
	file2.SetCellStyle(tab1, "L1", "L1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Lager Siteca")
	file2.SetColWidth(tab1, "M", "M", 8)
	file2.SetCellStyle(tab1, "M1", "M1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Bestellung KNT")
	file2.SetColWidth(tab1, "N", "N", 8)
	file2.SetCellStyle(tab1, "N1", "N1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Bestellung Siteca")
	file2.SetColWidth(tab1, "O", "O", 8)
	file2.SetCellStyle(tab1, "O1", "O1", style)
	lineWriter(file2, tab1, &colNum, &rowNum, "Beisteller")
	file2.SetColWidth(tab1, "P", "P", 20)
	lineWriter(file2, tab1, &colNum, &rowNum, "Quelle")
	file2.SetColWidth(tab1, "Q", "Q", 20)
	lineWriter(file2, tab1, &colNum, &rowNum, "Beschreibung")
	file2.SetColWidth(tab1, "R", "R", 70)
	rowNum++
	var summe int
	for _, b := range lagerbestand {
		for _, bb := range b.Artikel {
			colNum = 0
			lineWriter(file2, tab1, &colNum, &rowNum, b.BMK.FunktionaleZuordnung)
			lineWriter(file2, tab1, &colNum, &rowNum, b.BMK.Funktionskennzeichen)
			lineWriter(file2, tab1, &colNum, &rowNum, b.BMK.Aufstellungsort)
			lineWriter(file2, tab1, &colNum, &rowNum, b.BMK.Ortskennzeichen)
			lineWriter(file2, tab1, &colNum, &rowNum, b.BMK.BMK)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Bestellnummer)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.ERP)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.ERP_KNT)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Hersteller)
			summe = colNum
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Stueckzahl)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Beistellung_Stueckzahl)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Bestellung_Moeller)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Lager_Siteca)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Bestellung_KNT)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Bestellung_Siteca)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Beistellung)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Quelle)
			lineWriter(file2, tab1, &colNum, &rowNum, bb.Beschreibung)
			rowNum++
		}
	}
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, summe, 2, summe, rowNum-1)
	formulaWriter(file2, tab1, &summe, &rowNum, 10, rowNum, summe-1, rowNum)

	disable := false
	file2.AddTable(tab1, &excelize.Table{
		Range:             "A1:R" + fmt.Sprintf("%d", rowNum),
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowRowStripes:    &disable,
		ShowColumnStripes: true,
	})

	if err := file2.SaveAs(pfad + ".xlsx"); err != nil {
		fmt.Println("--2--")
		fmt.Println(err)
		fmt.Println("--2--")
	}
}
func formulaWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val1 int, val2 int, val3 int, val4 int) {
	//formulaType, ref := excelize.STCellFormulaTypeNormal, fmt.Sprintf("%s%d:%s%d", string(rune(65+*colNum)), *rowNum, string(rune(65+*colNum)), *rowNum)

	formula := fmt.Sprintf("=SUMME(%s:%s)", fmt.Sprintf("%s%d", string(rune(65+val1)), val2), fmt.Sprintf("%s%d", string(rune(65+val3)), val4))
	fmt.Println(formula)
	if err := file.SetCellValue(sheet, excelMakeCell(colNum, rowNum), formula); err != nil {
		fmt.Println(err)
		return
	}
	*colNum++
}
func excelMakeColumn(colNum *int, add bool) string {
	col := string(rune(65 + *colNum))
	if add {
		*colNum++
	}
	return col
}
func excelMakeCell(colNum *int, rowNum *int) string {
	return fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum)
}
func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val interface{}) {

	file.SetCellValue(sheet, excelMakeCell(colNum, rowNum), val)
	*colNum++
}

func setHeaderRow(file *excelize.File, sheet string, colNum *int, rowNum *int, val string, width float64) {
	file.SetCellValue(sheet, excelMakeCell(colNum, rowNum), val)
	file.SetColWidth(sheet, string(rune(65+*colNum)), string(rune(65+*colNum)), width)
	*colNum++
}

func (Liste *LAGERLISTE) writeStueckliste(name string) {
	file2 := excelize.NewFile()
	defer func() {
		if err := file2.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	rowNum := 2
	colNum := 0
	tab1 := "Sheet1"

	enable, disable := true, false
	if err := file2.AddPicture(tab1, "A1", "\\\\ME-Datenbank-1\\Database\\Vorlagen SITECA\\SITECA Logo\\Logo Farbe\\Siteca Logo_CMYK.jpg",
		&excelize.GraphicOptions{
			PrintObject:     &enable,
			LockAspectRatio: false,
			OffsetX:         15,
			OffsetY:         10,
			Locked:          &disable,
		}); err != nil {
		fmt.Println(err)
	}
	style, err := file2.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			TextRotation: 90,
			Horizontal:   "left",
		},
		Font: &excelize.Font{
			Bold:  true,
			Color: "FFFFFF",
			Size:  14,
		},
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("starting excel")
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Bestellnummer", 30)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "ERP", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "ERP KNT", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "EPlan", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Hersteller", 30)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Preis Siteca", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Preis KNT", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Lager Moeller", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Lager Siteca", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Lager KNT", 20)
	file2.SetCellStyle(tab1, excelMakeCell(&colNum, &rowNum), excelMakeCell(&colNum, &rowNum), style)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Source Moeller", 8)
	file2.SetCellStyle(tab1, excelMakeCell(&colNum, &rowNum), excelMakeCell(&colNum, &rowNum), style)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Source Siteca", 8)
	file2.SetCellStyle(tab1, excelMakeCell(&colNum, &rowNum), excelMakeCell(&colNum, &rowNum), style)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Source KNT", 8)
	file2.SetCellStyle(tab1, excelMakeCell(&colNum, &rowNum), excelMakeCell(&colNum, &rowNum), style)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Source Eplan", 8)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "Beschreibung", 70)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "++", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "+", 20)
	setHeaderRow(file2, tab1, &colNum, &rowNum, "-", 20)
	rowNum++
	for _, b := range Liste.Artikel {

		colNum = 0
		lineWriter(file2, tab1, &colNum, &rowNum, b.Bestellnummer)
		lineWriter(file2, tab1, &colNum, &rowNum, b.ERP)
		lineWriter(file2, tab1, &colNum, &rowNum, b.ERP_KNT)
		lineWriter(file2, tab1, &colNum, &rowNum, b.ArtikelnummerEplan)
		lineWriter(file2, tab1, &colNum, &rowNum, b.Hersteller)
		lineWriter(file2, tab1, &colNum, &rowNum, b.EK_Siteca)
		lineWriter(file2, tab1, &colNum, &rowNum, b.EK_KNT)
		lineWriter(file2, tab1, &colNum, &rowNum, b.Bestellung_Moeller)
		lineWriter(file2, tab1, &colNum, &rowNum, b.Lager_Siteca)
		lineWriter(file2, tab1, &colNum, &rowNum, b.Bestellung_KNT)
		lineWriter(file2, tab1, &colNum, &rowNum, booltoString(b.DataSource.Moeller))
		lineWriter(file2, tab1, &colNum, &rowNum, booltoString(b.DataSource.Siteca))
		lineWriter(file2, tab1, &colNum, &rowNum, booltoString(b.DataSource.KNT))
		lineWriter(file2, tab1, &colNum, &rowNum, booltoString(b.DataSource.Eplan))
		lineWriter(file2, tab1, &colNum, &rowNum, b.Beschreibung)
		lineWriter(file2, tab1, &colNum, &rowNum, "")
		lineWriter(file2, tab1, &colNum, &rowNum, "")
		lineWriter(file2, tab1, &colNum, &rowNum, "")
		rowNum++

	}
	file2.AddTable(tab1, &excelize.Table{
		Range:             "B1:R" + fmt.Sprintf("%d", rowNum),
		Name:              "table",
		StyleName:         "TableStyleMedium1",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowRowStripes:    &disable,
		ShowColumnStripes: true,
	})

	if err := file2.SaveAs("\\\\ME-Datenbank-1\\Database\\Software\\Blame\\Data\\Lager\\" + name + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}

func booltoString(test bool) string {
	var value string
	if test {
		value = "X"
	}
	return value
}
