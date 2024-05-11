package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

func readjson() {

	jsonFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\Test_Project\\blame_SITECA.json")
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	lagerbestand := make(map[string][]ARTIKEL)

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &lagerbestand)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for _, b := range lagerbestand {
		for _, bb := range b {
			fmt.Printf("BMK: %-10s", bb.BMK.FunktionaleZuordnung)
			fmt.Printf(" %-10s", bb.BMK.Funktionskennzeichen)
			fmt.Printf(" %-10s", bb.BMK.Aufstellungsort)
			fmt.Printf(" %-10s", bb.BMK.Ortskennzeichen)
			fmt.Printf(" %-10s", bb.BMK.BMK)
			fmt.Printf("ERP: %-30s", bb.ERP)
			fmt.Printf("Bestellnummer: %-30s", bb.Bestellnummer)
			fmt.Printf("Hersteller: %-30s", bb.Hersteller)
			//fmt.Printf("Beschreibung: %-30s", bb.Beschreibung)
			fmt.Printf("Beistellung: %-30s\n", bb.Beistellung)

		}
	}

}
