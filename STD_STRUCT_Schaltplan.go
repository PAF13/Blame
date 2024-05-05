package main

type Bauteil struct {
	BMK                             BMK
	Funktionstext                   string
	TechnischeKenngrößen            string
	FunktionsdefinitionKategorie    string
	FunktionsdefinitionGruppe       string
	FunktionsdefinitionBeschreibung string
	AnschlussbezeichnungderFunktion string
	Funktionsdefinition             string
	Symbolname                      string
	Symbolvariante                  string
}
type BMK struct {
	//BMK
	BMKVoll              string
	BMKID                string
	BMK                  string
	FunktionaleZuordnung string
	Funktionskennzeichen string
	Aufstellungsort      string
	Ortskennzeichen      string
	BMKKennbuchstabe     string
}

type Kable struct {
	BMK        BMK
	Verbindung []Verbindung
}
type Verbindung struct {
	//Info from XML
	ID                        string
	Bauteil                   [2]Bauteil
	Name                      BMK
	VerbindungZugehörigkeit   string
	Verbindungsquerschnitt    string
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