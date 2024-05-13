package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func ImportFile2(fileName string, kunde string, fileType string) {
	defer wg.Done()
	excelSize := 0

	file, err := excelize.OpenFile(rootPfadInput + fileName + ".xlsx")
	if err != nil {
		return
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

	for _, b := range rows {
		if len(b) > excelSize {
			excelSize = len(b)
		}
	}
	excelSize = excelSize + 1
	newRow := NewExcelImport(kunde, excelSize, fileType, fileName+".xlsx")
	for a, b := range rows {
		excelRow := make([]string, excelSize)
		newRow.Rows = append(newRow.Rows, excelRow)
		copy(newRow.Rows[a], b)
	}

	content, err := json.MarshalIndent(newRow, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfadOutput+"Blame_Import_"+fileName+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFile2(fileName string) {
	defer wg.Done()
	jsonFile, err := os.Open(rootPfadOutput + "Blame_Import_" + fileName + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var excelFile EXCEL_IMPORT
	json.Unmarshal([]byte(byteValue), &excelFile)

	artikelliste := ARTIKELLISTE{
		Header:  excelFile.Header,
		Artikel: map[string][]ARTIKEL{},
	}

	for _, newRow := range excelFile.Rows {
		bestellnummer := bestellnummerClean(newRow[excelFile.Columns.Bestellnummer])
		stueckzahl, _ := strconv.ParseFloat(newRow[excelFile.Columns.Stueckzahl], 64)
		if (bestellnummer != "" && excelFile.Header.Name == "lager") || excelFile.Header.Name != "lager" {
			artikelliste.Artikel[bestellnummer] = append(artikelliste.Artikel[bestellnummer], ARTIKEL{
				ERP:               newRow[excelFile.Columns.ERP],
				ERP_KNT:           newRow[excelFile.Columns.ERP_KNT],
				Bestellnummer:     bestellnummer,
				Stueckzahl:        stueckzahl,
				Bestellung_Siteca: stueckzahl,
				Quelle:            excelFile.Header.Source,
				Beschreibung:      newRow[excelFile.Columns.Beschreibung],
				Hersteller:        newRow[excelFile.Columns.Hersteller],
				Beistellung:       newRow[excelFile.Columns.Beistellung],
				Fehler:            []string{},
				BMK: BETRIEBSMITELLKENNZEICHEN{
					FunktionaleZuordnung: newRow[excelFile.Columns.FunktionaleZuordnung],
					Funktionskennzeichen: newRow[excelFile.Columns.Funktionskennzeichen],
					Aufstellungsort:      newRow[excelFile.Columns.Aufstellungsort],
					Ortskennzeichen:      newRow[excelFile.Columns.Ortskennzeichen],
					BMK:                  newRow[excelFile.Columns.BMK],
				},
			})
		}

	}
	content, err := json.MarshalIndent(artikelliste, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfadOutput+"Blame_Clean_"+fileName+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func sumListe(stueckliste string) {
	jsonFile_KNT, err := os.Open(rootPfadOutput + "Blame_Clean_Kopie von Lagerhueter_26_04_2024.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_KNT.Close()
	jsonFile_SITECA, err := os.Open(rootPfadOutput + "Blame_Clean_Topix_Artikel20240502.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_SITECA.Close()
	jsonFile_LISTE, err := os.Open(rootPfadOutput + "Blame_Clean_" + stueckliste + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()

	byteValue_KNT, _ := ioutil.ReadAll(jsonFile_KNT)
	byteValue_SITECA, _ := ioutil.ReadAll(jsonFile_SITECA)
	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)

	artikel_KNT := ARTIKELLISTE{}
	artikel_SITECA := ARTIKELLISTE{}
	artikel_LISTE := ARTIKELLISTE{}
	artikel_KNT_CLEAN := NewArtikelliste("Type dont know", "Source dont know")
	artikel_LISTE_CLEAN := NewArtikelliste("Type dont know", "Source dont know")

	json.Unmarshal(byteValue_KNT, &artikel_KNT)
	json.Unmarshal(byteValue_SITECA, &artikel_SITECA)
	json.Unmarshal(byteValue_LISTE, &artikel_LISTE)

	for _, b := range artikel_LISTE.Artikel {
		for _, bb := range b {
			_, ok := artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen]
			if !ok {
				artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen] = append(artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen], bb)
			} else {
				artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Stueckzahl = artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Stueckzahl + bb.Stueckzahl
				artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Bestellung_Siteca = artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Bestellung_Siteca + bb.Stueckzahl
			}
		}
	}

	for _, b := range artikel_KNT.Artikel {
		for _, bb := range b {
			_, ok := artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen]
			if !ok {
				artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen] = append(artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen], bb)
			} else {
				artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Stueckzahl = artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen][0].Stueckzahl + bb.Stueckzahl
			}
		}
	}

	for a, b := range artikel_LISTE_CLEAN.Artikel {
		_, ok := artikel_SITECA.Artikel[b[0].Bestellnummer]
		if ok {
			for aa, bb := range artikel_SITECA.Artikel[b[0].Bestellnummer] {

				if len(artikel_SITECA.Artikel[b[0].Bestellnummer]) >= 1 && len(artikel_SITECA.Artikel[b[0].Bestellnummer]) == aa+1 {
					artikel_LISTE_CLEAN.Artikel[a][0].ERP = artikel_LISTE_CLEAN.Artikel[a][0].ERP + bb.ERP
				} else if len(artikel_SITECA.Artikel[b[0].Bestellnummer]) > 1 {
					artikel_LISTE_CLEAN.Artikel[a][0].ERP = bb.ERP + " | "
				}

			}

		}
		if b[0].Beistellung == "SITECA" {
			_, ok2 := artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer]
			if ok2 {

				if artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca > artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl {
					//set KNT bestellung
					artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_KNT = artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl
					//set Rest stueckzahl
					artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca = artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca - artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl
					//set lager rest
					artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl = 0
				} else {
					//set KNT bestellung
					artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_KNT = artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca

					artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl = artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl - artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca
					//set lager rest
					artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca = 0
				}

			}
		}
	}

	content, err := json.MarshalIndent(artikel_LISTE_CLEAN, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfadOutput+"Blame_Clean2_"+stueckliste+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}

	writeStueckliste(artikel_LISTE_CLEAN.Artikel, "sdf", "sdf")
}

func writeStueckliste(lagerbestand map[string][]ARTIKEL, quelle string, fileType string) {
	file2 := excelize.NewFile()
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
		for _, bb := range b {
			colNum = 0
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.BMK.FunktionaleZuordnung)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.BMK.Funktionskennzeichen)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.BMK.Aufstellungsort)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, bb.BMK.Ortskennzeichen)
			lineWriter(file2, "Sheet1", &colNum, &rowNum, "")
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

	if err := file2.SaveAs(rootPfadOutput + "Blame_Clean2_" + fileType + "_" + quelle + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val string) {
	file.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum), val)
	*colNum++
}
