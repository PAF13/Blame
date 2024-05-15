package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func (a *App) ReturnOrte() []string {
	defer wg.Done()
	jsonFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Blame_Clean2_stueckliste.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue_jsonFile, _ := ioutil.ReadAll(jsonFile)

	artikel_jsonFile := ARTIKELLISTE{}

	json.Unmarshal(byteValue_jsonFile, &artikel_jsonFile)
	bmks := []string{}
	for _, b := range artikel_jsonFile.BMK_Liste {
		bmks = append(bmks, b)
	}
	fmt.Println(artikel_jsonFile.BMK_Liste)
	fmt.Println(bmks)
	return bmks
}
func ImportFile(pfad string, kunde string, fileType string, fileName string) {
	defer wg.Done()
	file, err := excelize.OpenFile(pfad)
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
	excelSize := 0
	for _, b := range rows {
		if len(b) > excelSize {
			excelSize = len(b)
		}
	}
	excelSize = excelSize + 1
	newRow := NewExcelImport(kunde, excelSize, fileType, fileName)
	for a, b := range rows {
		excelRow := make([]string, excelSize)
		newRow.Rows = append(newRow.Rows, excelRow)
		copy(newRow.Rows[a], b)
	}
	content, err := json.MarshalIndent(newRow, "", "\t")
	if err != nil {
		fmt.Println(err)
	}
	err = ioutil.WriteFile(rootPfadOutput+"Blame_Import1_"+kunde+"_"+fileName+"_"+fileType+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func loadFile(kunde string, fileType string, fileName string) {
	defer wg.Done()
	jsonFile, err := os.Open(rootPfadOutput + "Blame_Import1_" + kunde + "_" + fileName + "_" + fileType + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var excelFile EXCEL_IMPORT
	json.Unmarshal([]byte(byteValue), &excelFile)

	artikelliste := ARTIKELLISTE{
		Header:    excelFile.Header,
		BMK_Liste: map[string]string{},
		Artikel:   map[string][]ARTIKEL{},
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
					//BMKVollständig:       newRow[excelFile.Columns.Ortskennzeichen],
					BMKVollständig: newRow[excelFile.Columns.FunktionaleZuordnung] + newRow[excelFile.Columns.Funktionskennzeichen] + newRow[excelFile.Columns.Aufstellungsort] + newRow[excelFile.Columns.Ortskennzeichen],
				},
			})
		}

	}
	content, err := json.MarshalIndent(artikelliste, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfadOutput+"Blame_Import2_"+kunde+"_"+fileName+"_"+fileType+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func sumListe(kunde string, fileType string, fileName string) []string {
	defer wg.Done()
	jsonFile_KNT, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Blame_Import2_KNT_Lagerhueter_Lager.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_KNT.Close()

	jsonFile_MOELLER, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Blame_Import2_MOELLER_Moeller_Lager.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_KNT.Close()

	jsonFile_SITECA, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\Blame_Import2_SITECA_Topix_Lager.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_SITECA.Close()

	jsonFile_LISTE, err := os.Open(rootPfadOutput + "Blame_Import2_" + kunde + "_" + fileName + "_" + fileType + ".json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()

	byteValue_KNT, _ := ioutil.ReadAll(jsonFile_KNT)
	byteValue_MOELLER, _ := ioutil.ReadAll(jsonFile_MOELLER)
	byteValue_SITECA, _ := ioutil.ReadAll(jsonFile_SITECA)
	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)

	artikel_KNT := ARTIKELLISTE{}
	artikel_MOELLER := ARTIKELLISTE{}
	artikel_SITECA := ARTIKELLISTE{}
	artikel_LISTE := ARTIKELLISTE{}
	artikel_KNT_CLEAN := NewArtikelliste(fileName, "Source dont know")
	artikel_MOELLER_CLEAN := NewArtikelliste(fileName, "Source dont know")
	artikel_LISTE_CLEAN := NewArtikelliste(fileName, "Source dont know")

	json.Unmarshal(byteValue_KNT, &artikel_KNT)
	json.Unmarshal(byteValue_MOELLER, &artikel_MOELLER)
	json.Unmarshal(byteValue_SITECA, &artikel_SITECA)
	json.Unmarshal(byteValue_LISTE, &artikel_LISTE)
	fmt.Println("summing list")
	//sum stueckliste
	for _, b := range artikel_LISTE.Artikel {
		for _, bb := range b {
			if bb.Beistellung == "SITECA" {
				orten := bb.BMK.Ortskennzeichen
				if orten != "" {
					artikel_LISTE_CLEAN.BMK_Liste[orten] = orten
				}
				_, ok := artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten]
				if !ok {
					artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten] = append(artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten], bb)
				} else {
					artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl = artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl + bb.Stueckzahl
					artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten][0].Bestellung_Siteca = artikel_LISTE_CLEAN.Artikel[bb.Bestellnummer+orten][0].Bestellung_Siteca + bb.Stueckzahl

				}
			}
		}
	}
	fmt.Println("summing lager")
	//Sum Lager
	for _, b := range artikel_MOELLER.Artikel {
		for _, bb := range b {
			orten := bb.BMK.Ortskennzeichen
			_, ok := artikel_MOELLER_CLEAN.Artikel[bb.Bestellnummer+orten]
			if !ok {
				artikel_MOELLER_CLEAN.Artikel[bb.Bestellnummer+orten] = append(artikel_MOELLER_CLEAN.Artikel[bb.Bestellnummer+orten], bb)
			} else {
				artikel_MOELLER_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl = artikel_MOELLER_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl + bb.Stueckzahl
			}
		}
	}
	for _, b := range artikel_KNT.Artikel {
		for _, bb := range b {
			orten := bb.BMK.Ortskennzeichen
			_, ok := artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+orten]
			if !ok {
				artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+orten] = append(artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+orten], bb)
			} else {
				artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl = artikel_KNT_CLEAN.Artikel[bb.Bestellnummer+orten][0].Stueckzahl + bb.Stueckzahl
			}
		}
	}

	for a, b := range artikel_LISTE_CLEAN.Artikel {
		_, ok := artikel_SITECA.Artikel[b[0].Bestellnummer]
		if ok {
			fmt.Println("checking erp")
			for aa, bb := range artikel_SITECA.Artikel[b[0].Bestellnummer] {

				if len(artikel_SITECA.Artikel[b[0].Bestellnummer]) >= 1 && len(artikel_SITECA.Artikel[b[0].Bestellnummer]) == aa+1 {
					artikel_LISTE_CLEAN.Artikel[a][0].ERP = artikel_LISTE_CLEAN.Artikel[a][0].ERP + bb.ERP
				} else if len(artikel_SITECA.Artikel[b[0].Bestellnummer]) > 1 {
					artikel_LISTE_CLEAN.Artikel[a][0].ERP = bb.ERP + " | "
				}

			}

		}
		if true {
			fmt.Println("spliting lager")
			bestellungMoeller := &artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Moeller
			bestellungKNT := &artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_KNT
			bestellungSiteca := &artikel_LISTE_CLEAN.Artikel[a][0].Bestellung_Siteca
			var lagerMoeller *float64
			var lagerKNT *float64
			_, ok_Moeller := artikel_MOELLER_CLEAN.Artikel[b[0].Bestellnummer]
			if ok_Moeller {
				lagerMoeller = &artikel_MOELLER_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl
			}

			_, ok_KNT := artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer]
			if ok_KNT {
				lagerKNT = &artikel_KNT_CLEAN.Artikel[b[0].Bestellnummer][0].Stueckzahl
			}

			if *bestellungSiteca != 0 {

				if ok_Moeller && *bestellungSiteca > *lagerMoeller {
					*bestellungSiteca = *bestellungSiteca - *lagerMoeller
					*bestellungMoeller = *lagerMoeller
					*lagerMoeller = 0

				} else if ok_Moeller {
					*bestellungMoeller = *bestellungSiteca
					*bestellungSiteca = 0
					*lagerMoeller = 0
				}

			}

			if *bestellungSiteca != 0 {

				if ok_KNT && *bestellungSiteca > *lagerKNT {
					*bestellungSiteca = *bestellungSiteca - *lagerKNT
					*bestellungKNT = *lagerKNT
					*lagerKNT = 0
				} else if ok_KNT {
					*bestellungKNT = *bestellungSiteca
					*bestellungSiteca = 0
					*lagerKNT = 0
				}
			}

		}
	}
	temp := []string{}

	for a := range artikel_LISTE_CLEAN.BMK_Liste {

		temp = append(temp, a)
	}

	content, err := json.MarshalIndent(artikel_LISTE_CLEAN, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfadOutput+"Blame_Import3_"+kunde+"_"+fileName+"_"+fileType+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	return temp
	/*
		writeStueckliste(artikel_LISTE_CLEAN.Artikel, kunde, fileType, fileName)
		writeCSV(artikel_LISTE_CLEAN, fileName)
	*/
}

