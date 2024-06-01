package main

import "time"

func setHeader(fileType string, fileName string) JSON_HEADER {
	return JSON_HEADER{
		Name:    fileType,
		Version: [3]int{0, 0, 1},
		Time:    time.Now(),
		Source:  fileName,
	}

}
func setImportColumns(kunde string, maxSize1 int, fileType string) EXCEL_SIMPLE {
	maxSize := maxSize1 - 1
	var columnNumber EXCEL_SIMPLE
	switch {
	case kunde == "SITECA" && fileType == "Lager":
		columnNumber =
			EXCEL_SIMPLE{
				Header:               3,
				BMKVollständig:       maxSize,
				FunktionaleZuordnung: maxSize,
				Funktionskennzeichen: maxSize,
				Aufstellungsort:      maxSize,
				Ortskennzeichen:      maxSize,
				BMK:                  maxSize,
				ERP:                  2,
				ERP_KNT:              maxSize,
				Bestellnummer:        72,
				Bezeichnung:          maxSize,
				Beschreibung:         24,
				Stueckzahl:           5,
				Einheit:              maxSize,
				Verpackungseinheit:   maxSize,
				Lagerort:             maxSize,
				Hersteller:           6,
				Beistellung:          maxSize,
			}
	case kunde == "KNT" && fileType == "Lager":
		columnNumber = EXCEL_SIMPLE{
			Header:               4,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: maxSize,
			Funktionskennzeichen: maxSize,
			Aufstellungsort:      maxSize,
			Ortskennzeichen:      maxSize,
			BMK:                  maxSize,
			ERP:                  maxSize,
			ERP_KNT:              1,
			Bestellnummer:        13,
			Bezeichnung:          maxSize,
			Beschreibung:         11,
			Stueckzahl:           3,
			Einheit:              maxSize,
			Verpackungseinheit:   maxSize,
			Lagerort:             maxSize,
			Hersteller:           maxSize,
			Beistellung:          maxSize,
		}
	case kunde == "MOELLER" && fileType == "Lager":
		columnNumber = EXCEL_SIMPLE{
			Header:               1,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: maxSize,
			Funktionskennzeichen: maxSize,
			Aufstellungsort:      maxSize,
			Ortskennzeichen:      maxSize,
			BMK:                  maxSize,
			ERP:                  5,
			Bestellnummer:        2,
			Bezeichnung:          1,
			Beschreibung:         maxSize,
			Stueckzahl:           4,
			Einheit:              maxSize,
			Verpackungseinheit:   maxSize,
			Lagerort:             maxSize,
			Hersteller:           0,
			Beistellung:          maxSize,
		}
	case kunde == "TOPIX" && fileType == "stueckliste":
		columnNumber = EXCEL_SIMPLE{
			Header:               1,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: maxSize,
			Funktionskennzeichen: maxSize,
			Aufstellungsort:      maxSize,
			Ortskennzeichen:      maxSize,
			BMK:                  maxSize,
			ERP:                  83,
			ERP_KNT:              maxSize,
			Bestellnummer:        84,
			Bezeichnung:          maxSize,
			Beschreibung:         85,
			Stueckzahl:           80,
			Einheit:              81,
			Verpackungseinheit:   81,
			Lagerort:             maxSize,
			Hersteller:           11,
			Beistellung:          maxSize,
		}
	case kunde == "KNT" && fileType == "stueckliste":
		columnNumber = EXCEL_SIMPLE{
			Header:               7,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: 0,
			Funktionskennzeichen: 1,
			Aufstellungsort:      2,
			Ortskennzeichen:      3,
			BMK:                  4,
			ERP:                  maxSize,
			ERP_KNT:              5,
			Bestellnummer:        7,
			Bezeichnung:          maxSize,
			Beschreibung:         10,
			Stueckzahl:           6,
			Einheit:              maxSize,
			Verpackungseinheit:   maxSize,
			Lagerort:             maxSize,
			Hersteller:           8,
			Beistellung:          12,
		}

	default:
		columnNumber = EXCEL_SIMPLE{
			Header:               maxSize,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: maxSize,
			Funktionskennzeichen: maxSize,
			Aufstellungsort:      maxSize,
			Ortskennzeichen:      maxSize,
			BMK:                  maxSize,
			ERP:                  maxSize,
			Bestellnummer:        maxSize,
			Bezeichnung:          maxSize,
			Beschreibung:         maxSize,
			Stueckzahl:           maxSize,
			Einheit:              maxSize,
			Verpackungseinheit:   maxSize,
			Lagerort:             maxSize,
			Hersteller:           maxSize,
			Beistellung:          maxSize,
		}
	}
	return columnNumber
}

func NewArtikelliste(fileType string, source string) *ARTIKELLISTE {
	return &ARTIKELLISTE{
		Header:    setHeader(fileType, source),
		BMK_Liste: map[string]string{},
		Artikel:   map[string][]ARTIKEL{},
	}
}

func NewExcelImport(kunde string, maxSize int, fileType string, fileName string) *EXCEL_IMPORT {
	return &EXCEL_IMPORT{
		Header:  setHeader(fileType, fileName),
		Columns: setImportColumns(kunde, maxSize, fileType),
		Rows:    make([][]string, 0),
	}
}

func NewProjekt(fileType string, source string) *PROJEKT {
	return &PROJEKT{
		HEADER:               setHeader(fileType, source),
		PROJEKT_NUMMER:       "",
		PROJEKT_BESCHREIBUNG: "",
		BAUJAHR:              2024,
		AKTIV:                true,
		Produkte:             nil,
	}
}
func NewVerbindungMap() map[string]VERBINDUNG {
	return make(map[string]VERBINDUNG)
}
func NewVerbindungArray() *[]VERBINDUNG {
	return &[]VERBINDUNG{}
}
func NewVerbindung() *VERBINDUNG {
	return &VERBINDUNG{}
}
func NewVerbindung1() map[string]VERBINDUNG {
	return make(map[string]VERBINDUNG)
}
func NewProdukt() *PRODUKT {
	return &PRODUKT{
		Bestellt:               false,
		KundeFraigabe:          false,
		EPlanFraigabe:          false,
		SchaltplanQuelle:       "",
		BeschriftungFraigabe:   false,
		BetriebsmittelFreigabe: false,
		Betriebsmittel:         make(map[string]BETRIEBSMITELL),
		VerbindungenFreigabe:   false,
		Verbindungen:           make(map[string]VERBINDUNG),
		NCDatenFraigabe:        false,
	}
}

func NewEplanAuswertungXML() *EplanAuswertungXML {
	return &EplanAuswertungXML{}
}

func NewBool() *bool {
	b := false
	return &b
}

func NewBetriebsmittel() *BETRIEBSMITELL {
	return &BETRIEBSMITELL{
		Artikel: [50]ARTIKEL{},
	}
}

func NewCounter() *int {
	d := 0
	return &d
}

func newArtikel(url string) *ArtikelSeedRaw {
	return &ArtikelSeedRaw{
		URL:           url,
		Artikelnummer: []string{},
		Bestellnummer: []string{},
	}
}
