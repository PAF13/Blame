package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

// stücklistevergleich
func (a *App) ExcelChoice1(file1 string, file2 string) {
	fmt.Println("Stueckliste Compare: Recieving Stücklisten")
	CompareStueckliste(LoadStueckliste(file1), LoadStueckliste(file2))
	fmt.Println("Stueckliste Compare: Stücklisten recieved")
}

func (a *App) ImportStueckliste(file1 string, file2 string) (map[string][]string, map[string][]string) {

	return LoadStueckliste(file1), LoadStueckliste(file2)

}

func LoadStueckliste(x string) map[string][]string {
	fmt.Println("Stueckliste Compare: Creating map for:", x)
	//creating map of list
	stuecklisteMap := make(map[string][]string)
	headSkip := 1
	skip := 0
	//opening spreadsheet
	spreadsheet, err := excelize.OpenFile(x)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := spreadsheet.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	fmt.Println(spreadsheet.GetSheetList()[0])
	//Selecting rows from first sheet
	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}
	//reading rows
	for line, row := range rows {
		//cleaning empty erp numbers
		if row[6] == "" {
			t := strconv.Itoa(line)
			row[6] = "Empty" + t
		}
		//loading map with data as string
		if skip > headSkip {
			stuecklisteMap[row[6]] = row
			//fmt.Println(stuecklisteMap[row[6]])
		}
		skip++
	}
	return stuecklisteMap
}
func CompareStueckliste(old map[string][]string, new map[string][]string) {
	file := excelize.NewFile()
	headers := []string{
		"Artikelnummer",
		"»»» Stücklisten/Sets «««",
		"ist Stückliste",
		"Stücklistenart",
		"Positionen ausblenden",
		"SL-Pos.Rang",
		"SL-Pos.Nummer",
		"SL-Pos.Menge",
		"Löschen",
	}
	headers2 := []string{
		"Stücklisten-Kopfartikel, dieser muß schon in T8 angelegt sein.",
		"»»» Stücklisten/Sets «««",
		"1=JA",
		"0=Kopf ohne Positionen, 1=Pos. ohne Preise, 2=Pos mit Preisen",
		"A, B, L ,R",
		"GANZ WICHTIG: durchgehend nummerieren, sonst werden keine neuen Positionen angefügt",
		"Artikelnummer des zugehörigen Artikels",
		"Stücklisten-menge",
		"Hersteller",
		"Typnummer",
		"Artikelnummer",
		"Artikel: Bezeichnung",
	}
	for i, header := range headers {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for i, header := range headers2 {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 2), header)
	}
	line := 3
	for newValue := range new {
		_, Match := old[newValue]
		if Match && new[newValue][7] != old[newValue][7] {
			for i := 6; i < len(new[newValue]); i++ {
				if i == 7 {
					mengeOld, _ := strconv.Atoi(old[newValue][i])
					mengeNew, _ := strconv.Atoi(new[newValue][i])
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew-mengeOld)
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), new[newValue][i])
				}
			}
			line++
		} else if !Match {
			for i := 6; i < len(new[newValue]); i++ {
				if i == 7 {
					mengeNew, _ := strconv.Atoi(new[newValue][i])
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew)
				} else {
					file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), new[newValue][i])
				}
			}
			delete(new, newValue)
			line++
		}
		delete(old, newValue)

	}
	for newValue := range old {
		for i := 6; i < len(old[newValue]); i++ {
			if i == 7 {
				mengeNew, _ := strconv.Atoi(old[newValue][i])
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), mengeNew*-1)
			} else {
				file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), line), old[newValue][i])
			}

		}
		line++
		delete(old, newValue)
	}

	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Stueckliste.xlsx"); err != nil {
		fmt.Println(err)
	}
}