func writeStueckliste(lagerbestand map[string][]ARTIKEL, kunde string, fileType string, fileName string) {
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

	if err := file2.SaveAs(rootPfadOutput + "Blame_Sum_" + kunde + "_" + fileName + "_" + fileType + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val string) {

	file.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum), val)
	*colNum++
}
func bestellnummerClean3(x string) string {
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	x = strings.ReplaceAll(x, ".", "")
	x = strings.ReplaceAll(x, "/", "")
	x = strings.ReplaceAll(x, ",", "")
	x = strings.ReplaceAll(x, "ü", "ue")
	x = strings.ReplaceAll(x, "ä", "ae")
	x = strings.ReplaceAll(x, "ö", "oe")
	x = strings.ReplaceAll(x, "Ü", "Ue")
	x = strings.ReplaceAll(x, "Ä", "Ae")
	x = strings.ReplaceAll(x, "Ö", "Oe")
	return x
}
func writeCSV(lagerbestand *ARTIKELLISTE, fileName string, orte []string) {
	for _, a := range orte {
		teilvorhanden := false
		var err error
		var file *os.File
		var file2 *os.File
		var w *csv.Writer
		var w2 *csv.Writer
		var bestellSiteca []string
		var bestellKNT []string
		interStueck := 1
		externStueck := 1
		for _, record := range lagerbestand.Artikel {
			switch {
			case a == record[0].BMK.Ortskennzeichen:
				if !teilvorhanden {
					file, err = os.Create(rootPfadOutput + "Blame_SITECA_" + fileName + "_" + bestellnummerClean3(a) + ".csv")
					if err != nil {
						log.Fatalln("failed to open file", err)
					}
					defer file.Close()

					file2, err = os.Create(rootPfadOutput + "Blame_KNT_" + fileName + "_" + bestellnummerClean3(a) + ".csv")
					if err != nil {
						log.Fatalln("failed to open file", err)
					}

					defer file2.Close()

					w = csv.NewWriter_REFAC(file)
					defer w.Flush()
					w2 = csv.NewWriter_REFAC(file2)
					defer w2.Flush()

					headers := []string{
						"",
						"",
						"",
						"Stuecklistenart",
						"",
						"Pos.",
						"ERP",
						"Menge",
						"Herstellernummer",
						"Hersteller",
						"Quelle",
					}

					if err := w.Write(headers); err != nil {
						log.Fatalln("error writing record to file", err)
					}
					if err := w2.Write(headers); err != nil {
						log.Fatalln("error writing record to file", err)
					}

					bestellSiteca = []string{}
					for i := 1; i < 20; i++ {
						bestellSiteca = append(bestellSiteca, "")
					}
					bestellKNT = []string{}
					for i := 1; i < 20; i++ {
						bestellKNT = append(bestellKNT, "")
					}
					interStueck = 1
					externStueck = 1
					teilvorhanden = true
				}

				if record[0].Bestellung_Moeller != 0 {
					bestellSiteca[0] = bestellnummerClean3(fileName)
					bestellSiteca[1] = ""
					bestellSiteca[2] = "1"
					bestellSiteca[3] = "2"
					bestellSiteca[4] = "ABLR"
					bestellSiteca[5] = fmt.Sprintf("%d", interStueck)
					bestellSiteca[6] = record[0].ERP
					bestellSiteca[7] = fmt.Sprintf("%.0f", record[0].Bestellung_Moeller)
					bestellSiteca[8] = record[0].Bestellnummer
					bestellSiteca[9] = record[0].Hersteller
					bestellSiteca[10] = "Moeller"
					interStueck++
					if err := w.Write(bestellSiteca); err != nil {
						log.Fatalln("error writing record to file", err)
					}
				}
				if record[0].Bestellung_KNT != 0 {
					bestellKNT[0] = bestellnummerClean3(fileName)
					bestellKNT[1] = ""
					bestellKNT[2] = "1"
					bestellKNT[3] = "2"
					bestellKNT[4] = "ABLR"
					bestellKNT[5] = fmt.Sprintf("%d", externStueck)
					bestellKNT[6] = record[0].ERP_KNT
					bestellKNT[7] = fmt.Sprintf("%.0f", record[0].Bestellung_KNT)
					bestellKNT[8] = record[0].Bestellnummer
					bestellKNT[9] = record[0].Hersteller
					bestellKNT[10] = "KNT"
					externStueck++
					if err := w2.Write(bestellKNT); err != nil {
						log.Fatalln("error writing record to file", err)
					}
				}
				if record[0].Bestellung_Siteca != 0 {
					bestellSiteca[0] = bestellnummerClean3(fileName)
					bestellSiteca[1] = ""
					bestellSiteca[2] = "1"
					bestellSiteca[3] = "2"
					bestellSiteca[4] = "ABLR"
					bestellSiteca[5] = fmt.Sprintf("%d", interStueck)
					bestellSiteca[6] = record[0].ERP
					bestellSiteca[7] = fmt.Sprintf("%.0f", record[0].Bestellung_Siteca)
					bestellSiteca[8] = record[0].Bestellnummer
					bestellSiteca[9] = record[0].Hersteller
					bestellSiteca[10] = "Siteca"
					interStueck++
					if err := w.Write(bestellSiteca); err != nil {
						log.Fatalln("error writing record to file", err)
					}
				}

			default:
			}
		}
	}
}
