package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Write_Lager2(pfad string, ListeRaw map[string]*Artikel) {
	file_Lagerbestand_Kroenert := excelize.NewFile()

	row_Lagerbestand_Kroenert := 3

	headers := []string{
		pfad,
	}
	headers2 := []string{
		"ERP",
		"Menge",
		"Hersteller",
		"Bestellnummer",
		"Eplannummmer",
		"Beschreibung",
		"Warengruppe",
		"Quelle",
		"Stand",
		"Bereitsteller",
		"Ort",
	}
	for i, header := range headers {
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for i, header := range headers2 {
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 2), header)
	}

	for name := range ListeRaw {
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+0)), row_Lagerbestand_Kroenert), ListeRaw[name].ERP)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+1)), row_Lagerbestand_Kroenert), ListeRaw[name].Stueckzahl)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+2)), row_Lagerbestand_Kroenert), ListeRaw[name].Hersteller)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+3)), row_Lagerbestand_Kroenert), ListeRaw[name].Bestellnummer)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+4)), row_Lagerbestand_Kroenert), ListeRaw[name].ArtikelnummerEplan)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+5)), row_Lagerbestand_Kroenert), ListeRaw[name].Beschreibung)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+6)), row_Lagerbestand_Kroenert), ListeRaw[name].Warengruppe)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+7)), row_Lagerbestand_Kroenert), ListeRaw[name].Quelle)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+8)), row_Lagerbestand_Kroenert), ListeRaw[name].Stand)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+9)), row_Lagerbestand_Kroenert), ListeRaw[name].Beistellung)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+10)), row_Lagerbestand_Kroenert), ListeRaw[name].Aufstellungsort+ListeRaw[name].Ortskennzeichen)
		row_Lagerbestand_Kroenert++

	}

	if err := file_Lagerbestand_Kroenert.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}

}

func Write_Stueckliste(lagerbestand_Siteca map[string]*Artikel, lagerbestand_Kroenert map[string]*Artikel, stueckliste_Projekt map[string]*Artikel) {

	Bestell_Siteca := make(map[string]*Artikel)   // Key: Bestellnummer
	Bestell_Kroenert := make(map[string]*Artikel) // Key: Bestellnummer
	Bestell_Extern := make(map[string]*Artikel)   // Key: Bestellnummer

	for artikel := range stueckliste_Projekt {
		_, Siteca_ok := lagerbestand_Siteca[artikel]
		//var lagerbestand_Siteca_num bool
		_, Kroenert_ok := lagerbestand_Kroenert[artikel]
		//var lagerbestand_Kroenert_num bool

		var lastValue float64 = stueckliste_Projekt[artikel].Stueckzahl

		if Siteca_ok && lastValue != 0 {
			if lagerbestand_Siteca[artikel].Stueckzahl != 0 {
				if lagerbestand_Siteca[artikel].Stueckzahl >= lastValue && lagerbestand_Siteca[artikel].Stueckzahl != 0 && lastValue != 0 {
					sortDatatoFile(Bestell_Siteca, stueckliste_Projekt, artikel, lastValue)
					lastValue = 0
				} else {
					sortDatatoFile(Bestell_Siteca, stueckliste_Projekt, artikel, lastValue)
					lastValue = lastValue - lagerbestand_Siteca[artikel].Stueckzahl
				}
			}
		}
		if Kroenert_ok {
			if lagerbestand_Kroenert[artikel].Stueckzahl >= lastValue && lagerbestand_Kroenert[artikel].Stueckzahl != 0 && lastValue != 0 {
				if lagerbestand_Kroenert[artikel].Stueckzahl != 0 {
					sortDatatoFile(Bestell_Kroenert, stueckliste_Projekt, artikel, lastValue)
					lastValue = 0
				} else {
					sortDatatoFile(Bestell_Kroenert, stueckliste_Projekt, artikel, lastValue)
					lastValue = lastValue - lagerbestand_Kroenert[artikel].Stueckzahl
				}
			}
		}
		if lastValue != 0 {
			sortDatatoFile(Bestell_Extern, stueckliste_Projekt, artikel, lastValue)
		}
	}

	fmt.Println("Writing Bestellung Siteca")
	Stueckliste("Bestellung_Siteca.xlsx", Bestell_Siteca)
	fmt.Println("Writing Bestellung Kroenert")
	Stueckliste("Bestellung_Kroenert.xlsx", Bestell_Kroenert)
	fmt.Println("Writing Bestellung Extern")
	Stueckliste("Bestellung_Extern.xlsx", Bestell_Extern)

}

func sortDatatoFile(bestell map[string]*Artikel, stueckliste_Projekt map[string]*Artikel, artikel string, lastValue float64) {
	bestell[artikel] = &Artikel{
		ERP:                stueckliste_Projekt[artikel].ERP,
		Bestellnummer:      stueckliste_Projekt[artikel].Bestellnummer,
		ArtikelnummerEplan: stueckliste_Projekt[artikel].ArtikelnummerEplan,
		Hersteller:         stueckliste_Projekt[artikel].Hersteller,
		Beschreibung:       stueckliste_Projekt[artikel].Beschreibung,
		Stueckzahl:         lastValue,
		Einheit:            stueckliste_Projekt[artikel].Einheit,
		Warengruppe:        stueckliste_Projekt[artikel].Warengruppe,
		Quelle:             stueckliste_Projekt[artikel].Quelle,
	}
}

func Stueckliste(pfad string, liste map[string]*Artikel) {
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
	for _, value := range liste {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+6)), line), value.ERP)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+7)), line), value.Stueckzahl)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+8)), line), value.Hersteller)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+9)), line), value.Typ)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+10)), line), value.Bestellnummer)
		//file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+11)), line), value.Note)
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+12)), line), value.Aufstellungsort+value.Ortskennzeichen)
		line++

	}

	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}
}
