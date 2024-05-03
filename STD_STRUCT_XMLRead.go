package main

import "encoding/xml"

//Eplan Stammdaten Export XML
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

//Eplan Export XML
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
	Header  []Header `xml:"Header"`
	Footer  []Footer `xml:"Footer"`
}

type Line struct {
	XMLName xml.Name `xml:"Line"`
	Id      string   `xml:"source_id,attr"`
	Labels  []Label  `xml:"Label"`
}

type Header struct {
	XMLName    xml.Name   `xml:"Header"`
	Properties []Property `xml:"Property"`
}
type Footer struct {
	XMLName    xml.Name   `xml:"Footer"`
	Properties []Property `xml:"Property"`
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
