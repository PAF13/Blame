package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

var rootPfad string

func ImportFile() {
	log.Println("Starting import")
	rootPfad = "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\"
	pfad := [][]string{
		//{rootPfad + "Topix_Artikel20240502.xlsx", "Lager", "SITECA"},
		//{rootPfad + "Kopie von Lagerhueter_26_04_2024.xlsx", "Lager", "KNT"},
		//{rootPfad + "8000772_Stückliste.xlsx", "Steuckliste", "KNT"},
		//{rootPfad + "EPlanOutput\\EPlan_Verbindungsliste.xml","verbindungsliste", "SITECA"},
		//{rootPfad + "EPlanOutput\\EPlan_Betriebsmitttel.xml","Steuckliste", "SITECA"},
	}
	for _, pfad1 := range pfad {
		file, err := os.Open(pfad1[0])
		if err != nil {
			log.Println(err)
		}
		defer file.Close()

		byteValue, _ := io.ReadAll(file)

		switch {
		case pfad1[1] == "Lager":
			if pfad1[2] == "KNT" {
				excelImport := EXCEL_SIMPLE{
					Header:               0,
					BMKVollständig:       500,
					FunktionaleZuordnung: 500,
					Funktionskennzeichen: 500,
					Aufstellungsort:      500,
					Ortskennzeichen:      500,
					BMK:                  500,
					ERP:                  1,
					Bestellnummer:        13,
					Bezeichnung:          500,
					Beschreibung:         11,
					Stueckzahl:           3,
					Einheit:              500,
					Verpackungseinheit:   500,
					Lagerort:             500,
					Hersteller:           500,
					Beistellung:          500,
				}
				FileFactory(byteValue, pfad1[0], pfad1[2], excelImport, pfad1[1])
			} else if pfad1[2] == "SITECA" {
				excelImport := EXCEL_SIMPLE{
					Header:               0,
					BMKVollständig:       500,
					FunktionaleZuordnung: 500,
					Funktionskennzeichen: 500,
					Aufstellungsort:      500,
					Ortskennzeichen:      500,
					BMK:                  500,
					ERP:                  2,
					Bestellnummer:        72,
					Bezeichnung:          500,
					Beschreibung:         24,
					Stueckzahl:           5,
					Einheit:              500,
					Verpackungseinheit:   500,
					Lagerort:             500,
					Hersteller:           6,
					Beistellung:          500,
				}
				FileFactory(byteValue, pfad1[0], pfad1[2], excelImport, pfad1[1])
			}
		case pfad1[1] == "Steuckliste":
			if pfad1[2] == "KNT" {
				excelImport := EXCEL_SIMPLE{
					Header:               0,
					BMKVollständig:       500,
					FunktionaleZuordnung: 0,
					Funktionskennzeichen: 1,
					Aufstellungsort:      2,
					Ortskennzeichen:      3,
					BMK:                  4,
					ERP:                  7,
					Bestellnummer:        9,
					Bezeichnung:          500,
					Beschreibung:         10,
					Stueckzahl:           5,
					Einheit:              500,
					Verpackungseinheit:   500,
					Lagerort:             500,
					Hersteller:           11,
					Beistellung:          12,
				}
				FileFactory(byteValue, pfad1[0], pfad1[2], excelImport, pfad1[1])
			}
		case pfad1[2] == "BLAME":
		default:
			fmt.Println("The following was not found" + pfad1[2])
		}

	}
	loadFile()
}
func FileFactory(byteValue []byte, pfad string, quelle string, excelImport EXCEL_SIMPLE, fileType string) {

	file, err := excelize.OpenFile(pfad)
	if err != nil {
		log.Fatal(err)
	}
	rows, err := file.GetRows(file.GetSheetList()[0])
	if err != nil {
		log.Fatal(err)
	}

	lagerbestand := map[string][]ARTIKEL{}

	for _, row := range rows {
		var newRow [501]string
		copy(newRow[:], row)
		bestellnummer := bestellnummerClean(newRow[excelImport.Bestellnummer])
		stueckzahl, _ := strconv.ParseFloat(newRow[excelImport.Stueckzahl], 64)

		lagerbestand[bestellnummer] = append(lagerbestand[bestellnummer], ARTIKEL{
			ERP:           newRow[excelImport.ERP],
			Bestellnummer: bestellnummer,
			Stueckzahl:    stueckzahl,
			Quelle:        quelle,
			Beschreibung:  newRow[excelImport.Beschreibung],
			Hersteller:    newRow[excelImport.Hersteller],
			Beistellung:   newRow[excelImport.Beistellung],
			Fehler:        []string{},
			BMK: BETRIEBSMITELLKENNZEICHEN{
				FunktionaleZuordnung: newRow[excelImport.FunktionaleZuordnung],
				Funktionskennzeichen: newRow[excelImport.Funktionskennzeichen],
				Aufstellungsort:      newRow[excelImport.Aufstellungsort],
				Ortskennzeichen:      newRow[excelImport.Ortskennzeichen],
				BMK:                  newRow[excelImport.BMK],
			},
		})

	}
	content, err := json.MarshalIndent(lagerbestand, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfad+"Test_Project\\Blame_"+fileType+"_"+quelle+".json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
	writeStueckliste(lagerbestand, quelle, fileType)
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
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Hersteller")
	file2.SetColWidth("Sheet1", "H", "H", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Menge")
	file2.SetColWidth("Sheet1", "I", "I", 10)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung Moeller")
	file2.SetColWidth("Sheet1", "J", "J", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung KNT")
	file2.SetColWidth("Sheet1", "K", "K", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Bestellung Siteca")
	file2.SetColWidth("Sheet1", "L", "L", 20)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Beisteller")
	file2.SetColWidth("Sheet1", "M", "M", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Quelle")
	file2.SetColWidth("Sheet1", "N", "N", 13)
	lineWriter(file2, "Sheet1", &colNum, &rowNum, "Beschreibung")
	file2.SetColWidth("Sheet1", "O", "O", 70)
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
		Range:             "A1:O" + fmt.Sprintf("%d", rowNum),
		Name:              "table",
		StyleName:         "TableStyleMedium2",
		ShowFirstColumn:   true,
		ShowLastColumn:    true,
		ShowRowStripes:    &disable,
		ShowColumnStripes: true,
	})

	if err := file2.SaveAs(rootPfad + "Test_Project\\Blame_" + fileType + "_" + quelle + ".xlsx"); err != nil {
		fmt.Println(err)
	}
}
func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val string) {
	file.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum), val)
	*colNum++
}
func loadFile() {
	pfad := [][]string{
		{rootPfad + "Test_Project\\Blame_Lager_KNT.json", "Lager", "BLAME"},
		{rootPfad + "Test_Project\\Blame_Lager_SITECA.json", "Lager", "BLAME"},
		{rootPfad + "Test_Project\\Blame_Steuckliste_KNT.json", "Steuckliste", "BLAME"},
	}

	jsonFile_KNT, err := os.Open(pfad[0][0])
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_KNT.Close()
	jsonFile_SITECA, err := os.Open(pfad[1][0])
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_SITECA.Close()
	jsonFile_LISTE, err := os.Open(pfad[2][0])
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile_LISTE.Close()

	byteValue_KNT, _ := ioutil.ReadAll(jsonFile_KNT)
	byteValue_SITECA, _ := ioutil.ReadAll(jsonFile_SITECA)
	byteValue_LISTE, _ := ioutil.ReadAll(jsonFile_LISTE)

	artikel_KNT := make(map[string][]ARTIKEL)
	artikel_SITECA := make(map[string][]ARTIKEL)
	artikel_LISTE := make(map[string][]ARTIKEL)
	artikel_LISTE_CLEAN := make(map[string]*ARTIKEL)
	artikel_LISTE_WRITE := make(map[string][]ARTIKEL)

	json.Unmarshal(byteValue_KNT, &artikel_KNT)
	json.Unmarshal(byteValue_SITECA, &artikel_SITECA)
	json.Unmarshal(byteValue_LISTE, &artikel_LISTE)

	for a, b := range artikel_LISTE {
		for aa, bb := range b {

			_, ok := artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen]
			if !ok {
				artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen] = &artikel_LISTE[a][aa]
			} else {
				fmt.Printf("1Bestellnummer: %-40s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].Bestellnummer)
				fmt.Printf("BMK: %-20s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].BMK.Aufstellungsort)
				fmt.Printf(" %-20s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].BMK.Ortskennzeichen)
				fmt.Printf("Stueckzahl alt: %-20.0f", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].Stueckzahl)
				fmt.Printf("Stueckzahl neu: %-20.0f\n", bb.Stueckzahl)
				sadfd := artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].Stueckzahl + bb.Stueckzahl
				artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].updateStueckzahl(sadfd)
				fmt.Printf("2Bestellnummer: %-40s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].Bestellnummer)
				fmt.Printf("BMK: %-20s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].BMK.Aufstellungsort)
				fmt.Printf(" %-20s", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].BMK.Ortskennzeichen)
				fmt.Printf("Stueckzahl alt: %-20.0f\n", artikel_LISTE_CLEAN[bb.Bestellnummer+bb.BMK.Aufstellungsort+bb.BMK.Ortskennzeichen].Stueckzahl)
			}

		}
	}

	for a, b := range artikel_LISTE_CLEAN {
		_, ok := artikel_SITECA[b.Bestellnummer]
		stueckzahl := b.Stueckzahl
		if b.Beistellung == "SITECA" {
			_, ok_KNT := artikel_KNT[b.Bestellnummer]
			if ok_KNT {
				for a2, b2 := range artikel_KNT[b.Bestellnummer] {
					if stueckzahl >= b2.Stueckzahl {
						b.Bestellung_KNT = b.Bestellung_KNT + b2.Stueckzahl
						artikel_KNT[b.Bestellnummer][a2].Stueckzahl = 0
						stueckzahl = stueckzahl - b2.Stueckzahl
					} else {
						b.Bestellung_KNT = stueckzahl
						artikel_KNT[b.Bestellnummer][a2].Stueckzahl = artikel_KNT[b.Bestellnummer][a2].Stueckzahl - stueckzahl
						stueckzahl = 0
					}
				}
			}
			b.Bestellung_Siteca = stueckzahl
		}
		if ok {
			erpList := ""
			hersteller := ""
			for _, bb := range artikel_SITECA[b.Bestellnummer] {
				if len(artikel_SITECA[bb.Bestellnummer]) > 1 {
					erpList = erpList + bb.ERP + " | "
				} else {
					erpList = bb.ERP
				}
				hersteller = bb.Hersteller
			}
			b.ERP = erpList
			b.Quelle = "SITECA"
			b.Hersteller = hersteller

			artikel_LISTE_WRITE[a] = append(artikel_LISTE_WRITE[a], *b)
		} else {
			artikel_LISTE_WRITE[a] = append(artikel_LISTE_WRITE[a], *b)
		}

	}
	writeStueckliste(artikel_LISTE_WRITE, "_Stueckliste", "_Sum")

	content, err := json.MarshalIndent(artikel_LISTE_CLEAN, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile(rootPfad+"Test_Project\\Blame_Stueckliste_Clean_KNT.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func (s *ARTIKEL) updateStueckzahl(bb float64) *ARTIKEL {
	s.Stueckzahl = bb
	return s
}
func (structType *EplanAuswertungXML) convertFile2(byteValue []byte) {
	betriebsmittel = map[string][50]ARTIKEL{}
	artikelSlice := [50]ARTIKEL{}
	xml.Unmarshal(byteValue, &structType)
	line := structType.Document.Page.Line
	for a, aa := range line {
		fmt.Printf("Source ID: %-50s", aa.Label.SourceID)
		fmt.Printf("Anzahl: %-50d\n", a+1)
		betriebsmittelkennzeichen := BETRIEBSMITELLKENNZEICHEN{}
		P_betriebsmittelkennzeichen := &betriebsmittelkennzeichen
		artikel := []ARTIKEL{}
		P_artikel := &artikel
		for _, bb := range aa.Label.Property {

			switch bb.PropertyName {
			case "BMK (vollständig)":
				P_betriebsmittelkennzeichen.BMKVollständig = bb.PropertyValue

			case "Name (identifizierend)":
				P_betriebsmittelkennzeichen.BMKidentifizierung = bb.PropertyValue

			case "Funktionale Zuordnung":
				P_betriebsmittelkennzeichen.FunktionaleZuordnung = bb.PropertyValue

			case "Funktionskennzeichen":
				P_betriebsmittelkennzeichen.Funktionskennzeichen = bb.PropertyValue

			case "Aufstellungsort":
				P_betriebsmittelkennzeichen.Aufstellungsort = bb.PropertyValue

			case "Ortskennzeichen":
				P_betriebsmittelkennzeichen.Ortskennzeichen = bb.PropertyValue

			case "BMK (identifizierend, ohne Projektstrukturen)":
				P_betriebsmittelkennzeichen.BMK = bb.PropertyValue

			case "BMK: Kennbuchstabe":
				P_betriebsmittelkennzeichen.Kennbuchstabe = bb.PropertyValue

			case "Artikelnummer":
				*P_artikel = append(*P_artikel, ARTIKEL{ArtikelnummerEplan: bb.PropertyValue})
			default:
				fmt.Printf("Missing | Name: %-50s", bb.PropertyName)
				fmt.Printf("Value: %-50s", bb.PropertyValue)
				fmt.Printf("\n")
			}

		}
		copy(artikelSlice[:], artikel)
		betriebsmittel[betriebsmittelkennzeichen.BMKidentifizierung] = artikelSlice
	}

	content, err := json.MarshalIndent(betriebsmittel, "", "\t")
	if err != nil {
		fmt.Println(err)
	}

	err = ioutil.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\blame_Stueckliste.json", content, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
