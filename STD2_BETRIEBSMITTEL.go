package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
)

func (a *App) ImportBetriebsmittel() {
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Betriebsmitttel.xml")
	if err != nil {
		fmt.Println(err)
	}

	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	XML := NewEplanAuswertungXML()

	xml.Unmarshal(byteValue, XML)
	exportArray := &[][]string{}
	exportArray2 := &[]BETRIEBSMITELL{}

	line := XML.Document.Page.Line
	for _, aa := range line {
		betriebsmittel := NewBetriebsmittel()
		artikelPos := NewCounter()
		Bestellnummerpos := NewCounter()

		for _, bb := range aa.Label.Property {
			switch bb.PropertyName {
			case "BMK (vollständig)":
				betriebsmittel.BMK.BMKVollständig = bb.PropertyValue
			case "Name (identifizierend)":
				betriebsmittel.BMK.BMKidentifizierung = bb.PropertyValue
			case "Artikelnummer":
				betriebsmittel.Artikel[*artikelPos].ArtikelnummerEplan = bb.PropertyValue
				*artikelPos++
			case "Bestellnummer":
				betriebsmittel.Artikel[*Bestellnummerpos].Bestellnummer = bb.PropertyValue
				*Bestellnummerpos++
			default:
				//fmt.Printf("Missing | Name: %-50s", bb.PropertyName)
				//fmt.Printf("Value: %-50s", bb.PropertyValue)
				//fmt.Printf("\n")
			}
		}

		betriebsmittel.writeCSVFileBuffer(exportArray)
		*exportArray2 = append(*exportArray2, *betriebsmittel)
	}
	writeJsonFile("EPlan_Steuckliste_Raw", exportArray2)
	//writeCSVFile("EPlan_Steuckliste", exportArray, "EPlanBetriebsmittel")
}
