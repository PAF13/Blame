package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func (a *App) Verbindungsliste(verbindungsliste string) {

	//Excel Collum names
	//quelleBMKVollmitAnschluss:="";
	//quelleBMKVoll:="";
	//quellefunktionaleZuordnung:="";
	//quellefunktionskennzeichen:="";
	//quelleaufstellungsort:="";
	//quelleortskennzeichen:="";
	//quelleBMK:="";
	//quelleBMKKennbuchstabe:="";
	//quellefunktionstext:="";
	//quelletechnischeKenngroesse:="";
	//quellefunkdefKategorie:="";
	//quellefunkdefGruppe:="";
	//quellefunkdefBeschreibung:="";
	//quelleanschlussbezeichnung:="";
	//quellefunkdef:="";
	//quellesymbolname:="";
	//quellesymbolvariante:="";

	//ZielBMKVollmitAnschluss:="";
	//ZielBMKVoll:="";
	//ZielfunktionaleZuordnung:="";
	//Zielfunktionskennzeichen:="";
	//Zielaufstellungsort:="";
	//Zielortskennzeichen:="";
	//ZielBMK:="";
	//ZielBMKKennbuchstabe:="";
	//Zielfunktionstext:="";
	//ZieltechnischeKenngroesse:="";
	//ZielfunkdefKategorie:="";
	//ZielfunkdefGruppe:="";
	//ZielfunkdefBeschreibung:="";
	//Zielanschlussbezeichnung:="";
	//Zielfunkdef:="";
	//Zielsymbolname:="";
	//Zielsymbolvariante:="";

	//verbindungBMKVoll:="";
	//verbindungzugehoerigkeit:="";
	//verbindungfunkdef:="";
	//verbindungdurchmesser:="";
	//verbindungfarbe:="";
	//verbindunglaenge:="";
	//verbindungnetzname:="";
	//verbindungsignalname:="";
	//verbindungpotenzialname:="";
	//verbindungpotenzialtyp:="";
	//verbindungpotenzialwert:="";
	//verbindungnetzindex:="";

	spreadsheet, err := excelize.OpenFile(verbindungsliste)
	if err != nil {
		fmt.Println(err)
	}
	defer func() {
		// Close the spreadsheet.
		if err := spreadsheet.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}

	type Verbindung struct {
		ID             int
		verbindungen   [2]string
		ort            [2]string
		farbe          string
		querschnitt    int
		laenge         int
		funkdef        string
		zugehoerigkeit string
	}
	verbindungMap := make(map[string][]Verbindung)
	bauteilBMK := make(map[string][]string)

	type Pfad struct {
		verbindungen []Verbindung
	}
	fmt.Println(len(rows))
	for line, row := range rows {
		fmt.Println("line: ", line, "row: ", row)

		verbindung := Verbindung{ID: line, verbindungen: [2]string{safeArrayPull(row, 0), safeArrayPull(row, 17)}, ort: [2]string{safeArrayPull(row, 1), safeArrayPull(row, 18)}, farbe: safeArrayPull(row, 38), querschnitt: safeIntConvert(row, 37), laenge: safeIntConvert(row, 39), funkdef: row[35], zugehoerigkeit: row[36]}

		verbindungMap[safeArrayPull(row, 0)] = append(verbindungMap[safeArrayPull(row, 0)], verbindung)
		verbindungMap[safeArrayPull(row, 17)] = append(verbindungMap[safeArrayPull(row, 17)], verbindung)

		bauteilBMK[safeArrayPull(row, 1)] = append(bauteilBMK[safeArrayPull(row, 1)], safeArrayPull(row, 0))
		//fmt.Println(safeArrayPull(row, 0))
		//fmt.Println(bauteilBMK)
	}
	//for line, row := range bauteilBMK {
	//	fmt.Println("BMK:", line, "Anschluesse:", row)
	//}
}

func duplicateAnschluss(s []string, ss string) bool {
	for _, anschluss := range s {
		//fmt.Println("line:", line, "row:", row)
		if anschluss != ss {
			fmt.Println("1", anschluss)
			fmt.Println("2", ss)
			return true
		}

	}
	return false
}

func safeIntConvert(r []string, n int) int {
	if len(r) > n {
		if r[n] == "" {
			return 0
		}
		querschnittClean := strings.Replace(r[n], ",", "", -1)
		querschnitt, _ := strconv.Atoi(querschnittClean)
		return querschnitt
	}
	return 0
}

func safeArrayPull(r []string, n int) string {
	if len(r) > n {
		if r[n] == "" {
			return ""
		}
		return r[n]
	}
	return ""
}

/*func runner(a string){
		if !duplicateAnschluss(bauteilBMK[safeArrayPull(row, 1)], safeArrayPull(row, 17)) {
			bauteilBMK[safeArrayPull(row, 18)] = append(bauteilBMK[safeArrayPull(row, 18)], safeArrayPull(row, 17))
		}
	runner()
}*/
