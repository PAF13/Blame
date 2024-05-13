package main

type Verbindung struct {
	//Info from XML
	ID                        string
	VerbindungZugehörigkeit   string
	Verbindungsquerschnitt    float64
	Verbindungsfarbeundnummer string
	VerbindungLänge           int
	Netzname                  string
	Signalname                string
	Potenzialname             string
	Potenzialtyp              string
	Potenzialwert             string
	Netzindex                 string
	Fehler                    []string
}
