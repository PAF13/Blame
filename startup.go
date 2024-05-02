package main

import (
	"fmt"
	"strings"
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

	fmt.Println("Loading...")

	STD_Read_Lagerbestand(stueckliste_Topix, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240502.xlsx", &lagerbestand_Siteca, "Siteca")
	STD_Write_Stueckliste("Liste_Siteca.xlsx", lagerbestand_Siteca)
	for _, value := range lagerbestand_Siteca {
		if strings.Contains(value.Hersteller, "RIT") {
			fmt.Printf("ERP: %-20s", value.ERP)
			fmt.Printf("Stückzahl: %-10.2f", value.Stueckzahl)
			fmt.Printf("Bestellnummer: %-20s", value.Bestellnummer)
			fmt.Printf("Hersteller: %-20s", value.Hersteller)
			fmt.Printf("Bestellnr_L1: %-20s", value.Bestellnr_L1)
			fmt.Printf("Herstellertyp: %-20s", value.Herstellertyp)
			fmt.Printf("Bezeichnung: %-20s", value.Bezeichnung)
			fmt.Printf("Beistellung: %-20s", value.Beistellung)
			fmt.Printf("\n")
		}
	}
	for _, b := range lagerbestand_Siteca {
		lagerbestand_Siteca_Map[b.Bestellnummer] = b
	}

	STD_Read_Lagerbestand(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", &lagerbestand_Kroenert, "KNT")
	STD_Write_Stueckliste("Liste_Kroenert.xlsx", lagerbestand_Kroenert)
	STD_Clean_Lagerbestand(lagerbestand_Kroenert)
	STD_Write_Stueckliste("Liste_Kroenert_Clean.xlsx", lagerbestand_Kroenert)

	STD_Read_Lagerbestand(stueckliste_projekt, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\8000772_Stückliste.xlsx", &Stueckliste, "KNT_Stückliste")
	STD_Write_Stueckliste("Liste_Stueckliste.xlsx", Stueckliste)
	STD_Clean_Lagerbestand(Stueckliste)
	STD_Write_Stueckliste("Liste_Stueckliste_Clean.xlsx", Stueckliste)

	STD_Sum(&Stueckliste, &Stueckliste2, Stueckliste_Map)

	STD_Write_Stueckliste("Liste_Stueckliste_Clean_Sum.xlsx", Stueckliste2)

	fmt.Println("Startup finished")

	return true
}
