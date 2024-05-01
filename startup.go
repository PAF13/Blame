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

	fmt.Println("Loading...")

	STD_Read_Lagerbestand(stueckliste_Topix, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240430.xlsx", &lagerbestand_Siteca, "Siteca")
	STD_Write_Stueckliste("Liste_Siteca.xlsx", lagerbestand_Siteca)

	for _, b := range lagerbestand_Siteca {
		lagerbestand_Siteca_Map[b.Bestellnummer] = b
	}

	STD_Read_Lagerbestand(stueckliste_Kroenert, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", &lagerbestand_Kroenert, "KNT")
	STD_Write_Stueckliste("Liste_Kroenert.xlsx", lagerbestand_Kroenert)
	STD_Clean_Lagerbestand(lagerbestand_Kroenert)
	STD_Write_Stueckliste("Liste_Kroenert_Clean.xlsx", lagerbestand_Kroenert)

	STD_Read_Lagerbestand(stueckliste_projekt, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\8000772_Stückliste.xlsx", &Stueckliste, "KNT_Stückliste")
	for _, value := range Stueckliste {
		fmt.Printf("ERP: %-40s", value.ERP)
		fmt.Printf("Artikelnummer: %-50s", value.ArtikelnummerEplan)
		fmt.Printf("Bestellnummer: %-40s", value.Bestellnummer)
		fmt.Printf("Hersteller: %-40s", value.Hersteller)
		fmt.Printf("Stückzahl: %-30f", value.Stueckzahl)
		fmt.Printf("Typ: %-40s", value.Typ)
		fmt.Printf("Beistellung: %-40s", value.Beistellung)
		fmt.Printf("\n")
	}
	STD_Write_Stueckliste("Liste_Stueckliste.xlsx", Stueckliste)
	STD_Clean_Lagerbestand(Stueckliste)
	STD_Write_Stueckliste("Liste_Stueckliste_Clean.xlsx", Stueckliste)

	for _, b := range Stueckliste {
		_, ok := Stueckliste_Map[b.Ort+b.Bestellnummer]

		if ok {
			Stueckliste_Map[b.Ort+b.Bestellnummer].STD_Stueckliste_Update(b.Stueckzahl)
		} else {
			Stueckliste_Map[b.Ort+b.Bestellnummer] = b
		}
	}

	for _, b := range Stueckliste_Map {
		Stueckliste2 = append(Stueckliste2, b)
	}

	STD_Write_Stueckliste("Liste_Stueckliste_Clean_Sum.xlsx", Stueckliste2)

	fmt.Println("Startup finished")

	return true
}
