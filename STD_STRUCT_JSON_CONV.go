package main

var stueckliste_Topix = Stuckliste_ImportTemplate{
	Aufstellungsort:    300,
	Ortskennzeichen:    300,
	ERP:                2,
	Bestellnummer:      72,
	Hersteller:         6,
	ArtikelnummerEplan: 187,
	Beschreibung:       24,
	Stueckzahl:         50,
	Einheit:            12,
	FirstValue:         3,
	Warengruppe:        300,
	Beistellung:        300,
	Ort:                300,
	Herstellertyp:      188,
	HerstellerEplan:    300,
	Bestellnr_L1:       272,
	Bezeichnung:        11,
}
var stueckliste_Kroenert = Stuckliste_ImportTemplate{
	Aufstellungsort:    300,
	Ortskennzeichen:    300,
	ERP:                1,
	Bestellnummer:      13,
	Hersteller:         20,
	ArtikelnummerEplan: 20,
	Beschreibung:       9,
	Stueckzahl:         3,
	Einheit:            20,
	FirstValue:         4,
	Warengruppe:        20,
	Beistellung:        12,
	Ort:                3,
	Herstellertyp:      300,
	HerstellerEplan:    300,
	Bestellnr_L1:       300,
	Bezeichnung:        300,
}
var stueckliste_projekt = Stuckliste_ImportTemplate{
	Aufstellungsort:    2,
	Ortskennzeichen:    3,
	ERP:                7,
	Bestellnummer:      9,
	Hersteller:         11,
	ArtikelnummerEplan: 20,
	Beschreibung:       10,
	Stueckzahl:         5,
	Einheit:            20,
	FirstValue:         7,
	Warengruppe:        20,
	Beistellung:        12,
	Ort:                20,
	Herstellertyp:      300,
	HerstellerEplan:    300,
	Bestellnr_L1:       300,
	Bezeichnung:        300,
}
var externalReadExcel map[string]EXTERN_READ_EXCEL

type EINSTELLUNGEN struct {
	ALLEMEIN_EINSTELLUNG ALLGEMEIN_EINSTELLUNG
	KUNDE_EINSTELLUNG    map[string]KUNDE_EINSTELLUNG
}
type ALLGEMEIN_EINSTELLUNG struct {
	PROJEKTORDNER PROJEKTORDNER
	STD_PFAD      STD_PFAD
}
type STD_PFAD struct {
	ROOT string
}
type PROJEKTORDNER struct {
	ROOT            string
	DOKUMENTE       DOKUMENTE
	SCHALTPLAN      SCHALTPLAN
	MATERIAL        MATERIAL
	FERTIGUNGSDATEN FERTIGUNGSDATEN
	FOTOS           FOTOS
	PRUEFPROTOKOLL  PRUEFPROTOKOLL
}
type DOKUMENTE struct {
	ROOT string
}
type SCHALTPLAN struct {
	ROOT               string
	KUNDENAUSFERTIGUNG string
}
type MATERIAL struct {
	ROOT                       string
	STUECKLISTE_KUNDENFREIGABE string
	STUECKLISTE_INTERN         string
	STUECKLISTE_TOPIX          string
}
type FERTIGUNGSDATEN struct {
	ROOT         string
	NC           string
	DRAHT        string
	BESCHRIFTUNG string
	UNTERLAGEN   string
}
type FOTOS struct {
	ROOT        string
	BEISTELLUNG string
	PRODUKTION  string
}
type PRUEFPROTOKOLL struct {
	ROOT string
}

type KUNDE_BMK struct {
	BMK_VOLL                   int
	BMK_ID                     int
	FUNKTIONALEZUORDNUNG       int
	FUNKTIONSKENNZEICHEN       int
	AUFSTELLUNGSORT            int
	ORTSKENNZEICHEN            int
	BMK                        int
	DOKUMENTENART              int
	BENUTZERDEFINIERTESTRUKTUR int
	ANLAGENNUMMER              int
	KENNBUCHSTABE              int
}