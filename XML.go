package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

func (a *App) XMLTest() {
	artikelstammdaten := make(map[string]*Part) // Key: Bestellnummer

	// Open our xmlFile
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlan_Artikelstammdaten.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)
	// we initialize our Users array
	var partManager PartManager
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &partManager)

	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for i := 0; i < len(partManager.Parts); i++ {
		if partManager.Parts[i].Bestellnummer != "" {
			artikelstammdaten[partManager.Parts[i].Bestellnummer] = &Part{
				ERP:           partManager.Parts[i].ERP,
				Artikelnummer: partManager.Parts[i].Artikelnummer,
				Bestellnummer: partManager.Parts[i].Bestellnummer,
				Hersteller:    partManager.Parts[i].Hersteller,
				Typ:           partManager.Parts[i].Typ,
				Note:          partManager.Parts[i].Note,
				//Note:          partManager.Parts[i].Note,
			}
		}
	}
}
