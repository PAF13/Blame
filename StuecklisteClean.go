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
	Herstellertyp              string //
	HerstellerEplan            string //
	Bestellnr_L1               string //
	Bezeichnung                string //
}

func (a *App) Dailyvergleich() {

	fmt.Println("Fertig")
}
