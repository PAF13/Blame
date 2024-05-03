package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func Write_Lager(pfad string, ListeRaw *[]*Artikel) {
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

	for _, artikelListe := range *ListeRaw {

		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+0)), row_Lagerbestand_Kroenert), artikelListe.ERP)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+1)), row_Lagerbestand_Kroenert), artikelListe.Stueckzahl)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+2)), row_Lagerbestand_Kroenert), artikelListe.Hersteller)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+3)), row_Lagerbestand_Kroenert), artikelListe.Bestellnummer)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+4)), row_Lagerbestand_Kroenert), artikelListe.ArtikelnummerEplan)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+5)), row_Lagerbestand_Kroenert), artikelListe.Beschreibung)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+6)), row_Lagerbestand_Kroenert), artikelListe.Warengruppe)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+7)), row_Lagerbestand_Kroenert), artikelListe.Quelle)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+8)), row_Lagerbestand_Kroenert), artikelListe.Stand)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+9)), row_Lagerbestand_Kroenert), artikelListe.Beistellung)
		file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+10)), row_Lagerbestand_Kroenert), artikelListe.Aufstellungsort+artikelListe.Ortskennzeichen)
		row_Lagerbestand_Kroenert++

	}

	if err := file_Lagerbestand_Kroenert.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}

}

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

func Write_Lager3(pfad string, ListeRaw map[string]map[string]*Artikel) {
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

	for _, arti := range ListeRaw {
		for _, ann := range arti {
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+0)), row_Lagerbestand_Kroenert), ann.ERP)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+1)), row_Lagerbestand_Kroenert), ann.Stueckzahl)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+2)), row_Lagerbestand_Kroenert), ann.Hersteller)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+3)), row_Lagerbestand_Kroenert), ann.Bestellnummer)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+4)), row_Lagerbestand_Kroenert), ann.ArtikelnummerEplan)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+5)), row_Lagerbestand_Kroenert), ann.Beschreibung)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+6)), row_Lagerbestand_Kroenert), ann.Warengruppe)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+7)), row_Lagerbestand_Kroenert), ann.Quelle)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+8)), row_Lagerbestand_Kroenert), ann.Stand)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+9)), row_Lagerbestand_Kroenert), ann.Beistellung)
			file_Lagerbestand_Kroenert.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+10)), row_Lagerbestand_Kroenert), ann.Aufstellungsort+ann.Ortskennzeichen)
			row_Lagerbestand_Kroenert++
		}

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
	//P_file := file
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

	rowNum := 3
	for _, value := range liste {
		colNum := 6

		lineWriter(file, "Sheet1", &colNum, &rowNum, value.ERP)
		lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%f", value.Stueckzahl))
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Hersteller)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Typ)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnummer)
		lineWriter(file, "Sheet1", &colNum, &rowNum, "")
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Aufstellungsort+value.Ortskennzeichen)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Herstellertyp)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.HerstellerEplan)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnr_L1)
		rowNum++

	}

	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}
}

func lineWriter(file *excelize.File, sheet string, colNum *int, rowNum *int, val string) {
	file.SetCellValue(sheet, fmt.Sprintf("%s%d", string(rune(65+*colNum)), *rowNum), val)
	*colNum++

}
func STD_Write_Stueckliste(pfad string, Lagerbestand []*Artikel) {
	file := excelize.NewFile()
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
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 1), header)
	}
	for i, header := range headers2 {
		file.SetCellValue("Sheet1", fmt.Sprintf("%s%d", string(rune(65+i)), 2), header)
	}

	rowNum := 3
	for _, value := range Lagerbestand {
		colNum := 0

		lineWriter(file, "Sheet1", &colNum, &rowNum, value.ERP)
		lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", value.Stueckzahl))
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Hersteller)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnummer)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.ArtikelnummerEplan)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Beschreibung)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Warengruppe)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Quelle)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Stand)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Beistellung)
		lineWriter(file, "Sheet1", &colNum, &rowNum, value.Aufstellungsort+value.Ortskennzeichen)
		rowNum++
	}

	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}
}

func STD_Write_Stueckliste2(pfad string, Lagerbestand []*Artikel, ort string) {
	file := excelize.NewFile()

	rowNum := 1

	for _, value := range Lagerbestand {
		colNum := 0
		fmt.Printf("ERP: %-20s", value.ERP)
		fmt.Printf("Stückzahl: %-10.2f", value.Stueckzahl)
		fmt.Printf("Bestellnummer: %-25s", value.Bestellnummer)
		fmt.Printf("Ort: %-20s", value.Ort)
		fmt.Printf("\n")

		if value.Quelle != "Siteca" {
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Hersteller)
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnummer)
		} else {
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
		}

		lineWriter(file, "Sheet1", &colNum, &rowNum, "1")
		lineWriter(file, "Sheet1", &colNum, &rowNum, "2")
		lineWriter(file, "Sheet1", &colNum, &rowNum, "ABLR")
		lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%d", rowNum))

		if value.Quelle != "Siteca" {
			lineWriter(file, "Sheet1", &colNum, &rowNum, "")
		} else {
			lineWriter(file, "Sheet1", &colNum, &rowNum, value.ERP)
		}

		lineWriter(file, "Sheet1", &colNum, &rowNum, fmt.Sprintf("%.0f", value.Stueckzahl))
		//lineWriter(file, "Sheet1", &colNum, &rowNum, value.Bestellnummer)
		rowNum++
	}

	fmt.Println("Finished writing: " + pfad)
	if err := file.SaveAs("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\BlameOutput\\" + pfad); err != nil {
		fmt.Println(err)
	}
}
