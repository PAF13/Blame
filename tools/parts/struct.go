package parts

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
	Symbolname                 string
	Platzierung                string
}

type VERBINDUNG struct {
	//Info from XML
	UUID                                  string
	Quelle                                BETRIEBSMITELLKENNZEICHEN
	Ziel                                  BETRIEBSMITELLKENNZEICHEN
	VerbindungZugehörigkeit               string
	Verbindungsquerschnitt                float64
	Verbindungsfarbeundnummer             string
	VerbindungLänge                       int
	Netzname                              string
	Signalname                            string
	Potenzialname                         string
	Potenzialtyp                          string
	Potenzialwert                         string
	Netzindex                             string
	Funktionsdefinition                   string
	Darstellungsart                       string
	QuelleAnschlussbezeichnungderFunktion string
	ZielAnschlussbezeichnungderFunktion   string
	QuelleKlemmenbezeichnung              string
	ZielKlemmenbezeichnung                string
	QuelleFunktionKatagorie               string
	ZielFunktionKatagorie                 string
	QuelleAnschluss                       string
	ZielAnschluss                         string
}

type ARTIKELLISTE struct {
	Header    JSON_HEADER
	BMK_Liste map[string]string
	Artikel   map[string][]ARTIKEL
}
type DATASOURCE struct {
	KNT     bool
	Siteca  bool
	Moeller bool
	Eplan   bool
}
type ARTIKEL struct {
	UUID               UUID
	Beigestellt        bool
	ERP                string
	ERP_KNT            string
	Bestellnummer      string
	ArtikelnummerEplan string
	Hersteller         string
	Beschreibung       string
	ARTIKELINFO        ARTIKELINFO
	Einheit            string
	Quelle             string
	Stand              string
	Beistellung        string
	Strom              string
	Symbol             string
	Fehler             []string
	Funktionsgruppe    string
	DataSource         DATASOURCE

	EK_KNT    float64
	EK_Siteca float64

	Stueckzahl             float64
	Beistellung_Stueckzahl float64
	Bestellung_Moeller     float64
	Lager_Siteca           float64
	Bestellung_KNT         float64
	Bestellung_Siteca      float64
}
type ARTIKELINFO struct {
	gewerk             string
	produktgruppe      string
	produktuntergruppe string
}
type EPLANGRUPPEN struct {
	gewerk             map[int]string
	produktgruppe      map[int]string
	produktuntergruppe map[int]string
}
type PROJEKT struct {
	HEADER               JSON_HEADER
	PROJEKT_NUMMER       string
	PROJEKT_BESCHREIBUNG string
	BAUJAHR              int
	AKTIV                bool

	Produkte []*PRODUKT
}
type PRODUKT struct {
	Bestellt               bool
	KundeFraigabe          bool
	EPlanFraigabe          bool
	SchaltplanQuelle       string
	BeschriftungFraigabe   bool
	BetriebsmittelFreigabe bool
	Betriebsmittel         map[string]BETRIEBSMITELL
	VerbindungenFreigabe   bool
	Verbindungen           map[string]VERBINDUNG
	NCDatenFraigabe        bool
}
type BETRIEBSMITELL struct {
	BMK BETRIEBSMITELLKENNZEICHEN

	Artikel []*ARTIKEL
}

type FILTER struct {
	FunktionaleZuordnung       bool //==
	Funktionskennzeichen       bool //=
	Aufstellungsort            bool //++
	Ortskennzeichen            bool //+
	Dokumentenart              bool //&
	BenutzerdefinierteStruktur bool //#
	Anlagennummer              bool //empty?
	BMK                        bool //-
	Filter                     map[string]bool
	Product                    map[string]string
}
type BETRIEBSMITELLLISTE struct {
	Filter         map[string]bool
	ProduktDef     FILTER
	Produkte       map[string]*BETRIEBSMITELLKENNZEICHEN
	Betriebsmittel map[string]*BETRIEBSMITELL
}
type LAGERLISTE struct {
	//Betriebsmittel map[string]*ARTIKEL
	Artikel map[string]*ARTIKEL
}
