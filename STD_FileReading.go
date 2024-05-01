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
	Herstellertyp              int //
	HerstellerEplan            int //
	Bestellnr_L1               int //

}

var stueckliste_Topix = Stuckliste_ImportTemplate{
	Aufstellungsort:    300,
	Ortskennzeichen:    300,
	ERP:                2,
	Bestellnummer:      72,
	Hersteller:         6,
	ArtikelnummerEplan: 187,
	Beschreibung:       24,
	Stueckzahl:         50,
	Einheit:            12,
	FirstValue:         3,
	Warengruppe:        300,
	Beistellung:        300,
	Ort:                300,
	Herstellertyp:      188,
	HerstellerEplan:    300,
	Bestellnr_L1:       272,
}
var stueckliste_Kroenert = Stuckliste_ImportTemplate{
	Aufstellungsort:    300,
	Ortskennzeichen:    300,
	ERP:                1,
	Bestellnummer:      13,
	Hersteller:         20,
	ArtikelnummerEplan: 20,
	Beschreibung:       9,
	Stueckzahl:         3,
	Einheit:            20,
	FirstValue:         4,
	Warengruppe:        20,
	Beistellung:        12,
	Ort:                3,
	Herstellertyp:      300,
	HerstellerEplan:    300,
	Bestellnr_L1:       300,
}
var stueckliste_projekt = Stuckliste_ImportTemplate{
	Aufstellungsort:    2,
	Ortskennzeichen:    3,
	ERP:                7,
	Bestellnummer:      9,
	Hersteller:         11,
	ArtikelnummerEplan: 20,
	Beschreibung:       10,
	Stueckzahl:         5,
	Einheit:            20,
	FirstValue:         7,
	Warengruppe:        20,
	Beistellung:        12,
	Ort:                20,
	Herstellertyp:      300,
	HerstellerEplan:    300,
	Bestellnr_L1:       300,
}

func STD_Read_Lagerbestand(stuecklisteCells Stuckliste_ImportTemplate, stuecklistepfad string, lagerbestand *[]*Artikel, quelle string) {
	headSkip := stuecklisteCells.FirstValue
	skip := 0

	fmt.Println("Opening: " + stuecklistepfad)
	spreadsheet, err := excelize.OpenFile(stuecklistepfad)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}

	for _, row := range rows {
		if skip >= headSkip {
			STD_Set_Lagerbestand(lagerbestand, row, stuecklisteCells, quelle)
		}
		skip++
	}

	if err := spreadsheet.Close(); err != nil {
		fmt.Println(err)
	}
}

func STD_Set_Lagerbestand(lagerbestand *[]*Artikel, row []string, stuecklisteCells Stuckliste_ImportTemplate, quelle string) {
	Stueckzahl, _ := strconv.ParseFloat(safeStringArrayPull(row, stuecklisteCells.Stueckzahl), 32)

	ort := safeStringArrayPull(row, stuecklisteCells.Aufstellungsort) + safeStringArrayPull(row, stuecklisteCells.Ortskennzeichen)

	*lagerbestand = append(*lagerbestand, &Artikel{
		ERP:                safeStringArrayPull(row, stuecklisteCells.ERP),
		Bestellnummer:      safeStringArrayPull(row, stuecklisteCells.Bestellnummer),
		ArtikelnummerEplan: safeStringArrayPull(row, stuecklisteCells.ArtikelnummerEplan),
		Hersteller:         safeStringArrayPull(row, stuecklisteCells.Hersteller),
		Beschreibung:       safeStringArrayPull(row, stuecklisteCells.Beschreibung),
		Stueckzahl:         Stueckzahl,
		Einheit:            safeStringArrayPull(row, stuecklisteCells.Einheit),
		Warengruppe:        safeStringArrayPull(row, stuecklisteCells.Warengruppe),
		Quelle:             quelle,
		Beistellung:        safeStringArrayPull(row, stuecklisteCells.Beistellung),
		Ort:                ort,
		Aufstellungsort:    safeStringArrayPull(row, stuecklisteCells.Aufstellungsort),
		Ortskennzeichen:    safeStringArrayPull(row, stuecklisteCells.Ortskennzeichen),
		Herstellertyp:      safeStringArrayPull(row, stuecklisteCells.Herstellertyp),
		HerstellerEplan:    safeStringArrayPull(row, stuecklisteCells.HerstellerEplan),
		Bestellnr_L1:       safeStringArrayPull(row, stuecklisteCells.Bestellnr_L1),
	})

}
