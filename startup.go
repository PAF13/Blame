package main

import (
	"fmt"
)

var lagerbestand_Siteca = []*Artikel{}
var lagerbestand_Kroenert = []*Artikel{}
var lagerbestand_Siteca_Map = map[string]*Artikel{}

func (a *App) BlameStartup() bool {
	lagerbestand_Siteca = []*Artikel{}
	lagerbestand_Siteca_Map = make(map[string]*Artikel) // key: Bestellnummer
	lagerbestand_Kroenert = []*Artikel{}
	Stueckliste := []*Artikel{}
	Stueckliste2 := []*Artikel{}
	Stueckliste_Map := make(map[string]*Artikel)

	fmt.Println("Reading Lagerbestand Siteca")
	lagerbestand_Siteca = ReadStueckliste(stueckliste_Topix, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Artikelexport.xlsx", lagerbestand_Siteca, "Siteca")
	fmt.Println("Reading Lagerbestand Kroenert")
	lagerbestand_Kroenert = ReadStueckliste(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", lagerbestand_Kroenert, "KNT")
	lagerbestand_Kroenert = Sitecavergleich(lagerbestand_Siteca, lagerbestand_Kroenert)
	fmt.Println("Startup finished")

	return true
}
