package main

/*
func INIT_ARTIKELSTAMMDATEN() {

	STD_Read_Lagerbestand2("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Topix_Artikel20240502.xlsx", "SITECA")
	STD_Read_Lagerbestand2("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Kopie von Lagerhueter_26_04_2024.xlsx", "KNT")
}

func STD_Read_Lagerbestand2(pfad string, kunde string) {
	log.Println("Reading " + kunde + " Lagerbestand")
	readNum := externalReadExcel[kunde].KUNDE_EINSTELLUNG.KUNDE_LAGERBESTAND.ARTIKEL
	lagerOrt := externalReadExcel[kunde].KUNDE_EINSTELLUNG.KUNDE_LAGERBESTAND
	lager := LAGER{
		EIGENTUEMER: kunde,
		LAGERORT: make(map[string]LAGERORT),
	}

	spreadsheet, err := excelize.OpenFile(pfad)
	if err != nil {
		fmt.Println(err)
	}

	rows, err := spreadsheet.GetRows(spreadsheet.GetSheetList()[0])
	if err != nil {
		fmt.Println(err)
	}
	if err := spreadsheet.Close(); err != nil {
		log.Println(err)
	}
	artikel_Success := 0
	for num,row := range rows{
		if externalReadExcel[kunde].KUNDE_EINSTELLUNG.KUNDE_LAGERBESTAND.FIRST_VALUE -1 < num {
			_,ok := lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)]
			if !ok{
				lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)] =  LAGERORT{
					ARTIKELANZAHL: 0,
					BAUTEIL: make(map[string]BAUTEIL),
				}
			}

			_,artikel_ok := lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)].BAUTEIL[safeStringArrayPull(row,readNum.BESTELLNUMMER)]
			if !artikel_ok {
				log.Println("new")
				artikel_Success++
				lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)].BAUTEIL[safeStringArrayPull(row,readNum.BESTELLNUMMER)] = lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)].BAUTEIL[safeStringArrayPull(row,readNum.BESTELLNUMMER)].STD_Read_Lagerbestand3(row,readNum,kunde)


			}else{
				log.Println("dup")
				artikel_Success++
				lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)].BAUTEIL["[DUPLICATE " + fmt.Sprintf("%d",artikel_Success) + "]"+ safeStringArrayPull(row,readNum.BESTELLNUMMER)] = lager.LAGERORT[safeStringArrayPull(row,lagerOrt.LAGERORT)].BAUTEIL["[DUPLICATE " + fmt.Sprintf("%d",artikel_Success) + "]"+ safeStringArrayPull(row,readNum.BESTELLNUMMER)].STD_Read_Lagerbestand3(row,readNum,kunde)
			}

			log.Println(artikel_Success)
		}

	}

	json, err := json.MarshalIndent(lager, "", "    ")
	if err != nil {
		log.Println(err)
	}

	err = os.WriteFile("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\Lager_"+kunde+".json", json, 0644)
	if err != nil {
		log.Println(err)
	}
}
func (lager2 BAUTEIL) STD_Read_Lagerbestand3(artikel []string, readNum ARTIKEL, kunde string) BAUTEIL{

	stueckzahl := 0.0
	var err error
	if safeStringArrayPull(artikel,int(readNum.STEUCKZAHL)) != ""{
		stueckzahl,err = strconv.ParseFloat(safeStringArrayPull(artikel,readNum.STEUCKZAHL),64)
		if err != nil {
			log.Println(err)
		}
	}
	lager2.ERP = safeStringArrayPull(artikel,readNum.ERP)
	lager2.ERP_QUELLE = kunde
	lager2.BESTELLNUMMER = strings.ToUpper(bestellnummerClean(safeStringArrayPull(artikel,readNum.BESTELLNUMMER)))
	lager2.HERSTELLER = safeStringArrayPull(artikel,readNum.HERSTELLER)
	lager2.ARTIKELNUMMER_EPLAN = safeStringArrayPull(artikel,readNum.ARTIKELNUMMER_EPLAN)
	lager2.BESCHREIBUNG = safeStringArrayPull(artikel,24)
	lager2.STEUCKZAHL = stueckzahl
	lager2.EINHEIT = safeStringArrayPull(artikel,readNum.EINHEIT)

	return lager2
}

func bestellnummerClean (x string) string{
	x = strings.ReplaceAll(x, " ", "")
	x = strings.ReplaceAll(x, "\t", "")
	x = strings.ReplaceAll(x, "\n", "")
	return x
}*/