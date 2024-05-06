package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/xuri/excelize/v2"
)

func INIT_ARTIKELSTAMMDATEN() {

	STD_Read_Lagerbestand2("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240502.xlsx", "SITECA")
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
	for _,b := range rows{
		lager.STD_Read_Lagerbestand3(b, kunde)
		fmt.Println(b)
	}
	b2, err := json.MarshalIndent(lager, "", "    ")
	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(b2))
	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\Lager_"+kunde+".json", b2, 0644)
	if err != nil {
		log.Println(err)
	}
	if err := spreadsheet.Close(); err != nil {
		log.Println(err)
	}
}
func (lager *LAGER) STD_Read_Lagerbestand3(artikel []string, kunde string){
	_,ok := lager.LAGERORT[kunde]
	if !ok{
		lager.LAGERORT[kunde] =  LAGERORT{
			LAGERNAME: "Lager1",
			BAUTEIL: make(map[string]BAUTEIL),
		}
	}
	if safeStringArrayPull(artikel,72) != ""{
		stueckzahl,err := strconv.ParseFloat(safeStringArrayPull(artikel,50),32)
	if err != nil {
		log.Println(err)
	}
		lager.LAGERORT[kunde].BAUTEIL[safeStringArrayPull(artikel,72)] = BAUTEIL{
			BMK: BMK2{},
			ERP: safeStringArrayPull(artikel,2),
			ERP_QUELLE: kunde,
			BESTELLNUMMER: safeStringArrayPull(artikel,72),
			HERSTELLER: safeStringArrayPull(artikel,6),
			ARTIKELNUMMER_EPLAN: safeStringArrayPull(artikel,187),
			BESCHREIBUNG: safeStringArrayPull(artikel,24),
			STEUCKZAHL: stueckzahl,
			EINHEIT: safeStringArrayPull(artikel,12),
		}
	}	
}