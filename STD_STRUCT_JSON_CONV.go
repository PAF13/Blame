package main

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