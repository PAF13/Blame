package main

import "time"

/*
	Current Imports:
		Lagerstand:
			KNT
			Siteca
			Moeller Electro
		Stueckliste:
			KNT
			Eplan
		Verbindungsliste:
			Eplan

*/

type UUID = string //UUID should show source and time of import data

type VerbindungName = string

type JSON_HEADER struct {
	Name    string
	Version [3]int
	Time    time.Time
	Source  string
}

//non standard generics
type LAGER struct {
	Kunde       string
	Ort         string
	Watengruppe string
}

type BETRIEBSMITELLKENNZEICHEN struct {
	BMKVollständig             string
	BMKidentifizierung         string
	FunktionaleZuordnung       string //==
	Funktionskennzeichen       string //=
	Aufstellungsort            string //++
	Ortskennzeichen            string //+
	Dokumentenart              string //&
	BenutzerdefinierteStruktur string //#
	Anlagennummer              string //empty?
	BMK                        string //-
	Kennbuchstabe              string
}

type VERBINDUNG struct {
	//Info from XML
	UUID                      string
	Quelle                    BETRIEBSMITELLKENNZEICHEN
	Ziel                      BETRIEBSMITELLKENNZEICHEN
	VerbindungZugehörigkeit   string
	Verbindungsquerschnitt    string
	Verbindungsfarbeundnummer string
	VerbindungLänge           string
	Netzname                  string
	Signalname                string
	Potenzialname             string
	Potenzialtyp              string
	Potenzialwert             string
	Netzindex                 string
}
type EXCEL_IMPORT struct {
	Header  JSON_HEADER
	Columns EXCEL_SIMPLE
	Rows    [][]string
}
type ARTIKELLISTE struct {
	Header  JSON_HEADER
	Artikel map[string][]ARTIKEL
}
type ARTIKEL struct {
	UUID               UUID
	BMK                BETRIEBSMITELLKENNZEICHEN
	ERP                string
	ERP_KNT            string
	Bestellnummer      string
	ArtikelnummerEplan string
	Hersteller         string
	Beschreibung       string
	Stueckzahl         float64
	Einheit            string
	Quelle             string
	Stand              string
	Beistellung        string
	Strom              string
	Symbol             string
	Fehler             []string
	Bestellung_Moeller float64
	Bestellung_KNT     float64
	Bestellung_Siteca  float64
}
type EXCEL_SIMPLE struct {
	Header               int
	BMKVollständig       int
	FunktionaleZuordnung int //==
	Funktionskennzeichen int //=
	Aufstellungsort      int //++
	Ortskennzeichen      int //+
	BMK                  int //-
	ERP                  int
	ERP_KNT              int
	Hersteller           int
	Bestellnummer        int
	Bezeichnung          int
	Beschreibung         int
	Stueckzahl           int
	Einheit              int
	Verpackungseinheit   int
	Lagerort             int
	Beistellung          int
}

type PROJEKT struct {
	HEADER               JSON_HEADER
	PROJEKT_NUMMER       string
	PROJEKT_BESCHREIBUNG string
	BAUJAHR              int
	AKTIV                bool
	PRODUKTE             []PRODUKTE
	Produktname          []string
}
