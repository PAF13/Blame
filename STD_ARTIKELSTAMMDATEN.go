package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/xuri/excelize/v2"
)

func INIT_ARTIKELSTAMMDATEN() {

	STD_Read_Lagerbestand2("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240502.xlsx", "SITECA")
	STD_Read_Lagerbestand2("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", "KNT")
}

func STD_Read_Lagerbestand2(pfad string, kunde string) {
	lager := LAGER{
		EIGENTUEMER: kunde,
		LAGERORT: make(map[string]LAGERORT),
	}
	
	fmt.Println("Opening: " + pfad)
	spreadsheet, err := excelize.OpenFile(pfad)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}
	for num,b := range rows{
		lager.STD_Read_Lagerbestand3(b, kunde)
		fmt.Println("Number of items: " + fmt.Sprintf("%d", num))
	}
	b2, err := json.MarshalIndent(lager, "", "    ")
	if err != nil {
		log.Println(err)
	}
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\Lager_"+kunde+".json", b2, 0644)
	if err != nil {
		log.Println(err)
	}
	if err := spreadsheet.Close(); err != nil {
		log.Println(err)
	}
}
func (lager *LAGER) STD_Read_Lagerbestand3(artikel []string, kunde string){
	readNum := externalReadExcel[kunde].KUNDE_EINSTELLUNG.KUNDE_LAGERBESTAND.ARTIKEL
	lagerort,ok := lager.LAGERORT[kunde]
	if !ok{
		lager.LAGERORT[kunde] =  LAGERORT{
			LAGERNAME: "Lager1",
			ARTIKELANZAHL: 0,
			BAUTEIL: make(map[string]BAUTEIL),
		}
	}
	if safeStringArrayPull(artikel,readNum.STEUCKZAHL) != ""{
		stueckzahl,err := strconv.ParseFloat(safeStringArrayPull(artikel,readNum.STEUCKZAHL),64)
		if err != nil {
			log.Println(err)
		}
		//fmt.Println(artikel,readNum.STEUCKZAHL)
	lagerort.ARTIKELANZAHL++
	lager.LAGERORT[kunde] = lagerort
	//log.Println(lagerort.ARTIKELANZAHL)
	lager.LAGERORT[kunde].BAUTEIL[strings.ToUpper(bestellnummerClean(safeStringArrayPull(artikel,readNum.UID))) + fmt.Sprintf("%d", lagerort.ARTIKELANZAHL)] = BAUTEIL{
		BMK: BMK2{},
		ERP: safeStringArrayPull(artikel,readNum.ERP),
		ERP_QUELLE: kunde,
		BESTELLNUMMER: strings.ToUpper(bestellnummerClean(safeStringArrayPull(artikel,readNum.BESTELLNUMMER))),
		HERSTELLER: safeStringArrayPull(artikel,readNum.HERSTELLER),
		ARTIKELNUMMER_EPLAN: safeStringArrayPull(artikel,readNum.ARTIKELNUMMER_EPLAN),
		BESCHREIBUNG: safeStringArrayPull(artikel,24),
		STEUCKZAHL: stueckzahl,
		EINHEIT: safeStringArrayPull(artikel,readNum.EINHEIT),
	}

	}	
}

func bestellnummerClean (x string) string{
	//x = strings.ReplaceAll(x, " ", "")
	//x = strings.ReplaceAll(x, "\t", "")
	//x = strings.ReplaceAll(x, "\n", "")
	return x
}