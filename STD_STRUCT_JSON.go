package main

type JSON_PARSING struct {
	NAME    string
	VERSION [3]int
	BODY    JSON_PROJEKT
}

type JSON_BMK struct {
	BMK_VOLL             string
	BMK_ID               string
	BMK                  string
	FUNKTIONALEZUORDNUNG string
	FUNKTIONSKENNZEICHEN string
	AUFSTELLUNGSORT      string
	ORTSKENNZEICHEN      string
	KENNBUCHSTABE        string
}

type JSON_BAUTEIL struct {
	BMK                 JSON_BMK
	ERP                 string //Nummer für ERP-System
	ERP_QUELLE          string //Quelle der ERPnummer zB Siteca, Kroenert, etc.
	BESTELLNUMMER       string //Herstellernummer
	ARTIKELNUMMER_EPLAN string //Herstellernummer nach EPlan format
	HERSTELLER          string //Herstellername
	STEUCKZAHL          string
	EINHEIT             string
	BEISTELLUNG         string //Bestellungspfligt ob wir bestellen oder kunde beistellung
	GELIEFERT           string
}

type JSON_KABLE struct {
	BMK        JSON_BMK
	VERBINDUNG []JSON_VERBINDUNG
}

type JSON_VERBINDUNG struct {
	ID                       string //Internal ident number. possibly not needed
	BAUTEIL                  [2]JSON_BAUTEIL
	VERBINDUNGZUGEHOERIGKEIT string
	QUERSCHNITT              string
	FARBE_NUMMER             string
	LAENGE                   string
	NETZ_NAME                string
	NETZ_INDEX               string
	SIGNAL_NAME              string
	POTENZIAL_NAME           string
	POTENZIAL_TYP            string
	POTENZIAL_WERT           string
	FEHLER                   []string
}

type JSON_PRODUKTEINFO_KONSTRUKTION struct {
	FREIGABE_KUNDE        string
	FREIGABE_KONSTRUKTION string
	EPLAN_PROJEKT         string
	SCHALTPLAN_QUELLE     string
	NC_DATEN              string
	BESCHRIFTUNG          string
	STUECKLISTE           string
	VERBINDUNGSLISTE      string
	KLEMMENAUFREIHPLAN    string
}
type JSON_PRODUKTEINFO_EINKAUF struct {
}
type JSON_PRODUKTEINFO_ARBEITSVORBEREITUNG struct {
}
type JSON_PRODUKTEINFO_FERTIGUNG_MECHANISCH struct {
}
type JSON_PRODUKTEINFO_FERTIGUNG_ELEKTRISCH struct {
}
type JSON_PRODUKTEINFO_FERTIGUNG_ENDMONTAGE struct {
}
type JSON_PRODUKTEINFO_ADMINISTRATION struct {
}
type JSON_PRODUKTEINFO_ZUSATZINFO struct {
}
type JSON_PRODUKTE struct { //produkt info schaltschrank etc
	BMK                    JSON_BMK
	LIEFERTERMIN           string
	LIEFERTERMIN_BAUSTELLE string
	PRIORITAET             string
	KONSTRUKTION           JSON_PRODUKTEINFO_KONSTRUKTION
	EINKAUF                JSON_PRODUKTEINFO_EINKAUF
	ARBEITSVORBEREITUNG    JSON_PRODUKTEINFO_ARBEITSVORBEREITUNG
	FERTIGUNG_MECHANISCH   JSON_PRODUKTEINFO_FERTIGUNG_MECHANISCH
	FERTIGUNG_ELEKTRISCH   JSON_PRODUKTEINFO_FERTIGUNG_ELEKTRISCH
	FERTIGUNG_ENDMONTAGE   JSON_PRODUKTEINFO_FERTIGUNG_ENDMONTAGE
	ADMINISTRATION         JSON_PRODUKTEINFO_ADMINISTRATION
	ZUSATZINFO             JSON_PRODUKTEINFO_ZUSATZINFO

	ARTIKEL []JSON_BAUTEIL
	DRAEHTE []JSON_VERBINDUNG
}

type JSON_PROJEKT struct {
	PROJEKT_NUMMER       string
	PROJEKT_BESCHREIBUNG string
	BAUJAHR              string
	AKTIV                string
	PRODUKTE             []JSON_PRODUKTE
}
