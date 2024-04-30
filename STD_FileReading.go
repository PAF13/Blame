package main

import (
	"fmt"
	"strconv"

	"github.com/xuri/excelize/v2"
)

type Stuckliste_ImportTemplate struct {
	//Cells
	FunktionaleZuordnung       int //==
	Funktionskennzeichen       int //=
	Aufstellungsort            int //++
	Ortskennzeichen            int //+
	Dokumentenart              int //&
	BenutzerdefinierteStruktur int //#
	Anlagennummer              int //empty?
	BMK                        int //-
	ERP                        int //
	Bestellnummer              int //
	Zusatzbestellnummer        int //
	Hersteller                 int //
	ArtikelnummerEplan         int //
	Beschreibung               int //
	Stueckzahl                 int //
	FirstValue                 int //
	Einheit                    int //
	Warengruppe                int //
	Beistellung                int
	Ort                        int
}

func ReadStueckliste(stuecklisteCells Stuckliste_ImportTemplate, stuecklistepfad string, lagerbestand map[string]*Artikel, quelle string) map[string]*Artikel {
	headSkip := stuecklisteCells.FirstValue
	skip := 0
	//opening spreadsheet
	fmt.Println("Opening: " + stuecklistepfad)
	spreadsheet, err := excelize.OpenFile(stuecklistepfad)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}
	//reading rows
	for _, row := range rows {
		//loading map with data as string
		if skip > headSkip {
			setLager(lagerbestand, row, stuecklisteCells, quelle)
		}
		skip++
	}
	if err := spreadsheet.Close(); err != nil {
		fmt.Println(err)
	}
	return lagerbestand
}

func setLager(lagerbestand map[string]*Artikel, row []string, stuecklisteCells Stuckliste_ImportTemplate, quelle string) {
	var source string
	var stueckzahl float64
	var bestellnummer string
	source = quelle
	menge, _ := strconv.ParseFloat(safeStringArrayPull(row, stuecklisteCells.Stueckzahl), 32)

	_, ok := lagerbestand[safeStringArrayPull(row, stuecklisteCells.Bestellnummer)]

	if ok {
		stueckzahl = lagerbestand[safeStringArrayPull(row, stuecklisteCells.Bestellnummer)].Stueckzahl + menge
	} else {
		stueckzahl = menge
	}

	if safeStringArrayPull(row, stuecklisteCells.Bestellnummer) != "" {
		bestellnummer = safeStringArrayPull(row, stuecklisteCells.Bestellnummer)
	} else {
		bestellnummer = "Leer"
	}

	lagerbestand[safeStringArrayPull(row, stuecklisteCells.Bestellnummer)] = &Artikel{
		ERP:                safeStringArrayPull(row, stuecklisteCells.ERP),
		Bestellnummer:      bestellnummer,
		ArtikelnummerEplan: safeStringArrayPull(row, stuecklisteCells.ArtikelnummerEplan),
		Hersteller:         safeStringArrayPull(row, stuecklisteCells.Hersteller),
		Beschreibung:       safeStringArrayPull(row, stuecklisteCells.Beschreibung),
		Stueckzahl:         stueckzahl,
		Einheit:            safeStringArrayPull(row, stuecklisteCells.Einheit),
		Warengruppe:        safeStringArrayPull(row, stuecklisteCells.Warengruppe),
		Quelle:             source,
		Beistellung:        safeStringArrayPull(row, stuecklisteCells.Beistellung),
		Ort:                safeStringArrayPull(row, stuecklisteCells.Ort),
		Aufstellungsort:    safeStringArrayPull(row, stuecklisteCells.Aufstellungsort),
		Ortskennzeichen:    safeStringArrayPull(row, stuecklisteCells.Ortskennzeichen),
	}

}
