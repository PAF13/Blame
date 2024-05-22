package main

import (
	"encoding/csv"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/xuri/excelize/v2"
)

func XMLEplan() {
	rootPfad := "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EDC\\"
	topixName := "Eplan2024_Datenbank.xml"

	headerNamen := []string{"ERP", "Bestellnummer", "Artikelnummer", "Hersteller", "Error"}
	richtig := [][]string{headerNamen}
	falsch := [][]string{headerNamen}

	artikelEPlanRaw := ImportXMLEPLAN(rootPfad, topixName)
	artikelEPlan := map[string]*Product{}
	for _, b := range artikelEPlanRaw.SO117 {
		artikelEPlan["BESTELLNUMMER|"+b.P22222+"|"+b.P22002] = newProductEplan1(b)
	}
	writeJsonFile(&artikelEPlan, "Step2_0_EPlan")

	artikelTopixRaw, length := ImportXLSX("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameInput\\Topix.xlsx", "Step1_0_TOPIX")
	artikelTopix := map[string]*Product{}
	for a := range length {
		key := "BESTELLNUMMER|" + artikelTopixRaw["Hersteller"][a] + "|" + artikelTopixRaw["HerstellerNummer"][a]
		_, ok := artikelTopix[key]
		if !ok {
			artikelTopix[key] = newProductTopix1(artikelTopixRaw, a)
		} else {
			fehler(&falsch, artikelTopixRaw, a)
		}

	}
	writeJsonFile(&artikelTopix, "Step2_0_Topix")

	for a, b := range artikelTopix {
		_, ok := artikelEPlan[a]
		if ok {
			falsch = append(falsch, []string{b.ERPTopix, b.Bestellnummer, b.Artikelnummer, b.Hersteller, "Nicht in Eplan / Falscher Herstellername"})
		} else {
			richtig = append(richtig, []string{b.ERPTopix, b.Bestellnummer, b.Artikelnummer, b.Hersteller, "Richtig"})
		}
	}

	writeCSVFile("richtig", &richtig)
	writeCSVFile("falsch", &falsch)
}
func fehler(csv *[][]string, raw map[string][]string, a int) {
	*csv = append(*csv, []string{
		raw["Artikelnummer"][a],
		raw["HerstellerNummer"][a],
		raw["EPlan_Artikelnr_"][a],
		raw["Hersteller"][a],
		"Duplicate Bestellnummer | Topix",
	})
}
func writeCSVFile(fileName string, rows *[][]string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered. Error:\n", r)
		}
	}()
	csvFile, err := os.Create("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EDC\\Blame_" + fileName + ".csv")
	if err != nil {
		log.Printf("failed creating file: %s\n", err)
	}
	defer csvFile.Close()
	w := csv.NewWriter_REFAC(csvFile)
	defer w.Flush()

	writeCSVFileBody(w, rows)

}
func writeCSVFileBody(w *csv.Writer, rows *[][]string) {
	for _, row := range *rows {
		if err := w.Write(row); err != nil {
			log.Println("error writing record to file", err)
		}
	}

}
func ImportXMLEPLAN(pfad string, fileName string) *EplanPxfRoot {
	xmlFile, err := os.Open(pfad + fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil
	}
	defer xmlFile.Close()

	byteValue, _ := ioutil.ReadAll(xmlFile)

	var root EplanPxfRoot
	err = xml.Unmarshal(byteValue, &root)
	if err != nil {
		fmt.Println("Error unmarshalling XML:", err)
		return nil
	}
	writeJsonFile(&root, "Step1_0_EPlan")
	return &root
}

func ImportXLSX(pfad string, fileName string) (map[string][]string, int) {
	// Open the Excel file
	f, err := excelize.OpenFile(pfad)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	// Get all the rows in the first sheet
	sheetName := f.GetSheetName(0)
	rows, err := f.GetRows(sheetName)
	if err != nil {
		log.Fatal(err)
	}

	if len(rows) == 0 {
		log.Fatal("The sheet is empty")
	}

	// Assuming the first row contains the headers
	length := 0
	header := 2
	headers := rows[header]
	columns := make(map[string][]string)

	// Initialize the map with headers
	for _, header := range headers {
		columns[header] = []string{}
	}

	for _, row := range rows[3:] {
		for i, cell := range row {
			columns[headers[i]] = append(columns[headers[i]], cell)
			if len(columns[headers[i]]) > length {
				length = len(columns[headers[i]])
			}
		}

	}

	for a := range columns {
		for len(columns[a]) < length {
			columns[a] = append(columns[a], "")
		}
	}

	for a, values := range columns {
		fmt.Printf("Header: %-40s Länge: %-30d\n", a, len(values))
	}

	writeJsonFile(columns, fileName)
	return columns, length
}
