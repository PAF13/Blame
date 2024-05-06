package main

import (
	"strconv"
)




func STD_Set_Lagerbestand(lagerbestand *[]*Artikel, row []string, stuecklisteCells Stuckliste_ImportTemplate, quelle string) {
	Stueckzahl, _ := strconv.ParseFloat(safeStringArrayPull(row, stuecklisteCells.Stueckzahl), 64)

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
		Bezeichnung:        safeStringArrayPull(row, stuecklisteCells.Bezeichnung),
	})

}
