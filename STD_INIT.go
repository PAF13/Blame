package main

var lagerbestand_Siteca = []*Artikel{}
var lagerbestand_Kroenert = []*Artikel{}
var lagerbestand_Siteca_Map = map[string]*Artikel{}
var Stueckliste2 *[]*Artikel

func (a *App) BlameStartup() bool {
	INIT_SETTINGS()
	INIT_ARTIKELSTAMMDATEN()
	return true
}
