package main

var lagerbestand_Siteca = []*Artikel{}
var lagerbestand_Kroenert = []*Artikel{}

func (a *App) BlameStartup() bool {
	INIT_SETTINGS()
	//INIT_ARTIKELSTAMMDATEN()
	INIT_VERBINDUNGSLITE()
	INIT_STUECKLISTE()
	return true
}
