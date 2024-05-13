package main

type BMK2 struct {
	BMK_VOLL             string
	BMK_ID               string
	BMK                  string
	FUNKTIONALEZUORDNUNG string
	FUNKTIONSKENNZEICHEN string
	AUFSTELLUNGSORT      string
	ORTSKENNZEICHEN      string
	KENNBUCHSTABE        string
}

type BAUTEIL struct {
	BMK                 BMK2
	ERP                 string //Nummer für ERP-System
	ERP_QUELLE          string //Quelle der ERPnummer zB Siteca, Kroenert, etc.
	BESTELLNUMMER       string //Herstellernummer
	ARTIKELNUMMER_EPLAN string //Herstellernummer nach EPlan format
	BESCHREIBUNG        string
	HERSTELLER          string //Herstellername
	STEUCKZAHL          float64
	EINHEIT             string
	BEISTELLUNG         string //Bestellungspfligt ob wir bestellen oder kunde beistellung
	GELIEFERT           string
}

type KABLE struct {
	BMK        BMK2
	VERBINDUNG []VERBINDUNG
}

type PRODUKTEINFO_KONSTRUKTION struct {
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
type PRODUKTEINFO_EINKAUF struct {
}
type PRODUKTEINFO_ARBEITSVORBEREITUNG struct {
}
type PRODUKTEINFO_FERTIGUNG_MECHANISCH struct {
}
type PRODUKTEINFO_FERTIGUNG_ELEKTRISCH struct {
}
type PRODUKTEINFO_FERTIGUNG_ENDMONTAGE struct {
}
type PRODUKTEINFO_ADMINISTRATION struct {
}
type PRODUKTEINFO_ZUSATZINFO struct {
}
type PRODUKTE struct { //produkt info schaltschrank etc
	BMK                    BMK2
	LIEFERTERMIN           string
	LIEFERTERMIN_BAUSTELLE string
	PRIORITAET             string
	KONSTRUKTION           PRODUKTEINFO_KONSTRUKTION
	EINKAUF                PRODUKTEINFO_EINKAUF
	ARBEITSVORBEREITUNG    PRODUKTEINFO_ARBEITSVORBEREITUNG
	FERTIGUNG_MECHANISCH   PRODUKTEINFO_FERTIGUNG_MECHANISCH
	FERTIGUNG_ELEKTRISCH   PRODUKTEINFO_FERTIGUNG_ELEKTRISCH
	FERTIGUNG_ENDMONTAGE   PRODUKTEINFO_FERTIGUNG_ENDMONTAGE
	ADMINISTRATION         PRODUKTEINFO_ADMINISTRATION
	ZUSATZINFO             PRODUKTEINFO_ZUSATZINFO
	ARTIKEL                []BAUTEIL
	DRAEHTE                []VERBINDUNG
}

type LAGERORT struct {
	LAGERNAME     string
	ARTIKELANZAHL int
	BAUTEIL       map[string]BAUTEIL
}

type EXTERN_READ_EXCEL struct {
	HEADER            JSON_HEADER
	KUNDE_EINSTELLUNG KUNDE_EINSTELLUNG
}

type KUNDE_EINSTELLUNG struct {
	KUNDE_STUECKLISTE  KUNDE_STUECKLISTE
	KUNDE_LAGERBESTAND KUNDE_LAGERBESTAND
}
type KUNDE_STUECKLISTE struct {
	FIRST_VALUE int
	ARTIKEL     ARTIKEL
	KUNDE_BMK   KUNDE_BMK
}

type KUNDE_LAGERBESTAND struct {
	LAGERORT    int
	FIRST_VALUE int
	ARTIKEL     ARTIKEL
}
type ARTIKEL2 struct {
	UID                 int
	ERP                 int
	ERP_QUELLE          string
	BESTELLNUMMER       int
	ARTIKELNUMMER_EPLAN int
	HERSTELLER          int
	STEUCKZAHL          int
	EINHEIT             int
	BEISTELLUNG         int
	BESCHREIBUNG        int
	GELIEFERT           int
}
