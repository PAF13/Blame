package main

import (
	"fmt"
)

var stueckliste_Topix = Stuckliste_ImportTemplate{
	ERP:                0,
	Bestellnummer:      70,
	Hersteller:         4,
	ArtikelnummerEplan: 300,
	Beschreibung:       188,
	Stueckzahl:         1,
	Einheit:            20,
	FirstValue:         3,
	Warengruppe:        300,
	Beistellung:        300,
	Ort:                300,
}
var stueckliste_Kroenert = Stuckliste_ImportTemplate{
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
	Beistellung:        20,
	Ort:                20,
}

var lagerbestand_Siteca map[string]*Artikel   // Key: Bestellnummer
var lagerbestand_Kroenert map[string]*Artikel // Key: Bestellnummer

func (a *App) BlameStartup() bool {
	lagerbestand_Siteca = map[string]*Artikel{}
	lagerbestand_Kroenert = map[string]*Artikel{}

	fmt.Println("Reading Lagerbestand Siteca")
	lagerbestand_Siteca = ReadStueckliste(stueckliste_Topix, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Artikelexport.xlsx", lagerbestand_Siteca, "Siteca")
	fmt.Println("Reading Lagerbestand Kroenert")
	lagerbestand_Kroenert = ReadStueckliste(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", lagerbestand_Kroenert, "KNT")
	lagerbestand_Kroenert = Sitecavergleich(lagerbestand_Siteca, lagerbestand_Kroenert)
	fmt.Println("Startup finished")

	return true
}
