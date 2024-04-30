package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

// XML parsing
type PartManager struct {
	XMLName xml.Name `xml:"partsmanagement"`
	Parts   []Part   `xml:"part"`
}

type Part struct {
	Artikelnummer string `xml:"P_ARTICLE_PARTNR,attr"`       //PXC.1088136
	ERP           string `xml:"P_ARTICLE_ERPNR,attr"`        //1005928
	Bestellnummer string `xml:"P_ARTICLE_ORDERNR,attr"`      //1088136
	Hersteller    string `xml:"P_ARTICLE_MANUFACTURER,attr"` //PXC
	Typ           string `xml:"P_ARTICLE_TYPENR,attr"`       //AXL F LPSDO8/3 1F
	Note          string `xml:"P_ARTICLE_NOTE,attr"`         //en_US@Motor circuit breaker, TeSys Deca, 3P, 1.6-2.5 A, thermal magnetic, screw clamp terminals
}

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
