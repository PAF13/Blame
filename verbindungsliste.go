package main

import (
	"encoding/xml"
	"fmt"
	"io"
	"os"
)

type EplanLabelling struct {
	XMLName  xml.Name `xml:"EplanLabelling"`
	Id       string   `xml:"source_id,attr"`
	Document Document `xml:"Document"`
}
type Document struct {
	XMLName xml.Name `xml:"Document"`
	Id      string   `xml:"source_id,attr"`
	Page    Page     `xml:"Page"`
}
type Page struct {
	XMLName xml.Name `xml:"Page"`
	Id      string   `xml:"source_id,attr"`
	Lines   []Line   `xml:"Line"`
}

type Line struct {
	XMLName xml.Name `xml:"Line"`
	Id      string   `xml:"source_id,attr"`
	Labels  []Label  `xml:"Label"`
}
type Label struct {
	XMLName    xml.Name   `xml:"Label"`
	Id         string     `xml:"source_id,attr"`
	Properties []Property `xml:"Property"`
}

type Property struct {
	XMLName       xml.Name `xml:"Property"`
	PropertyName  string   `xml:"PropertyName"`
	PropertyValue string   `xml:"PropertyValue"`
}
type VerbindungProperty struct {
	PropertyName  string
	PropertyValue string
}

/*
Name des Zielanschlusses (vollständig)
BMK (identifizierend)
Funktionale Zuordnung
Funktionskennzeichen
Aufstellungsort
Ortskennzeichen
BMK (identifizierend, ohne Projektstrukturen)
BMK: Kennbuchstabe
Funktionstext
Technische Kenngrößen
Funktionsdefinition: Kategorie
Funktionsdefinition: Gruppe
Funktionsdefinition: Beschreibung
Anschlussbezeichnung der Funktion
Funktionsdefinition
Symbolname
Symbolvariante
*/

func (a *App) VerbindungRead() {

	// Open our xmlFile
	xmlFile, err := os.Open("\\\\ME-Datenbank-1\\Database\\Schnittstelle\\EPlanOutput\\EPlan_Klemmen.xml")
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened EPlan_Klemmen.xml")
	// defer the closing of our xmlFile so that we can parse it later on
	defer xmlFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := io.ReadAll(xmlFile)

	// we initialize our Users array
	var eplanLabelling EplanLabelling
	// we unmarshal our byteArray which contains our
	// xmlFiles content into 'users' which we defined above
	xml.Unmarshal(byteValue, &eplanLabelling)
	// we iterate through every user within our users array and
	// print out the user Type, their name, and their facebook url
	// as just an example
	for a := 0; a < len(eplanLabelling.Document.Page.Lines); a++ {
		for b := 0; b < len(eplanLabelling.Document.Page.Lines[a].Labels); b++ {
			fmt.Printf("Id: %-20s", eplanLabelling.Document.Page.Lines[a].Labels[b].Id)
			fmt.Printf("\n")
			for c := 0; c < len(eplanLabelling.Document.Page.Lines[a].Labels[b].Properties); c++ {
				fmt.Printf("Property Name: %-50s", eplanLabelling.Document.Page.Lines[a].Labels[b].Properties[c].PropertyName)
				fmt.Printf("Property Value: %-50s", eplanLabelling.Document.Page.Lines[a].Labels[b].Properties[c].PropertyValue)
				fmt.Printf("\n")
			}
		}
	}

	fmt.Println("Verbindungsliste Fertig")
}
