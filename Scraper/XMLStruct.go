package main

import "encoding/xml"

// Define the structs based on the XML structure
type EplanPxfRoot struct {
	XMLName            xml.Name          `xml:"EplanPxfRoot"`
	Name               string            `xml:"Name,attr"`
	Type               string            `xml:"Type,attr"`
	Version            string            `xml:"Version,attr"`
	PxfVersion         string            `xml:"PxfVersion,attr"`
	SchemaVersion      string            `xml:"SchemaVersion,attr"`
	Source             string            `xml:"Source,attr"`
	SourceProject      string            `xml:"SourceProject,attr"`
	Description        string            `xml:"Description,attr"`
	ConfigurationFlags string            `xml:"ConfigurationFlags,attr"`
	NumMainObjects     int               `xml:"NumMainObjects,attr"`
	NumProjectSteps    int               `xml:"NumProjectSteps,attr"`
	NumMDSteps         int               `xml:"NumMDSteps,attr"`
	StreamSchema       string            `xml:"StreamSchema,attr"`
	DataConfiguration  DataConfiguration `xml:"DataConfiguration"`
	SO117              []SO117           `xml:"O117"`
	SO99               []SO99            `xml:"O99"`
	Objects            []Object          `xml:",any"` // ca// Catch all unspecified elements as raw XML
}

type DataConfiguration struct {
	Languages          string             `xml:"languages,attr"`
	XMemberDataEntries []XMemberDataEntry `xml:"XMemberDataEntry"`
	XPropertyEntries   []XPropertyEntry   `xml:"XPropertyEntry"`
	Objects            []Object           `xml:",any"` // ca// Catch all unspecified elements as raw XML
}

type XMemberDataEntry struct {
	Name        string   `xml:"Name,attr"`
	ID          int      `xml:"ID,attr"`
	Readonly    int      `xml:"Readonly,attr"`
	Description string   `xml:"Description,attr"`
	Objects     []Object `xml:",any"` // Catch all unspecified elements as raw XML
}

type XPropertyEntry struct {
	Name          string   `xml:"Name,attr"`
	RelationId    int      `xml:"RelationId,attr"`
	RelationIndex int      `xml:"RelationIndex,attr"`
	ID            int      `xml:"ID,attr"`
	Index         int      `xml:"Index,attr"`
	IdentName     string   `xml:"IdentName,attr"`
	Readonly      int      `xml:"Readonly,attr"`
	Description   string   `xml:"Description,attr"`
	Objects       []Object `xml:",any"` // Catch all unspecified elements as raw XML
}

type SO117 struct {
	Build   string      `xml:"Build,attr"`
	A1      string      `xml:"A1,attr"`
	A15     string      `xml:"A15,attr"`
	P22001  string      `xml:"P22001,attr"`
	P22002  string      `xml:"P22002,attr"`
	P22056  string      `xml:"P22056,attr"`
	P22003  string      `xml:"P22003,attr"`
	P22024  string      `xml:"P22024,attr"`
	P22145  string      `xml:"P22145,attr"`
	P22010  string      `xml:"P22010,attr"`
	P22007  string      `xml:"P22007,attr"`
	P22222  string      `xml:"P22222,attr"`
	P22008  string      `xml:"P22008,attr"`
	P22223  string      `xml:"P22223,attr"`
	P22258  string      `xml:"P22258,attr"`
	P22013  string      `xml:"P22013,attr"`
	P22012  string      `xml:"P22012,attr"`
	P22014  string      `xml:"P22014,attr"`
	P22033  string      `xml:"P22033,attr"`
	P22070  string      `xml:"P22070,attr"`
	P22052  string      `xml:"P22052,attr"`
	P22071  string      `xml:"P22071,attr"`
	P22073  string      `xml:"P22073,attr"`
	P22074  string      `xml:"P22074,attr"`
	P22072  string      `xml:"P22072,attr"`
	P22140  string      `xml:"P22140,attr"`
	P22142  string      `xml:"P22142,attr"`
	P22027  string      `xml:"P22027,attr"`
	P22139  string      `xml:"P22139,attr"`
	P22367  string      `xml:"P22367,attr"`
	M22004  TextElement `xml:"M22004"`
	M22005  TextElement `xml:"M22005"`
	M22006  TextElement `xml:"M22006"`
	M22009  TextElement `xml:"M22009"`
	Objects []Object    `xml:",any"` // Catch all unspecified elements as raw XML
}

type TextElement struct {
	Text    string   `xml:"T"`
	Objects []Object `xml:",any"` // Catch all unspecified elements as raw XML
}

type SO99 struct {
	Build   string   `xml:"Build,attr"`
	A15     string   `xml:"A15,attr"`
	P21004  string   `xml:"P21004,attr"`
	P21024  string   `xml:"P21024,attr"`
	P20038  string   `xml:"P20038,attr"`
	P20039  string   `xml:"P20039,attr"`
	P20296  string   `xml:"P20296,attr"`
	P20375  string   `xml:"P20375,attr"`
	P21001  string   `xml:"P21001,attr"`
	P20027  string   `xml:"P20027,attr"`
	P31004  string   `xml:"P31004,attr"`
	P31002  string   `xml:"P31002,attr"`
	Objects []Object `xml:",any"` // Catch all unspecified elements as raw XML
}

// Object struct to catch any missed tags
type Object struct {
	XMLName xml.Name
	Content string     `xml:",innerxml"`
	Attrs   []xml.Attr `xml:",any,attr"`
}
