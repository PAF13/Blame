package main

import "time"

func setHeader(fileType string, source string) JSON_HEADER {
	return JSON_HEADER{
		Name:    fileType,
		Version: [3]int{0, 0, 1},
		Time:    time.Now(),
		Source:  source,
	}

}
func setImportColumns(kunde string, maxSize1 int) EXCEL_SIMPLE {
	maxSize := maxSize1 - 1
	var columnNumber EXCEL_SIMPLE
	switch kunde {
	case "SITECA":
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
	case "KNT":
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
	case "stueckliste":
		columnNumber = EXCEL_SIMPLE{
			Header:               7,
			BMKVollständig:       maxSize,
			FunktionaleZuordnung: 0,
			Funktionskennzeichen: 1,
			Aufstellungsort:      2,
			Ortskennzeichen:      3,
			BMK:                  4,
			ERP:                  maxSize,
			ERP_KNT:              7,
			Bestellnummer:        9,
			Bezeichnung:          maxSize,
			Beschreibung:         10,
			Stueckzahl:           5,
			Einheit:              maxSize,
			Verpackungseinheit:   maxSize,
			Lagerort:             maxSize,
			Hersteller:           11,
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

func NewExcelImport(kunde string, maxSize int, fileType string, source string) *EXCEL_IMPORT {
	return &EXCEL_IMPORT{
		Header:  setHeader(fileType, source),
		Columns: setImportColumns(kunde, maxSize),
		Rows:    make([][]string, 0),
	}
}
