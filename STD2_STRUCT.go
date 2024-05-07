package main

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

var lagerbestand map[LAGER][]ARTIKEL // Lagerort

var verbindungsliste map[string]VERBINDUNG

var Betriebsmittel map[BETRIEBSMITELLKENNZEICHEN][50]ARTIKEL


//non standard generics
type LAGER struct {
	Kunde string
	Ort string
	Watengruppe string
}

type BETRIEBSMITELLKENNZEICHEN struct {
	BMKVollständig             string
	FunktionaleZuordnung       string //==
	Funktionskennzeichen       string //=
	Aufstellungsort         string //++
	Ortskennzeichen            string //+
	Dokumentenart              string //&
	BenutzerdefinierteStruktur string //#
	Anlagennummer              string //empty?
	BMK                        string //-
}

type VERBINDUNG struct {
	//Info from XML
	UUID                        string
	Quelle                   	BETRIEBSMITELLKENNZEICHEN
	Ziel 						BETRIEBSMITELLKENNZEICHEN
	VerbindungZugehörigkeit   	string
	Verbindungsquerschnitt    	float64
	Verbindungsfarbeundnummer 	string
	VerbindungLänge           	int
	Netzname                  	string
	Signalname                	string
	Potenzialname             	string
	Potenzialtyp              	string
	Potenzialwert             	string
	Netzindex                 	string
}

type ARTIKEL struct {
	UUID UUID	
	ERP                        string
	Bestellnummer              string
	ArtikelnummerEplan         string
	Hersteller                 string
	Beschreibung               string
	Stueckzahl                 float64
	Einheit                    string
	Quelle                     string
	Stand                      string
	Beistellung                string
}