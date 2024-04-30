package main

import (
	"fmt"
)

type Artikel struct {
	FunktionaleZuordnung       string //==
	Funktionskennzeichen       string //=
	Aufstellungsort            string //++
	Ortskennzeichen            string //+
	Dokumentenart              string //&
	BenutzerdefinierteStruktur string //#
	Anlagennummer              string //empty?
	BMK                        string //-
	ERP                        string
	Bestellnummer              string
	ArtikelnummerEplan         string
	Hersteller                 string
	Typ                        string
	Beschreibung               string
	Stueckzahl                 float64
	Einheit                    string
	Warengruppe                string
	Quelle                     string
	Stand                      string
	Beistellung                string
	Ort                        string
}

func (a *App) Dailyvergleich() {

	stueckliste_Projekt := make(map[string]*Artikel) // Key: Bestellnummer

	fmt.Println("Reading Stueckliste Projekt")
	stueckliste_Projekt = ReadStueckliste(stueckliste_projekt, "\\\\ME-Datenbank-1\\Database\\Schnittstelle\\8000772_Stückliste.xlsx", stueckliste_Projekt, "KNT_Stückliste")
	stueckliste_Projekt = Sitecavergleich(lagerbestand_Siteca, stueckliste_Projekt)

	fmt.Println("Writing Lagerbestand")

	Write_Lager2("lagerbestand_Siteca.xlsx", lagerbestand_Siteca)
	Write_Lager2("lagerbestand_Kroenert.xlsx", lagerbestand_Kroenert)
	Write_Lager2("stueckliste_Projekt.xlsx", stueckliste_Projekt)

	fmt.Println("Writing")
	Write_Stueckliste(lagerbestand_Siteca, lagerbestand_Kroenert, stueckliste_Projekt)

	fmt.Println("Fertig")
}
