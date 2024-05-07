package main

import "encoding/xml"

// EplanLabelling was generated 2024-05-06 17:49:48 by https://xml-to-go.github.io/ in Ukraine.
type EplanAuswertungXML struct {
	XMLName  xml.Name `xml:"EplanLabelling"`
	Text     string   `xml:",chardata"`
	Version  string   `xml:"version,attr"`
	Document struct {
		Text     string `xml:",chardata"`
		SourceID string `xml:"source_id,attr"`
		Page     struct {
			Text         string `xml:",chardata"`
			SourceID     string `xml:"source_id,attr"`
			Header       string `xml:"Header"`
			ColumnHeader []struct {
				Text         string `xml:",chardata"`
				DataType     string `xml:"DataType,attr"`
				PropertyName string `xml:"PropertyName"`
			} `xml:"ColumnHeader"`
			Line []struct {
				Text      string `xml:",chardata"`
				SourceID  string `xml:"source_id,attr"`
				Separator string `xml:"separator,attr"`
				Label     struct {
					Text     string `xml:",chardata"`
					SourceID string `xml:"source_id,attr"`
					Property []struct {
						Text             string `xml:",chardata"`
						FormattingType   string `xml:"FormattingType,attr"`
						FormattingLength string `xml:"FormattingLength,attr"`
						FormattingRAlign string `xml:"FormattingRAlign,attr"`
						PropertyName     string `xml:"PropertyName"`
						PropertyValue    string `xml:"PropertyValue"`
					} `xml:"Property"`
				} `xml:"Label"`
			} `xml:"Line"`
			Footer string `xml:"Footer"`
		} `xml:"Page"`
	} `xml:"Document"`
}  

// Settings was generated 2024-05-06 17:55:12 by https://xml-to-go.github.io/ in Ukraine.
type Settings struct {
	XMLName xml.Name `xml:"Settings"`
	Text    string   `xml:",chardata"`
	Format  string   `xml:"format,attr"`
	CAT     struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		MOD  struct {
			Text string `xml:",chardata"`
			Name string `xml:"name,attr"`
			LEV1 struct {
				Text     string `xml:",chardata"`
				Name     string `xml:"name,attr"`
				Nodekind string `xml:"nodekind,attr"`
				LEV2     struct {
					Text string `xml:",chardata"`
					Name string `xml:"name,attr"`
					LEV3 struct {
						Text     string `xml:",chardata"`
						Name     string `xml:"name,attr"`
						Nodekind string `xml:"nodekind,attr"`
						LEV4     []struct {
							Text    string `xml:",chardata"`
							Name    string `xml:"name,attr"`
							Setting []struct {
								Text string `xml:",chardata"`
								Name string `xml:"name,attr"`
								Type string `xml:"type,attr"`
								Val  string `xml:"Val"`
							} `xml:"Setting"`
							LEV5 struct {
								Text    string `xml:",chardata"`
								Name    string `xml:"name,attr"`
								Setting []struct {
									Text string   `xml:",chardata"`
									Name string   `xml:"name,attr"`
									Type string   `xml:"type,attr"`
									Val  []string `xml:"Val"`
								} `xml:"Setting"`
								LEV6 []struct {
									Text    string `xml:",chardata"`
									Name    string `xml:"name,attr"`
									Setting []struct {
										Text string `xml:",chardata"`
										Name string `xml:"name,attr"`
										Type string `xml:"type,attr"`
										Val  string `xml:"Val"`
									} `xml:"Setting"`
								} `xml:"LEV6"`
							} `xml:"LEV5"`
						} `xml:"LEV4"`
						Setting []struct {
							Text string `xml:",chardata"`
							Name string `xml:"name,attr"`
							Type string `xml:"type,attr"`
							Val  string `xml:"Val"`
						} `xml:"Setting"`
					} `xml:"LEV3"`
					Setting []struct {
						Text string `xml:",chardata"`
						Name string `xml:"name,attr"`
						Type string `xml:"type,attr"`
						Val  string `xml:"Val"`
					} `xml:"Setting"`
				} `xml:"LEV2"`
			} `xml:"LEV1"`
		} `xml:"MOD"`
	} `xml:"CAT"`
} 