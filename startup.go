package main

import (
	"fmt"
)

var lagerbestand_Siteca = []*Artikel{}
var lagerbestand_Kroenert = []*Artikel{}
var lagerbestand_Siteca_Map = map[string]*Artikel{}

func (a *App) BlameStartup() bool {
	INIT_SETTINGS()
	INIT_ARTIKELSTAMMDATEN()
	return true
}
func INIT_ARTIKELSTAMMDATEN(){
}
func (a *App) BlameStartup2() bool {
	lagerbestand_Siteca = []*Artikel{}
	lagerbestand_Siteca_Map = make(map[string]*Artikel) // key: Bestellnummer
	lagerbestand_Kroenert = []*Artikel{}
	startUp := false
	fmt.Println("Loading...")
	if startUp {
		STD_Read_Lagerbestand(stueckliste_Topix, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240502.xlsx", &lagerbestand_Siteca, "Siteca")
		STD_Write_Stueckliste("Liste_Siteca.xlsx", lagerbestand_Siteca)
		for _, b := range lagerbestand_Siteca {
			lagerbestand_Siteca_Map[b.Bestellnummer] = b
		}

		STD_Read_Lagerbestand(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", &lagerbestand_Kroenert, "KNT")
		//STD_Write_Stueckliste("Liste_Kroenert.xlsx", lagerbestand_Kroenert)
		STD_Clean_Lagerbestand(lagerbestand_Kroenert)
		//STD_Write_Stueckliste("Liste_Kroenert_Clean.xlsx", lagerbestand_Kroenert)
	}
	fmt.Println("Startup finished")

	return true
}
